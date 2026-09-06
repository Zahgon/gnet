//go:build dragonfly || freebsd || netbsd || openbsd

package socket

func SetBindToDevice(_ int, _ string) error { _ = "STUB: not implemented"; return nil }
