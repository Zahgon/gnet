//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd

package socket

import (
	"net"

	"golang.org/x/sys/unix"
)

func GetUnixSockAddr(proto, addr string) (sa unix.Sockaddr, family int, unixAddr *net.UnixAddr, err error) {
	_ = "STUB: not implemented"
	return *new(unix.Sockaddr), 0, nil, nil
}

func udsSocket(proto, addr string, passive bool, sockOptInts []Option[int], sockOptStrs []Option[string]) (fd int, netAddr net.Addr, err error) {
	_ = "STUB: not implemented"
	return 0, *new(net.Addr), nil
}
