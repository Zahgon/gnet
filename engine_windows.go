package gnet

import (
	"context"
	"sync/atomic"

	"golang.org/x/sync/errgroup"
)

type engine struct {
	listeners     []*listener
	opts          *Options
	eventLoops    loadBalancer
	inShutdown    atomic.Bool
	beingShutdown atomic.Bool
	turnOff       context.CancelFunc
	eventHandler  EventHandler
	concurrency   struct {
		*errgroup.Group

		ctx context.Context
	}
}

func (eng *engine) isShutdown() bool { _ = "STUB: not implemented"; return false }

func (eng *engine) shutdown(err error) { _ = "STUB: not implemented"; return }

func (eng *engine) closeEventLoops() { _ = "STUB: not implemented"; return }

func (eng *engine) start(ctx context.Context, numEventLoop int) error {
	_ = "STUB: not implemented"
	return nil
}

func (eng *engine) stop(ctx context.Context, engine Engine) { _ = "STUB: not implemented"; return }

func run(eventHandler EventHandler, listeners []*listener, options *Options, addrs []string) error {
	_ = "STUB: not implemented"
	return nil
}
