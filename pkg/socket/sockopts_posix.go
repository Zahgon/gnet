//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd

package socket

import (
	"net"
)

func SetNoDelay(fd, noDelay int) error { _ = "STUB: not implemented"; return nil }

func SetRecvBuffer(fd, size int) error { _ = "STUB: not implemented"; return nil }

func SetSendBuffer(fd, size int) error { _ = "STUB: not implemented"; return nil }

func SetReuseAddr(fd, reuseAddr int) error { _ = "STUB: not implemented"; return nil }

func SetIPv6Only(fd, ipv6only int) error { _ = "STUB: not implemented"; return nil }

func SetLinger(fd, sec int) error { _ = "STUB: not implemented"; return nil }

func SetMulticastMembership(proto string, udpAddr *net.UDPAddr) func(int, int) error {
	_ = "STUB: not implemented"
	return nil
}

func SetIPv4MulticastMembership(fd int, mcast net.IP, ifIndex int) error {
	_ = "STUB: not implemented"
	return nil
}

func SetIPv6MulticastMembership(fd int, mcast net.IP, ifIndex int) error {
	_ = "STUB: not implemented"
	return nil
}

func interfaceFirstIPv4Addr(ifIndex int) (net.IP, error) {
	_ = "STUB: not implemented"
	return *new(net.IP), nil
}
