// SPDX-License-Identifier: Apache-2.0
// Copyright Authors of Cilium

package policycell

import (
	"github.com/cilium/hive/cell"
	"github.com/spf13/pflag"

	cmtypes "github.com/cilium/cilium/pkg/clustermesh/types"
	"github.com/cilium/cilium/pkg/crypto/certificatemanager"
	"github.com/cilium/cilium/pkg/endpointmanager"
	"github.com/cilium/cilium/pkg/envoy"
	"github.com/cilium/cilium/pkg/identity"
	"github.com/cilium/cilium/pkg/identity/identitymanager"
	"github.com/cilium/cilium/pkg/metrics"
	"github.com/cilium/cilium/pkg/option"
	"github.com/cilium/cilium/pkg/policy"
	"github.com/cilium/cilium/pkg/policy/api"
)

// Cell provides the PolicyRepository and PolicyUpdater.
var Cell = cell.Module(
	"policy",
	"Contains policy rules",

	cell.Provide(newPolicyRepo),
	cell.Provide(newPolicyUpdater),
	cell.Provide(newPolicyImporter),
	cell.Config(defaultConfig),
)

type Config struct {
	PolicyQueueSize uint `mapstructure:"policy-queue-size"`
}

var defaultConfig = Config{
	PolicyQueueSize: 100,
}

func (def Config) Flags(flags *pflag.FlagSet) {
	flags.Uint("policy-queue-size", def.PolicyQueueSize, "Size of queue for policy-related events")
}

type policyRepoParams struct {
	cell.In

	Lifecycle       cell.Lifecycle
	CertManager     certificatemanager.CertificateManager
	SecretManager   certificatemanager.SecretManager
	IdentityManager identitymanager.IDManager
	ClusterInfo     cmtypes.ClusterInfo
	MetricsManager  api.PolicyMetrics
}

func newPolicyRepo(params policyRepoParams) policy.PolicyRepository {
	if option.Config.EnableWellKnownIdentities {
		// Must be done before calling policy.NewPolicyRepository() below.
		num := identity.InitWellKnownIdentities(option.Config, params.ClusterInfo)
		metrics.Identity.WithLabelValues(identity.WellKnownIdentityType).Add(float64(num))
		identity.WellKnown.ForEach(func(i *identity.Identity) {
			for labelSource := range i.Labels.CollectSources() {
				metrics.IdentityLabelSources.WithLabelValues(labelSource).Inc()
			}
		})
	}

	// policy repository: maintains list of active Rules and their subject
	// security identities. Also constructs the SelectorCache, a precomputed
	// cache of label selector -> identities for policy peers.
	policyRepo := policy.NewPolicyRepository(
		identity.ListReservedIdentities(), // Load SelectorCache with reserved identities
		params.CertManager,
		params.SecretManager,
		params.IdentityManager,
		params.MetricsManager,
	)
	policyRepo.SetEnvoyRulesFunc(envoy.GetEnvoyHTTPRules)

	params.Lifecycle.Append(cell.Hook{
		OnStart: func(hc cell.HookContext) error {
			policyRepo.GetSelectorCache().RegisterMetrics()
			return nil
		},
	})

	return policyRepo
}

type policyUpdaterParams struct {
	cell.In

	Lifecycle        cell.Lifecycle
	PolicyRepository policy.PolicyRepository
	EndpointManager  endpointmanager.EndpointManager
}

func newPolicyUpdater(params policyUpdaterParams) *policy.Updater {
	// policyUpdater: forces policy recalculation on all endpoints.
	// Called for various events, such as named port changes
	// or certain identity updates.
	policyUpdater := policy.NewUpdater(params.PolicyRepository, params.EndpointManager)

	params.Lifecycle.Append(cell.Hook{
		OnStop: func(hc cell.HookContext) error {
			policyUpdater.Shutdown()
			return nil
		},
	})

	return policyUpdater
}
