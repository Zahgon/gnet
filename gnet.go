package gnet

import (
	"context"
	"io"
	"net"
	"sync"
	"time"
)

type Action int

const (
	None Action = iota

	Close

	Shutdown
)

type Engine struct {
	eng *engine
}

func (e Engine) Validate() error { _ = "STUB: not implemented"; return nil }

func (e Engine) CountConnections() (count int) { _ = "STUB: not implemented"; return 0 }

func (e Engine) Register(ctx context.Context) (<-chan RegisteredResult, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (e Engine) Dup() (fd int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (e Engine) DupListener(network, addr string) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (e Engine) Stop(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

type Reader interface {
	io.Reader
	io.WriterTo

	Next(n int) (buf []byte, err error)

	Peek(n int) (buf []byte, err error)

	Discard(n int) (discarded int, err error)

	InboundBuffered() int
}

type Writer interface {
	io.Writer
	io.ReaderFrom

	SendTo(buf []byte, addr net.Addr) (n int, err error)

	Writev(bs [][]byte) (n int, err error)

	Flush() error

	OutboundBuffered() int

	AsyncWrite(buf []byte, callback AsyncCallback) (err error)

	AsyncWritev(bs [][]byte, callback AsyncCallback) (err error)
}

type AsyncCallback func(c Conn, err error) error

type Socket interface {
	Fd() int

	Dup() (int, error)

	SetReadBuffer(size int) error

	SetWriteBuffer(size int) error

	SetLinger(secs int) error

	SetKeepAlivePeriod(d time.Duration) error

	SetKeepAlive(enabled bool, idle, intvl time.Duration, cnt int) error

	SetNoDelay(noDelay bool) error
}

type Runnable interface {
	Run(ctx context.Context) error
}

type RunnableFunc func(ctx context.Context) error

func (fn RunnableFunc) Run(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

type RegisteredResult struct {
	Conn Conn
	Err  error
}

type EventLoop interface {
	Register(ctx context.Context, addr net.Addr) (<-chan RegisteredResult, error)

	Enroll(ctx context.Context, c net.Conn) (<-chan RegisteredResult, error)

	Execute(ctx context.Context, runnable Runnable) error

	Schedule(ctx context.Context, runnable Runnable, delay time.Duration) error

	Close(Conn) error
}

type Conn interface {
	Reader
	Writer
	Socket

	Context() (ctx any)

	SafeContext() (ctx any)

	EventLoop() EventLoop

	SetContext(ctx any)

	SetSafeContext(ctx any)

	LocalAddr() net.Addr

	RemoteAddr() net.Addr

	Wake(callback AsyncCallback) error

	CloseWithCallback(callback AsyncCallback) error

	Close() error

	SetDeadline(time.Time) error

	SetReadDeadline(time.Time) error

	SetWriteDeadline(time.Time) error
}

type (
	EventHandler interface {
		OnBoot(eng Engine) (action Action)

		OnShutdown(eng Engine)

		OnOpen(c Conn) (out []byte, action Action)

		OnClose(c Conn, err error) (action Action)

		OnTraffic(c Conn) (action Action)

		OnTick() (delay time.Duration, action Action)
	}

	BuiltinEventEngine struct{}
)

func (*BuiltinEventEngine) OnBoot(_ Engine) (action Action) {
	_ = "STUB: not implemented"
	return *new(Action)
}

func (*BuiltinEventEngine) OnShutdown(_ Engine) { _ = "STUB: not implemented"; return }

func (*BuiltinEventEngine) OnOpen(_ Conn) (out []byte, action Action) {
	_ = "STUB: not implemented"
	return nil, *new(Action)
}

func (*BuiltinEventEngine) OnClose(_ Conn, _ error) (action Action) {
	_ = "STUB: not implemented"
	return *new(Action)
}

func (*BuiltinEventEngine) OnTraffic(_ Conn) (action Action) {
	_ = "STUB: not implemented"
	return *new(Action)
}

func (*BuiltinEventEngine) OnTick() (delay time.Duration, action Action) {
	_ = "STUB: not implemented"
	return *new(time.Duration), *new(Action)
}

var MaxStreamBufferCap = 64 * 1024

func createListeners(addrs []string, opts ...Option) ([]*listener, *Options, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

func Run(eventHandler EventHandler, protoAddr string, opts ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

func Rotate(eventHandler EventHandler, addrs []string, opts ...Option) error {
	_ = "STUB: not implemented"
	return nil
}

var (
	allEngines sync.Map

	shutdownPollInterval = 500 * time.Millisecond
)

func Stop(ctx context.Context, protoAddr string) error { _ = "STUB: not implemented"; return nil }

func parseProtoAddr(protoAddr string) (string, string, error) {
	_ = "STUB: not implemented"
	return "", "", nil
}

func determineEventLoops(opts *Options) int { _ = "STUB: not implemented"; return 0 }
