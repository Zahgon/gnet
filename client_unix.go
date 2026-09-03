//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd

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

func (cli *Client) Dial(network, address string) (Conn, error) {
	_ = "STUB: not implemented"
	return *new(Conn), nil
}

func (cli *Client) DialContext(network, address string, ctx any) (Conn, error) {
	_ = "STUB: not implemented"
	return *new(Conn), nil
}

func (cli *Client) Enroll(c net.Conn) (Conn, error) {
	_ = "STUB: not implemented"
	return *new(Conn), nil
}

func (cli *Client) EnrollContext(c net.Conn, ctx any) (Conn, error) {
	_ = "STUB: not implemented"
	//nolint:errcheck
	return *new(Conn), nil
}

//nolint:errcheck
