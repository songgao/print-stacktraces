package SIGUSR1

import (
	"os"
	"syscall"

	"github.com/songgao/stacktraces"
)

func init() {
	stacktraces.Set(os.Stderr, syscall.SIGUSR1)
}
