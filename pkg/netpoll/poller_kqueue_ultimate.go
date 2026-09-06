//go:build (darwin || dragonfly || freebsd || netbsd || openbsd) && poll_opt

package netpoll

import (
	"github.com/panjf2000/gnet/v2/pkg/queue"
)

type Poller struct {
	fd                          int
	pipe                        []int
	wakeupCall                  int32
	asyncTaskQueue              queue.AsyncTaskQueue
	urgentAsyncTaskQueue        queue.AsyncTaskQueue
	highPriorityEventsThreshold int32
}

func OpenPoller() (poller *Poller, err error) { _ = "STUB: not implemented"; return nil, nil }

func (p *Poller) Close() error { _ = "STUB: not implemented"; return nil }

func (p *Poller) Trigger(priority queue.EventPriority, fn queue.Func, param any) (err error) {
	_ = "STUB: not implemented"
	return nil
}

func (p *Poller) Polling() error { _ = "STUB: not implemented"; return nil }

func (p *Poller) AddReadWrite(pa *PollAttachment, edgeTriggered bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Poller) AddRead(pa *PollAttachment, edgeTriggered bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Poller) AddWrite(pa *PollAttachment, edgeTriggered bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Poller) ModRead(pa *PollAttachment, _ bool) error { _ = "STUB: not implemented"; return nil }

func (p *Poller) ModReadWrite(pa *PollAttachment, edgeTriggered bool) error {
	_ = "STUB: not implemented"
	return nil
}

func (p *Poller) Delete(_ int) error { _ = "STUB: not implemented"; return nil }
