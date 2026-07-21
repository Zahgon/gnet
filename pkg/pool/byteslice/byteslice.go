package byteslice

import (
	"sync"
)

var builtinPool Pool

type Pool struct {
	pools [32]sync.Pool
}

func Get(size int) []byte { _ = "STUB: not implemented"; return nil }

func Put(buf []byte) { _ = "STUB: not implemented"; return }

func (p *Pool) Get(size int) []byte { _ = "STUB: not implemented"; return nil }

func (p *Pool) Put(buf []byte) { _ = "STUB: not implemented"; return }

func index(n uint32) uint32 { _ = "STUB: not implemented"; return 0 }
