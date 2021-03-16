package ethtool

import (
	"bytes"
	"unsafe"

	"golang.org/x/sys/unix"
)

const (
	SIOCETHTOOL      = 0x8946
	ETHTOOL_GDRVINFO = 0x00000003 // Get driver info

	IFNAMSIZ = 16
)

const (
	veth = "veth"
)

type ifreq struct {
	name [IFNAMSIZ]byte
	data uintptr
}

type ethtoolDrvInfo struct {
	cmd         uint32
	driver      [32]byte
	version     [32]byte
	fwVersion   [32]byte
	busInfo     [32]byte
	eromVersion [32]byte
	reserved2   [12]byte
	nPrivFlags  uint32
	nStats      uint32
	testInfoLen uint32
	eedumpLen   uint32
	regdumpLen  uint32
}

func ethtoolIoctl(iface string, info *ethtoolDrvInfo) error {
	var ifname [IFNAMSIZ]byte

	fd, err := unix.Socket(unix.AF_INET, unix.SOCK_DGRAM, unix.IPPROTO_IP)
	if err != nil {
		return err
	}
	copy(ifname[:], iface)
	req := ifreq{
		name: ifname,
		data: uintptr(unsafe.Pointer(info)),
	}
	_, _, ep := unix.Syscall(unix.SYS_IOCTL, uintptr(fd), SIOCETHTOOL, uintptr(unsafe.Pointer(&req)))
	if ep != 0 {
		return ep
	}
	return nil
}

func GetDeviceName(iface string) (string, error) {
	info := ethtoolDrvInfo{
		cmd: ETHTOOL_GDRVINFO,
	}

	if err := ethtoolIoctl(iface, &info); err != nil {
		return "", err
	}
	return string(bytes.Trim(info.driver[:], "\x00")), nil
}

func IsNotVirtualDriver(iface string) (bool, error) {
	drvName, err := GetDeviceName(iface)
	if err != nil {
		return false, err
	}

	switch drvName {
	case veth:
		return false, nil
	}

	return true, nil
}
