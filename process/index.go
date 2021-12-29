package process

import (
	"os"
	"os/signal"
	"syscall"
)

var C = make(chan os.Signal)

func init() {
	signal.Notify(C, syscall.SIGINT, syscall.SIGQUIT)
}
