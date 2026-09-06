//go:build (darwin || dragonfly || freebsd || linux || netbsd || openbsd) && poll_opt

package netpoll

import "unsafe"

func convertPollAttachment(ptr unsafe.Pointer, attachment *PollAttachment) {
	_ = "STUB: not implemented"
	return
}

func restorePollAttachment(ptr unsafe.Pointer) *PollAttachment {
	_ = "STUB: not implemented"
	return nil
}
