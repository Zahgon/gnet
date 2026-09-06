//go:build darwin || dragonfly || freebsd || netbsd || openbsd

package socket

func maxListenerBacklog() int { _ = "STUB: not implemented"; return 0 }
