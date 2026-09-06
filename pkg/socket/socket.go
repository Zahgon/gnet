//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd

package socket

import (
	"net"

	"golang.org/x/sys/unix"
)

type Option[T int | string] struct {
	SetSockOpt func(int, T) error
	Opt        T
}

func execSockOpts[T int | string](fd int, opts []Option[T]) error {
	_ = "STUB: not implemented"
	return nil
}

func TCPSocket(proto, addr string, passive bool, sockOptInts []Option[int], sockOptStrs []Option[string]) (int, net.Addr, error) {
	_ = "STUB: not implemented"
	return 0, *new(net.Addr), nil
}

func UDPSocket(proto, addr string, connect bool, sockOptInts []Option[int], sockOptStrs []Option[string]) (int, net.Addr, error) {
	_ = "STUB: not implemented"
	return 0, *new(net.Addr), nil
}

func UnixSocket(proto, addr string, passive bool, sockOptInts []Option[int], sockOptStrs []Option[string]) (int, net.Addr, error) {
	_ = "STUB: not implemented"
	return 0, *new(net.Addr), nil
}

func Accept(fd int) (int, unix.Sockaddr, error) {
	_ = "STUB: not implemented"
	return 0, *new(unix.Sockaddr), nil
}
