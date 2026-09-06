package ring

import (
	"errors"
	"io"
)

const (
	MinRead = 512

	DefaultBufferSize   = 1024
	bufferGrowThreshold = 4 * 1024
)

var ErrIsEmpty = errors.New("ring-buffer is empty")

type Buffer struct {
	buf     []byte
	size    int
	r       int
	w       int
	isEmpty bool
}

func New(size int) *Buffer { _ = "STUB: not implemented"; return nil }

func (rb *Buffer) Peek(n int) (head []byte, tail []byte) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (rb *Buffer) peekAll() (head []byte, tail []byte) { _ = "STUB: not implemented"; return nil, nil }

func (rb *Buffer) Discard(n int) (discarded int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (rb *Buffer) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (rb *Buffer) ReadByte() (b byte, err error) { _ = "STUB: not implemented"; return 0, nil }

func (rb *Buffer) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (rb *Buffer) WriteByte(c byte) error { _ = "STUB: not implemented"; return nil }

func (rb *Buffer) Buffered() int { _ = "STUB: not implemented"; return 0 }

func (rb *Buffer) Len() int { _ = "STUB: not implemented"; return 0 }

func (rb *Buffer) Cap() int { _ = "STUB: not implemented"; return 0 }

func (rb *Buffer) Available() int { _ = "STUB: not implemented"; return 0 }

func (rb *Buffer) WriteString(s string) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (rb *Buffer) Bytes() []byte { _ = "STUB: not implemented"; return nil }

func (rb *Buffer) ReadFrom(r io.Reader) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (rb *Buffer) WriteTo(w io.Writer) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (rb *Buffer) IsFull() bool { _ = "STUB: not implemented"; return false }

func (rb *Buffer) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (rb *Buffer) Reset() { _ = "STUB: not implemented"; return }

func (rb *Buffer) grow(newCap int) { _ = "STUB: not implemented"; return }
