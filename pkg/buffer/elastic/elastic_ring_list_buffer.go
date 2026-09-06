package elastic

import (
	"io"

	"github.com/panjf2000/gnet/v2/pkg/buffer/linkedlist"
)

type Buffer struct {
	maxStaticBytes int
	ringBuffer     RingBuffer
	listBuffer     linkedlist.Buffer
}

func New(maxStaticBytes int) (*Buffer, error) { _ = "STUB: not implemented"; return nil, nil }

func (mb *Buffer) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (mb *Buffer) Peek(n int) ([][]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (mb *Buffer) Discard(n int) (discarded int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (mb *Buffer) Write(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (mb *Buffer) Writev(bs [][]byte) (int, error) { _ = "STUB: not implemented"; return 0, nil }

func (mb *Buffer) ReadFrom(r io.Reader) (int64, error) { _ = "STUB: not implemented"; return 0, nil }

func (mb *Buffer) WriteTo(w io.Writer) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (mb *Buffer) Buffered() int { _ = "STUB: not implemented"; return 0 }

func (mb *Buffer) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (mb *Buffer) Reset(maxStaticBytes int) { _ = "STUB: not implemented"; return }

func (mb *Buffer) Release() { _ = "STUB: not implemented"; return }
