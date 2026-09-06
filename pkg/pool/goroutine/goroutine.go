package goroutine

import (
	"time"

	"github.com/panjf2000/ants/v2"

	"github.com/panjf2000/gnet/v2/pkg/logging"
)

const (
	DefaultAntsPoolSize = 1 << 18

	ExpiryDuration = 10 * time.Second

	Nonblocking = true
)

func init() {

	ants.Release()
}

var DefaultWorkerPool = Default()

type Pool = ants.Pool

type antsLogger struct {
	logging.Logger
}

func (l antsLogger) Printf(format string, args ...any) { _ = "STUB: not implemented"; return }

func Default() *Pool { _ = "STUB: not implemented"; return nil }
