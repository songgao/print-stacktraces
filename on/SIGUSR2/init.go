package SIGUSR2

import (
	"os"
	"syscall"

	"github.com/songgao/print-stracktraces"
)

func init() {
	stacktraces.Set(os.Stderr, syscall.SIGUSR2)
}