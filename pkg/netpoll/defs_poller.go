//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd

package netpoll

type PollEventHandler func(int, IOEvent, IOFlags) error

type PollAttachment struct {
	FD       int
	Callback PollEventHandler
}
