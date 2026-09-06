//go:build darwin || dragonfly || freebsd || openbsd

package netpoll

type IOFlags = uint16

type IOEvent = int16
