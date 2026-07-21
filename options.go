package gnet

import (
	"time"

	"github.com/panjf2000/gnet/v2/pkg/logging"
)

type Option func(opts *Options)

func loadOptions(options ...Option) *Options { _ = "STUB: not implemented"; return nil }

type TCPSocketOpt int

const (
	TCPNoDelay TCPSocketOpt = iota
	TCPDelay
)

type Options struct {
	LB LoadBalancing

	ReuseAddr bool

	ReusePort bool

	MulticastInterfaceIndex int

	BindToDevice string

	Multicore bool

	NumEventLoop int

	ReadBufferCap int

	WriteBufferCap int

	LockOSThread bool

	Ticker bool

	TCPKeepAlive time.Duration

	TCPKeepInterval time.Duration

	TCPKeepCount int

	TCPNoDelay TCPSocketOpt

	SocketRecvBuffer int

	SocketSendBuffer int

	LogPath string

	LogLevel logging.Level

	Logger logging.Logger

	EdgeTriggeredIO bool

	EdgeTriggeredIOChunk int
}

func WithOptions(options Options) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMulticore(multicore bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithLockOSThread(lockOSThread bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithReadBufferCap(readBufferCap int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithWriteBufferCap(writeBufferCap int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithLoadBalancing(lb LoadBalancing) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithNumEventLoop(numEventLoop int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithReusePort(reusePort bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithReuseAddr(reuseAddr bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTCPKeepAlive(tcpKeepAlive time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithTCPKeepInterval(tcpKeepInterval time.Duration) Option {
	_ = "STUB: not implemented"
	return *new(Option)
}

func WithTCPKeepCount(tcpKeepCount int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTCPNoDelay(tcpNoDelay TCPSocketOpt) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithSocketRecvBuffer(recvBuf int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithSocketSendBuffer(sendBuf int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithTicker(ticker bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithLogPath(fileName string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithLogLevel(lvl logging.Level) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithLogger(logger logging.Logger) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithMulticastInterfaceIndex(idx int) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithBindToDevice(iface string) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithEdgeTriggeredIO(et bool) Option { _ = "STUB: not implemented"; return *new(Option) }

func WithEdgeTriggeredIOChunk(chunk int) Option { _ = "STUB: not implemented"; return *new(Option) }
