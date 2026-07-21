//go:build poll_opt

package netpoll

func epollCtl(epfd int, op int, fd int, event *epollevent) error {
	_ = "STUB: not implemented"
	return nil
}
