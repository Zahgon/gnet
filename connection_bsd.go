//go:build darwin || dragonfly || freebsd || netbsd || openbsd

package gnet

import (
	"github.com/panjf2000/gnet/v2/pkg/netpoll"
)

func (c *conn) processIO(_ int, filter netpoll.IOEvent, flags netpoll.IOFlags) (err error) {
	_ = "STUB: not implemented"
	return nil
}
