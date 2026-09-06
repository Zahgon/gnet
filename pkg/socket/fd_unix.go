//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd

package socket

import (
	"sync/atomic"
)

func Dup(fd int) (int, error) { _ = "STUB: not implemented"; return 0, nil }

var tryDupCloexec atomic.Bool

func init() {
	tryDupCloexec.Store(true)
}

func dupCloseOnExec(fd int) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func dupCloseOnExecOld(fd int) (int, error) { _ = "STUB: not implemented"; return 0, nil }
