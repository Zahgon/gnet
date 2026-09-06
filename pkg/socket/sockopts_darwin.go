package socket

func SetKeepAlivePeriod(fd, secs int) error { _ = "STUB: not implemented"; return nil }

func SetKeepAlive(fd int, enabled bool, idle, intvl, cnt int) error {
	_ = "STUB: not implemented"
	return nil
}

func SetBindToDevice(_ int, _ string) error { _ = "STUB: not implemented"; return nil }
