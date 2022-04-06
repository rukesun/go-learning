package xxsignal

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/rukesun/go-learning/examples/pprof/internal/xxprof"
)

func init() {
	go func() {
		var profiler *xxprof.Profile

		signals := make(chan os.Signal, 1)
		signal.Notify(signals, syscall.SIGUSR1, syscall.SIGUSR2)

		for {
			v := <-signals
			log.Printf("Got signal:", v)
			switch v {
			case syscall.SIGUSR1:
			case syscall.SIGUSR2:
				if profiler == nil {
					profiler = xxprof.NewProfile()
					profiler.Start()
				} else {
					profiler.Stop()
					profiler = nil
				}
			default:
				log.Printf("Got unregistered signal:", v)
			}
		}
	}()
}
