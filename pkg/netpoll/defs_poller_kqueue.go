//go:build darwin || dragonfly || freebsd || netbsd || openbsd

package netpoll

import "golang.org/x/sys/unix"

const (
	InitPollEventsCap = 64

	MaxPollEventsCap = 512

	MinPollEventsCap = 16

	MaxAsyncTasksAtOneTime = 128

	ReadEvents = unix.EVFILT_READ

	WriteEvents = unix.EVFILT_WRITE

	ReadWriteEvents = ReadEvents | WriteEvents

	ErrEvents = unix.EV_EOF | unix.EV_ERROR
)

func IsReadEvent(event IOEvent) bool { _ = "STUB: not implemented"; return false }

func IsWriteEvent(event IOEvent) bool { _ = "STUB: not implemented"; return false }

func IsErrorEvent(_ IOEvent, flags IOFlags) bool { _ = "STUB: not implemented"; return false }

type eventList struct {
	size   int
	events []unix.Kevent_t
}

func newEventList(size int) *eventList { _ = "STUB: not implemented"; return nil }

func (el *eventList) expand() { _ = "STUB: not implemented"; return }

func (el *eventList) shrink() { _ = "STUB: not implemented"; return }
