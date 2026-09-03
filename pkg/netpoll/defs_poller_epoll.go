//go:build linux

package netpoll

import "golang.org/x/sys/unix"

type IOFlags = uint16

type IOEvent = uint32

const (
	InitPollEventsCap = 128

	MaxPollEventsCap = 1024

	MinPollEventsCap = 32

	MaxAsyncTasksAtOneTime = 256

	ReadEvents = unix.EPOLLIN | unix.EPOLLPRI

	WriteEvents = unix.EPOLLOUT

	ReadWriteEvents = ReadEvents | WriteEvents

	ErrEvents = unix.EPOLLERR | unix.EPOLLHUP
)

func IsReadEvent(event IOEvent) bool { _ = "STUB: not implemented"; return false }

func IsWriteEvent(event IOEvent) bool { _ = "STUB: not implemented"; return false }

func IsErrorEvent(event IOEvent, _ IOFlags) bool { _ = "STUB: not implemented"; return false }

type eventList struct {
	size   int
	events []epollevent
}

func newEventList(size int) *eventList { _ = "STUB: not implemented"; return nil }

func (el *eventList) expand() { _ = "STUB: not implemented"; return }

func (el *eventList) shrink() { _ = "STUB: not implemented"; return }
