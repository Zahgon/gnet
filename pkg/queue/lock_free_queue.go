package queue

import (
	"unsafe"
)

type lockFreeQueue struct {
	head   unsafe.Pointer
	tail   unsafe.Pointer
	length int32
}

type node struct {
	value *Task
	next  unsafe.Pointer
}

func NewLockFreeQueue() AsyncTaskQueue { _ = "STUB: not implemented"; return *new(AsyncTaskQueue) }

func (q *lockFreeQueue) Enqueue(task *Task) { _ = "STUB: not implemented"; return }

func (q *lockFreeQueue) Dequeue() *Task { _ = "STUB: not implemented"; return nil }

func (q *lockFreeQueue) IsEmpty() bool { _ = "STUB: not implemented"; return false }

func (q *lockFreeQueue) Length() int32 { _ = "STUB: not implemented"; return 0 }

func load(p *unsafe.Pointer) (n *node) { _ = "STUB: not implemented"; return nil }

func cas(p *unsafe.Pointer, old, new *node) bool {
	_ = "STUB: not implemented" //nolint:revive
	return false
}
