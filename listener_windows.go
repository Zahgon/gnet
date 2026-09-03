package gnet

import (
	"net"
	"sync"
)

type listener struct {
	openOnce, closeOnce sync.Once
	network             string
	address             string
	lc                  *net.ListenConfig
	ln                  net.Listener
	pc                  net.PacketConn
	addr                net.Addr
}

func (l *listener) dup() (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (l *listener) open() (err error) { _ = "STUB: not implemented"; return nil }

func (l *listener) close() { _ = "STUB: not implemented"; return }

func initListener(network, addr string, options *Options) (*listener, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
