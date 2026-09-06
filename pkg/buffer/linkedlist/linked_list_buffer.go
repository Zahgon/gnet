package linkedlist

import (
	"io"
)

type node struct {
	buf  []byte
	next *node
}

func (b *node) len() int { _ = "STUB: not implemented"; return 0 }

type Buffer struct {
	head  *node
	tail  *node
	size  int
	bytes int
}

func (llb *Buffer) Read(p []byte) (n int, err error) { _ = "STUB: not implemented"; return 0, nil }

func (llb *Buffer) AllocNode(n int) []byte { _ = "STUB: not implemented"; return nil }

func (llb *Buffer) FreeNode(p []byte) { _ = "STUB: not implemented"; return }

func (llb *Buffer) Append(p []byte) { _ = "STUB: not implemented"; return }

func (llb *Buffer) Pop() []byte { _ = "STUB: not implemented"; return nil }

func (llb *Buffer) PushFront(p []byte) { _ = "STUB: not implemented"; return }

func (llb *Buffer) PushBack(p []byte) { _ = "STUB: not implemented"; return }

func (llb *Buffer) Peek(maxBytes int) ([][]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (llb *Buffer) PeekWithBytes(maxBytes int, bs ...[]byte) ([][]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (llb *Buffer) Discard(n int) (discarded int, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

const minRead = 512

func (llb *Buffer) ReadFrom(r io.Reader) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (llb *Buffer) WriteTo(w io.Writer) (n int64, err error) {
	_ = "STUB: not implemented"
	return 0, nil
}

func (llb *Buffer) Len() int { _ = "STUB: not implemented"; return 0 }

func (llb *Buffer) Buffered() int { _ = "STUB: not implemented"; return 0 }

func (llb *Buffer) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (llb *Buffer) Reset() { _ = "STUB: not implemented"; return }

func (llb *Buffer) pop() *node { _ = "STUB: not implemented"; return nil }

func (llb *Buffer) pushFront(b *node) { _ = "STUB: not implemented"; return }

func (llb *Buffer) pushBack(b *node) { _ = "STUB: not implemented"; return }
