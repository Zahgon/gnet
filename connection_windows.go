package gnet

import (
	"io"
	"net"
	"sync/atomic"
	"time"

	"github.com/panjf2000/gnet/v2/pkg/buffer/elastic"
	bbPool "github.com/panjf2000/gnet/v2/pkg/pool/bytebuffer"
)

type netErr struct {
	c   *conn
	err error
}

type tcpConn struct {
	c *conn
	b *bbPool.ByteBuffer
}

type udpConn struct {
	c *conn
}

type openConn struct {
	c  *conn
	cb func()
}

type conn struct {
	pc            net.PacketConn
	ctx           any
	safeCtx       atomic.Pointer[any]
	loop          *eventloop
	buffer        *bbPool.ByteBuffer
	cache         []byte
	rawConn       net.Conn
	localAddr     net.Addr
	remoteAddr    net.Addr
	inboundBuffer elastic.RingBuffer
}

func packTCPConn(c *conn, buf []byte) *tcpConn { _ = "STUB: not implemented"; return nil }

func unpackTCPConn(tc *tcpConn) *conn { _ = "STUB: not implemented"; return nil }

func packUDPConn(c *conn, buf []byte) *udpConn { _ = "STUB: not implemented"; return nil }

func newStreamConn(el *eventloop, nc net.Conn, ctx any) (c *conn) {
	_ = "STUB: not implemented"
	return nil
}

func (c *conn) release() { _ = "STUB: not implemented"; return }

func newUDPConn(el *eventloop, pc net.PacketConn, rc net.Conn, localAddr, remoteAddr net.Addr, ctx any) *conn {
	_ = "STUB: not implemented"
	return nil
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

func (c *conn) Fd() (fd int) { _ = "STUB: not implemented"; return 0 }

func (c *conn) Dup() (fd int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (c *conn) SetReadBuffer(bytes int) error { _ = "STUB: not implemented"; return nil }

func (c *conn) SetWriteBuffer(bytes int) error { _ = "STUB: not implemented"; return nil }

func (c *conn) SetLinger(sec int) error { _ = "STUB: not implemented"; return nil }

func (c *conn) SetNoDelay(noDelay bool) error { _ = "STUB: not implemented"; return nil }

func (c *conn) SetKeepAlivePeriod(d time.Duration) error { _ = "STUB: not implemented"; return nil }

func (c *conn) SetKeepAlive(enabled bool, idle, intvl time.Duration, cnt int) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *conn) AsyncWrite(buf []byte, cb AsyncCallback) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *conn) AsyncWritev(bs [][]byte, cb AsyncCallback) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *conn) Wake(cb AsyncCallback) (err error) { _ = "STUB: not implemented"; return nil }

func (c *conn) Close() (err error) { _ = "STUB: not implemented"; return nil }

func (c *conn) CloseWithCallback(cb AsyncCallback) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (c *conn) EventLoop() EventLoop { _ = "STUB: not implemented"; return *new(EventLoop) }

func (*conn) SetDeadline(_ time.Time) error { _ = "STUB: not implemented"; return nil }

func (*conn) SetReadDeadline(_ time.Time) error { _ = "STUB: not implemented"; return nil }

func (*conn) SetWriteDeadline(_ time.Time) error { _ = "STUB: not implemented"; return nil }

func (c *conn) SafeContext() (ctx any) { _ = "STUB: not implemented"; return *new(any) }

func (c *conn) SetSafeContext(ctx any) { _ = "STUB: not implemented"; return }
