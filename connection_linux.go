package gnet

import (
	"github.com/panjf2000/gnet/v2/pkg/netpoll"
)

func (c *conn) processIO(_ int, ev netpoll.IOEvent, _ netpoll.IOFlags) error {
	_ = "STUB: not implemented"
	return nil
}
