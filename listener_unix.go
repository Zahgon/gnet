//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd

package gnet

import (
	"net"
	"sync"

	"github.com/panjf2000/gnet/v2/pkg/netpoll"
	"github.com/panjf2000/gnet/v2/pkg/socket"
)

type listener struct {
	openOnce, closeOnce sync.Once
	fd                  int
	addr                net.Addr
	address, network    string
	sockOptInts         []socket.Option[int]
	sockOptStrs         []socket.Option[string]
	pollAttachment      *netpoll.PollAttachment
}

func (ln *listener) packPollAttachment(handler netpoll.PollEventHandler) *netpoll.PollAttachment {
	_ = "STUB: not implemented"
	return nil
}

func (ln *listener) dup() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (ln *listener) open() (err error) { _ = "STUB: not implemented"; return nil }

func (ln *listener) close() { _ = "STUB: not implemented"; return }

func initListener(network, addr string, options *Options) (ln *listener, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}
