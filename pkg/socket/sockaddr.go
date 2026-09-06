//go:build darwin || dragonfly || freebsd || linux || netbsd || openbsd

package socket

import (
	"net"

	"golang.org/x/sys/unix"
)

func NetAddrToSockaddr(addr net.Addr) unix.Sockaddr {
	_ = "STUB: not implemented"
	return *new(unix.Sockaddr)
}

func IPAddrToSockaddr(addr *net.IPAddr) unix.Sockaddr {
	_ = "STUB: not implemented"
	return *new(unix.Sockaddr)
}

func TCPAddrToSockaddr(addr *net.TCPAddr) unix.Sockaddr {
	_ = "STUB: not implemented"
	return *new(unix.Sockaddr)
}

func UDPAddrToSockaddr(addr *net.UDPAddr) unix.Sockaddr {
	_ = "STUB: not implemented"
	return *new(unix.Sockaddr)
}

func IPToSockaddr(ip net.IP, port int, zone string) unix.Sockaddr {
	_ = "STUB: not implemented"
	return *new(unix.Sockaddr)
}

func UnixAddrToSockaddr(addr *net.UnixAddr) (unix.Sockaddr, int) {
	_ = "STUB: not implemented"
	return *new(unix.Sockaddr), 0
}

func SockaddrToTCPOrUnixAddr(sa unix.Sockaddr) net.Addr {
	_ = "STUB: not implemented"
	return *new(net.Addr)
}

func SockaddrToUDPAddr(sa unix.Sockaddr) net.Addr { _ = "STUB: not implemented"; return *new(net.Addr) }

func ip6ZoneToInt(zone string) int { _ = "STUB: not implemented"; return 0 }

func ip6ZoneToString(zone uint32) string { _ = "STUB: not implemented"; return "" }

func itod(v uint) string { _ = "STUB: not implemented"; return "" }

const big = 0xFFFFFF

func dtoi(s string, i0 int) (n int, i int, ok bool) { _ = "STUB: not implemented"; return 0, 0, false }
