//go:build dragonfly || freebsd || linux

package socket

import "golang.org/x/sys/unix"

func sysSocket(family, sotype, proto int) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func sysAccept(fd int) (nfd int, sa unix.Sockaddr, err error) {
	_ = "STUB: not implemented"
	return 0, *new(unix.Sockaddr), nil
}
