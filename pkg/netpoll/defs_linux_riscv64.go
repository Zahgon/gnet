//go:build poll_opt

package netpoll

type epollevent struct {
	events    uint32
	pad_cgo_0 [4]byte
	data      [8]byte
}
