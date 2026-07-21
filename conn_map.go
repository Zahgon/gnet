//go:build (darwin || dragonfly || freebsd || linux || netbsd || openbsd) && !gc_opt

package gnet

type connMatrix struct {
	connCount int32
	connMap   map[int]*conn
}

func (cm *connMatrix) init() {
	cm.connMap = make(map[int]*conn)
}

func (cm *connMatrix) iterate(f func(*conn) bool) { _ = "STUB: not implemented"; return }

func (cm *connMatrix) incCount(_ int, delta int32) { _ = "STUB: not implemented"; return }

func (cm *connMatrix) loadCount() (n int32) { _ = "STUB: not implemented"; return 0 }

func (cm *connMatrix) addConn(c *conn, index int) { _ = "STUB: not implemented"; return }

func (cm *connMatrix) delConn(c *conn) { _ = "STUB: not implemented"; return }

func (cm *connMatrix) getConn(fd int) *conn { _ = "STUB: not implemented"; return nil }
