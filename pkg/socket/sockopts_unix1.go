//go:build darwin || dragonfly || linux || netbsd || openbsd

package socket

func SetReuseport(fd, reusePort int) error { _ = "STUB: not implemented"; return nil }
