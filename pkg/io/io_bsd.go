//go:build darwin || dragonfly || freebsd || netbsd || openbsd

package io

import (
	"golang.org/x/sys/unix"
)

func Writev(fd int, bs [][]byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:staticcheck

func Readv(fd int, bs [][]byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

//nolint:staticcheck

var _zero uintptr

func bytes2iovec(bs [][]byte) []unix.Iovec { _ = "STUB: not implemented"; return nil }
