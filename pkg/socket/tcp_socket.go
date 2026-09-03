//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd

package socket

import (
	"net"

	"golang.org/x/sys/unix"
)

var listenerBacklogMaxSize = maxListenerBacklog()

func GetTCPSockAddr(proto, addr string) (sa unix.Sockaddr, family int, tcpAddr *net.TCPAddr, ipv6only bool, err error) {
	_ = "STUB: not implemented"
	return *new(unix.Sockaddr), 0, nil, false, nil
}

func determineTCPProto(proto string, addr *net.TCPAddr) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func tcpSocket(proto, addr string, passive bool, sockOptInts []Option[int], sockOptStrs []Option[string]) (fd int, netAddr net.Addr, err error) {
	_ = "STUB: not implemented"
	return 0, *new(net.Addr), nil
}
