//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd

package gnet

import (
	"context"
	"sync/atomic"
	"time"

	"golang.org/x/sync/errgroup"
)

type engine struct {
	listeners    map[int]*listener
	opts         *Options
	ingress      *eventloop
	eventLoops   loadBalancer
	inShutdown   atomic.Bool
	turnOff      context.CancelFunc
	eventHandler EventHandler
	concurrency  struct {
		*errgroup.Group

		ctx context.Context
	}
}

func (eng *engine) isShutdown() bool { _ = "STUB: not implemented"; return false }

func (eng *engine) shutdown(err error) { _ = "STUB: not implemented"; return }

func (eng *engine) closeEventLoops() { _ = "STUB: not implemented"; return }

func (eng *engine) runEventLoops(ctx context.Context, numEventLoop int) error {
	_ = "STUB: not implemented"
	return nil
}

func (eng *engine) activateReactors(ctx context.Context, numEventLoop int) error {
	_ = "STUB: not implemented"
	return nil
}

func (eng *engine) start(ctx context.Context, numEventLoop int) error {
	_ = "STUB: not implemented"
	return nil
}

func (eng *engine) stop(ctx context.Context, s Engine) { _ = "STUB: not implemented"; return }

func run(eventHandler EventHandler, listeners []*listener, options *Options, addrs []string) error {
	_ = "STUB: not implemented"
	return nil
}

func setKeepAlive(fd int, enabled bool, idle, intvl time.Duration, cnt int) error {
	_ = "STUB: not implemented"
	return nil
}
