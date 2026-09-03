package gnet

import (
	"context"
	"net"
)

type contextKey struct{}

func NewContext(ctx context.Context, v any) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func FromContext(ctx context.Context) any { _ = "STUB: not implemented"; return *new(any) }

type connContextKey struct{}

func NewNetConnContext(ctx context.Context, c net.Conn) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func FromNetConnContext(ctx context.Context) (net.Conn, bool) {
	_ = "STUB: not implemented"
	return *new(net.Conn), false
}

type netAddrContextKey struct{}

func NewNetAddrContext(ctx context.Context, a net.Addr) context.Context {
	_ = "STUB: not implemented"
	return *new(context.Context)
}

func FromNetAddrContext(ctx context.Context) (net.Addr, bool) {
	_ = "STUB: not implemented"
	return *new(net.Addr), false
}
