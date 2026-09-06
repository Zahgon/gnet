//go:build !arm64 && !riscv64 && poll_opt

package netpoll

func epollWait(epfd int, events []epollevent, msec int) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}
