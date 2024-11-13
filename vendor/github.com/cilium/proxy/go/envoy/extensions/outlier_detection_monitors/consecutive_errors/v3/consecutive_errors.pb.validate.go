// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/outlier_detection_monitors/consecutive_errors/v3/consecutive_errors.proto

package consecutive_errorsv3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on ConsecutiveErrors with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ConsecutiveErrors) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConsecutiveErrors with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConsecutiveErrorsMultiError, or nil if none found.
func (m *ConsecutiveErrors) ValidateAll() error {
	return m.validate(true)
}

func (m *ConsecutiveErrors) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if wrapper := m.GetThreshold(); wrapper != nil {

		if wrapper.GetValue() > 100 {
			err := ConsecutiveErrorsValidationError{
				field:  "Threshold",
				reason: "value must be less than or equal to 100",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if wrapper := m.GetEnforcing(); wrapper != nil {

		if wrapper.GetValue() > 100 {
			err := ConsecutiveErrorsValidationError{
				field:  "Enforcing",
				reason: "value must be less than or equal to 100",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if all {
		switch v := interface{}(m.GetErrorBuckets()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConsecutiveErrorsValidationError{
					field:  "ErrorBuckets",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConsecutiveErrorsValidationError{
					field:  "ErrorBuckets",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetErrorBuckets()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConsecutiveErrorsValidationError{
				field:  "ErrorBuckets",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ConsecutiveErrorsMultiError(errors)
	}

	return nil
}

// ConsecutiveErrorsMultiError is an error wrapping multiple validation errors
// returned by ConsecutiveErrors.ValidateAll() if the designated constraints
// aren't met.
type ConsecutiveErrorsMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConsecutiveErrorsMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConsecutiveErrorsMultiError) AllErrors() []error { return m }

// ConsecutiveErrorsValidationError is the validation error returned by
// ConsecutiveErrors.Validate if the designated constraints aren't met.
type ConsecutiveErrorsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConsecutiveErrorsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConsecutiveErrorsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConsecutiveErrorsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConsecutiveErrorsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConsecutiveErrorsValidationError) ErrorName() string {
	return "ConsecutiveErrorsValidationError"
}

// Error satisfies the builtin error interface
func (e ConsecutiveErrorsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConsecutiveErrors.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConsecutiveErrorsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConsecutiveErrorsValidationError{}
