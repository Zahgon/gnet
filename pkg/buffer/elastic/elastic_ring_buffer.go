package elastic

import (
	"io"

	"github.com/panjf2000/gnet/v2/pkg/buffer/ring"
)

type RingBuffer struct {
	rb *ring.Buffer
}

func (b *RingBuffer) instance() *ring.Buffer { _ = "STUB: not implemented"; return nil }

func (b *RingBuffer) Done() { _ = "STUB: not implemented"; return }

func (b *RingBuffer) done() { _ = "STUB: not implemented"; return }

func (b *RingBuffer) Peek(n int) (head []byte, tail []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (b *RingBuffer) Discard(n int) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (b *RingBuffer) Read(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (b *RingBuffer) ReadByte() (byte, error) { _ = "STUB: not implemented"; return 0, nil }

func (b *RingBuffer) Write(p []byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (b *RingBuffer) WriteByte(c byte) error { _ = "STUB: not implemented"; return nil }

func (b *RingBuffer) Buffered() int { _ = "STUB: not implemented"; return 0 }

func (b *RingBuffer) Len() int { _ = "STUB: not implemented"; return 0 }

func (b *RingBuffer) Cap() int { _ = "STUB: not implemented"; return 0 }

func (b *RingBuffer) Available() int { _ = "STUB: not implemented"; return 0 }

func (b *RingBuffer) WriteString(s string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (b *RingBuffer) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (b *RingBuffer) ReadFrom(r io.Reader) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (b *RingBuffer) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (b *RingBuffer) IsFull() bool { _ = "STUB: not implemented"; return false }

func (b *RingBuffer) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (b *RingBuffer) Reset() { _ = "STUB: not implemented"; return }
