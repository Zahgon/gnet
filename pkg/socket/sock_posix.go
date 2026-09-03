//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd

package socket

import (
	"net"

	"golang.org/x/sys/unix"
)

func ipToSockaddrInet4(ip net.IP, port int) (unix.SockaddrInet4, error) {
	_ = "STUB: not implemented"
	return *new(unix.SockaddrInet4), nil
}

func ipToSockaddrInet6(ip net.IP, port int, zone string) (unix.SockaddrInet6, error) {
	_ = "STUB: not implemented"
	return *new(unix.SockaddrInet6), nil
}

func ipToSockaddr(family int, ip net.IP, port int, zone string) (unix.Sockaddr, error) {
	_ = "STUB: not implemented"
	return *new(unix.Sockaddr), nil
}
