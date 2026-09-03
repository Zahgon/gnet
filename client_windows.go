package gnet

import (
	"net"
)

type Client struct {
	opts *Options
	eng  *engine
}

func NewClient(eh EventHandler, opts ...Option) (cli *Client, err error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (cli *Client) Start() error { _ = "STUB: not implemented"; return nil }

func (cli *Client) Stop() error { _ = "STUB: not implemented"; return nil }

func (cli *Client) Dial(network, addr string) (Conn, error) {
	_ = "STUB: not implemented"
	return *new(Conn), nil
}

func (cli *Client) DialContext(network, addr string, ctx any) (Conn, error) {
	_ = "STUB: not implemented"
	return *new(Conn), nil
}

func (cli *Client) Enroll(nc net.Conn) (gc Conn, err error) {
	_ = "STUB: not implemented"
	return *new(Conn), nil
}

func (cli *Client) EnrollContext(nc net.Conn, ctx any) (gc Conn, err error) {
	_ = "STUB: not implemented"
	return *new(Conn), nil
}
