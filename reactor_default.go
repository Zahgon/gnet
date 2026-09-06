//go:build (darwin || dragonfly || freebsd || linux || netbsd || openbsd) && !poll_opt

package gnet

func (el *eventloop) rotate() error { _ = "STUB: not implemented"; return nil }

func (el *eventloop) orbit() error { _ = "STUB: not implemented"; return nil }

func (el *eventloop) run() error { _ = "STUB: not implemented"; return nil }
