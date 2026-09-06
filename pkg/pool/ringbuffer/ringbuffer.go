package ringbuffer

import (
	"sync"

	"github.com/panjf2000/gnet/v2/pkg/buffer/ring"
)

const (
	minBitSize = 6
	steps      = 20

	minSize = 1 << minBitSize

	calibrateCallsThreshold = 42000
	maxPercentile           = 0.95
)

type RingBuffer = ring.Buffer

type Pool struct {
	calls       [steps]uint64
	calibrating uint64

	defaultSize uint64
	maxSize     uint64

	pool sync.Pool
}

var builtinPool Pool

func Get() *RingBuffer { _ = "STUB: not implemented"; return nil }

func (p *Pool) Get() *RingBuffer { _ = "STUB: not implemented"; return nil }

func Put(b *RingBuffer) { _ = "STUB: not implemented"; return }

func (p *Pool) Put(b *RingBuffer) { _ = "STUB: not implemented"; return }

func (p *Pool) calibrate() { _ = "STUB: not implemented"; return }

type callSize struct {
	calls uint64
	size  uint64
}

type callSizes []callSize

func (ci callSizes) Len() int { _ = "STUB: not implemented"; return 0 }

func (ci callSizes) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (ci callSizes) Swap(i, j int) { _ = "STUB: not implemented"; return }

func index(n int) int { _ = "STUB: not implemented"; return 0 }
