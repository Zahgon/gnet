//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd

package socket

import (
	"net"

	"golang.org/x/sys/unix"
)

func GetUDPSockAddr(proto, addr string) (sa unix.Sockaddr, family int, udpAddr *net.UDPAddr, ipv6only bool, err error) {
	_ = "STUB: not implemented"
	return *new(unix.Sockaddr), 0, nil, false, nil
}

func determineUDPProto(proto string, addr *net.UDPAddr) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func udpSocket(proto, addr string, connect bool, sockOptInts []Option[int], sockOptStrs []Option[string]) (fd int, netAddr net.Addr, err error) {
	_ = "STUB: not implemented"
	return 0, *new(net.Addr), nil
}
