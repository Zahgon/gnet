//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd

package gnet

import (
	"github.com/panjf2000/gnet/v2/pkg/netpoll"
)

func (el *eventloop) accept0(fd int, _ netpoll.IOEvent, _ netpoll.IOFlags) error {
	_ = "STUB: not implemented"
	return nil
}

func (el *eventloop) accept(fd int, ev netpoll.IOEvent, flags netpoll.IOFlags) error {
	_ = "STUB: not implemented"
	return nil
}
