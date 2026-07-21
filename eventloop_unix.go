//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd

package gnet

import (
	"context"
	"net"
	"time"

	"github.com/panjf2000/gnet/v2/pkg/logging"
	"github.com/panjf2000/gnet/v2/pkg/netpoll"
)

type eventloop struct {
	listeners    map[int]*listener
	idx          int
	engine       *engine
	poller       *netpoll.Poller
	buffer       []byte
	connections  connMatrix
	eventHandler EventHandler
}

func (el *eventloop) Register(ctx context.Context, addr net.Addr) (<-chan RegisteredResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (el *eventloop) Enroll(ctx context.Context, c net.Conn) (<-chan RegisteredResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (el *eventloop) Execute(ctx context.Context, runnable Runnable) error {
	_ = "STUB: not implemented"
	return nil
}

func (el *eventloop) Schedule(context.Context, Runnable, time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (el *eventloop) Close(c Conn) error { _ = "STUB: not implemented"; return nil }

func (el *eventloop) getLogger() logging.Logger {
	_ = "STUB: not implemented"
	return *new(logging.Logger)
}

func (el *eventloop) countConn() int32 { _ = "STUB: not implemented"; return 0 }

func (el *eventloop) closeConns() { _ = "STUB: not implemented"; return }

type connWithCallback struct {
	c  *conn
	cb func()
}

func (el *eventloop) enroll(c net.Conn, addr net.Addr, ctx any) (resCh chan RegisteredResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

//nolint:errcheck

//nolint:errcheck

func (el *eventloop) register(a any) error { _ = "STUB: not implemented"; return nil }

func (el *eventloop) register0(c *conn) error { _ = "STUB: not implemented"; return nil }

func (el *eventloop) open(c *conn) error { _ = "STUB: not implemented"; return nil }

func (el *eventloop) read0(a any) error { _ = "STUB: not implemented"; return nil }

func (el *eventloop) read(c *conn) error { _ = "STUB: not implemented"; return nil }

func (el *eventloop) write0(a any) error { _ = "STUB: not implemented"; return nil }

const iovMax = 1024

func (el *eventloop) write(c *conn) error { _ = "STUB: not implemented"; return nil }

func (el *eventloop) close(c *conn, err error) error { _ = "STUB: not implemented"; return nil }

func (el *eventloop) wake(c *conn) error { _ = "STUB: not implemented"; return nil }

func (el *eventloop) ticker(ctx context.Context) { _ = "STUB: not implemented"; return }

func (el *eventloop) readUDP(fd int, _ netpoll.IOEvent, _ netpoll.IOFlags) error {
	_ = "STUB: not implemented"
	return nil
}

func (el *eventloop) handleAction(c *conn, action Action) error {
	_ = "STUB: not implemented"
	return nil
}
