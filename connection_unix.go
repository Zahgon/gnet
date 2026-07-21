//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd

package gnet

import (
	"io"
	"net"
	"sync/atomic"
	"time"

	"golang.org/x/sys/unix"

	"github.com/panjf2000/gnet/v2/internal/gfd"
	"github.com/panjf2000/gnet/v2/pkg/buffer/elastic"
	"github.com/panjf2000/gnet/v2/pkg/netpoll"
)

type conn struct {
	fd             int
	gfd            gfd.GFD
	ctx            any
	safeCtx        atomic.Pointer[any]
	remote         unix.Sockaddr
	proto          string
	localAddr      net.Addr
	remoteAddr     net.Addr
	loop           *eventloop
	outboundBuffer elastic.Buffer
	pollAttachment netpoll.PollAttachment
	inboundBuffer  elastic.RingBuffer
	buffer         []byte
	cache          []byte
	isDatagram     bool
	opened         bool
	isEOF          bool
}

func newStreamConn(proto string, fd int, el *eventloop, sa unix.Sockaddr, localAddr, remoteAddr net.Addr) (c *conn) {
	_ = "STUB: not implemented"
	return nil
}

func newUDPConn(fd int, el *eventloop, localAddr net.Addr, sa unix.Sockaddr, connected bool) (c *conn) {
	_ = "STUB: not implemented"
	return nil
}

func (c *conn) release() { _ = "STUB: not implemented"; return }

func (c *conn) open(buf []byte) error { _ = "STUB: not implemented"; return nil }

func (c *conn) write(data []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (c *conn) writev(bs [][]byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

type asyncWriteHook struct {
	callback AsyncCallback
	data     []byte
}

func (c *conn) asyncWrite(a any) (err error) { _ = "STUB: not implemented"; return nil }

type asyncWritevHook struct {
	callback AsyncCallback
	data     [][]byte
}

func (c *conn) asyncWritev(a any) (err error) { _ = "STUB: not implemented"; return nil }

func (c *conn) sendTo(buf []byte, addr unix.Sockaddr) (n int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *conn) resetBuffer() { _ = "STUB: not implemented"; return }

func (c *conn) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (c *conn) Next(n int) (buf []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (c *conn) Peek(n int) (buf []byte, err error) { _ = "STUB: not implemented"; return nil, nil }

func (c *conn) Discard(n int) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *conn) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *conn) SendTo(p []byte, addr net.Addr) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (c *conn) Writev(bs [][]byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *conn) ReadFrom(r io.Reader) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (c *conn) WriteTo(w io.Writer) (n int64, err error) { _ = "STUB: not implemented"; return 0, nil }

func (c *conn) Flush() error { _ = "STUB: not implemented"; return nil }

func (c *conn) InboundBuffered() int { _ = "STUB: not implemented"; return 0 }

func (c *conn) OutboundBuffered() int { _ = "STUB: not implemented"; return 0 }

func (c *conn) Context() any         { _ = "STUB: not implemented"; return *new(any) }
func (c *conn) SetContext(ctx any)   { _ = "STUB: not implemented"; return }
func (c *conn) LocalAddr() net.Addr  { _ = "STUB: not implemented"; return *new(net.Addr) }
func (c *conn) RemoteAddr() net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func (c *conn) Fd() int                        { _ = "STUB: not implemented"; return 0 }
func (c *conn) Dup() (fd int, err error)       { _ = "STUB: not implemented"; return 0, nil }
func (c *conn) SetReadBuffer(bytes int) error  { _ = "STUB: not implemented"; return nil }
func (c *conn) SetWriteBuffer(bytes int) error { _ = "STUB: not implemented"; return nil }
func (c *conn) SetLinger(sec int) error        { _ = "STUB: not implemented"; return nil }
func (c *conn) SetNoDelay(noDelay bool) error  { _ = "STUB: not implemented"; return nil }

func (c *conn) SetKeepAlivePeriod(d time.Duration) error { _ = "STUB: not implemented"; return nil }

func (c *conn) SetKeepAlive(enabled bool, idle, intvl time.Duration, cnt int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *conn) AsyncWrite(buf []byte, callback AsyncCallback) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *conn) AsyncWritev(bs [][]byte, callback AsyncCallback) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *conn) Wake(callback AsyncCallback) error { _ = "STUB: not implemented"; return nil }

func (c *conn) CloseWithCallback(callback AsyncCallback) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *conn) Close() error { _ = "STUB: not implemented"; return nil }

func (c *conn) EventLoop() EventLoop { _ = "STUB: not implemented"; return *new(EventLoop) }

func (*conn) SetDeadline(_ time.Time) error { _ = "STUB: not implemented"; return nil }

func (*conn) SetReadDeadline(_ time.Time) error { _ = "STUB: not implemented"; return nil }

func (*conn) SetWriteDeadline(_ time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *conn) SafeContext() (ctx any) { _ = "STUB: not implemented"; return *new(any) }

func (c *conn) SetSafeContext(ctx any) { _ = "STUB: not implemented"; return }
