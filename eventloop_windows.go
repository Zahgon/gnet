package gnet

import (
	"context"
	"net"
	"time"

	"github.com/panjf2000/gnet/v2/pkg/logging"
)

type eventloop struct {
	ch           chan any
	idx          int
	eng          *engine
	connCount    int32
	connections  map[*conn]struct{}
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

func (el *eventloop) enroll(c net.Conn, addr net.Addr, ctx any) (resCh chan RegisteredResult, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (el *eventloop) incConn(delta int32) { _ = "STUB: not implemented"; return }

func (el *eventloop) countConn() int32 { _ = "STUB: not implemented"; return 0 }

func (el *eventloop) run() (err error) { _ = "STUB: not implemented"; return nil }

func (el *eventloop) open(oc *openConn) error { _ = "STUB: not implemented"; return nil }

func (el *eventloop) read(c *conn) error { _ = "STUB: not implemented"; return nil }

func (el *eventloop) readUDP(c *conn) error { _ = "STUB: not implemented"; return nil }

func (el *eventloop) ticker(ctx context.Context) { _ = "STUB: not implemented"; return }

func (el *eventloop) wake(c *conn) error { _ = "STUB: not implemented"; return nil }

func (el *eventloop) close(c *conn, err error) error { _ = "STUB: not implemented"; return nil }

func (el *eventloop) handleAction(c *conn, action Action) error {
	_ = "STUB: not implemented"
	return nil
}
