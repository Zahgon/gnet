//go:build poll_opt

package netpoll

import "golang.org/x/sys/unix"

var (
	errEAGAIN error = unix.EAGAIN
	errEINVAL error = unix.EINVAL
	errENOENT error = unix.ENOENT
)

func errnoErr(e unix.Errno) error { _ = "STUB: not implemented"; return nil }

var zero uintptr
