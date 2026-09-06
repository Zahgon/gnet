package errors

import "errors"

var (
	ErrEmptyEngine = errors.New("gnet: the internal engine is empty")

	ErrEngineShutdown = errors.New("gnet: server is going to be shutdown")

	ErrEngineInShutdown = errors.New("gnet: server is already in shutdown")

	ErrAcceptSocket = errors.New("gnet: accept a new connection error")

	ErrTooManyEventLoopThreads = errors.New("gnet: too many event-loops under LockOSThread mode")

	ErrUnsupportedProtocol = errors.New("gnet: only unix, tcp/tcp4/tcp6, udp/udp4/udp6 are supported")

	ErrUnsupportedTCPProtocol = errors.New("gnet: only tcp/tcp4/tcp6 are supported")

	ErrUnsupportedUDPProtocol = errors.New("gnet: only udp/udp4/udp6 are supported")

	ErrUnsupportedUDSProtocol = errors.New("gnet: only unix is supported")

	ErrUnsupportedOp = errors.New("gnet: unsupported operation")

	ErrNegativeSize = errors.New("gnet: negative size is not allowed")

	ErrNoIPv4AddressOnInterface = errors.New("gnet: no IPv4 address on interface")

	ErrInvalidNetworkAddress = errors.New("gnet: invalid network address")

	ErrInvalidNetConn = errors.New("gnet: the net.Conn is empty")

	ErrNilRunnable = errors.New("gnet: nil runnable is not allowed")
)
