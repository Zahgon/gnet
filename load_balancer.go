package gnet

import (
	"net"
)

type LoadBalancing int

const (
	RoundRobin LoadBalancing = iota

	LeastConnections

	SourceAddrHash
)

type (
	loadBalancer interface {
		register(*eventloop)
		next(net.Addr) *eventloop
		index(int) *eventloop
		iterate(func(int, *eventloop) bool)
		len() int
	}

	baseLoadBalancer struct {
		eventLoops []*eventloop
		size       int
	}

	roundRobinLoadBalancer struct {
		baseLoadBalancer
		nextIndex uint64
	}

	leastConnectionsLoadBalancer struct {
		baseLoadBalancer
	}

	sourceAddrHashLoadBalancer struct {
		baseLoadBalancer
	}
)

func (lb *baseLoadBalancer) register(el *eventloop) { _ = "STUB: not implemented"; return }

func (lb *baseLoadBalancer) index(i int) *eventloop { _ = "STUB: not implemented"; return nil }

func (lb *baseLoadBalancer) iterate(f func(int, *eventloop) bool) {
	_ = "STUB: not implemented"
	return
}

func (lb *baseLoadBalancer) len() int { _ = "STUB: not implemented"; return 0 }

func (lb *roundRobinLoadBalancer) next(_ net.Addr) (el *eventloop) {
	_ = "STUB: not implemented"
	return nil
}

func (lb *leastConnectionsLoadBalancer) next(_ net.Addr) (el *eventloop) {
	_ = "STUB: not implemented"
	return nil
}

func (*sourceAddrHashLoadBalancer) hash(s string) int { _ = "STUB: not implemented"; return 0 }

func (lb *sourceAddrHashLoadBalancer) next(netAddr net.Addr) *eventloop {
	_ = "STUB: not implemented"
	return nil
}
