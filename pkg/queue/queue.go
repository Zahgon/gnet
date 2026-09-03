package queue

import "sync"

type Func func(any) error

type Task struct {
	Exec  Func
	Param any
}

var taskPool = sync.Pool{New: func() any { return new(Task) }}

func GetTask() *Task { _ = "STUB: not implemented"; return nil }

func PutTask(task *Task) { _ = "STUB: not implemented"; return }

type AsyncTaskQueue interface {
	Enqueue(*Task)
	Dequeue() *Task
	IsEmpty() bool
	Length() int32
}

type EventPriority int

const (
	HighPriority EventPriority = iota

	LowPriority
)
