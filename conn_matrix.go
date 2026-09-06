//go:build (darwin || dragonfly || freebsd || linux || netbsd || openbsd) && gc_opt

package gnet

import (
	"github.com/panjf2000/gnet/v2/internal/gfd"
)

type connMatrix struct {
	disableCompact bool
	connCounts     [gfd.ConnMatrixRowMax]int32
	row            int
	column         int
	table          [gfd.ConnMatrixRowMax][]*conn
	fd2gfd         map[int]gfd.GFD
}

func (cm *connMatrix) init() {
	cm.fd2gfd = make(map[int]gfd.GFD)
}

func (cm *connMatrix) iterate(f func(*conn) bool) { _ = "STUB: not implemented"; return }

func (cm *connMatrix) incCount(row int, delta int32) { _ = "STUB: not implemented"; return }

func (cm *connMatrix) loadCount() (n int32) { _ = "STUB: not implemented"; return 0 }

func (cm *connMatrix) addConn(c *conn, index int) { _ = "STUB: not implemented"; return }

func (cm *connMatrix) delConn(c *conn) { _ = "STUB: not implemented"; return }

func (cm *connMatrix) getConn(fd int) *conn { _ = "STUB: not implemented"; return nil }
