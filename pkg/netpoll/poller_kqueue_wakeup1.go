//go:build netbsd || openbsd

package netpoll

func (p *Poller) addWakeupEvent() error { _ = "STUB: not implemented"; return nil }

func (p *Poller) wakePoller() error { _ = "STUB: not implemented"; return nil }

func (p *Poller) drainWakeupEvent() { _ = "STUB: not implemented"; return }
