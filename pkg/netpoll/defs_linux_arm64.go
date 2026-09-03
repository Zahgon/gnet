//go:build poll_opt

package netpoll

type epollevent struct {
	events uint32
	_pad   uint32
	data   [8]byte
}
