//go:build dragonfly || freebsd || linux || netbsd

package socket

func SetKeepAlivePeriod(fd, secs int) error { _ = "STUB: not implemented"; return nil }

func SetKeepAlive(fd int, enabled bool, idle, intvl, cnt int) error {
	_ = "STUB: not implemented"
	return nil
}
