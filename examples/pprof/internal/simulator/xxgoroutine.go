package simulator

import (
	"log"
	"time"
)

type XXGoroutine struct {
}

func NewXXGoroutine() *XXGoroutine {
	return &XXGoroutine{}
}

func (g *XXGoroutine) Name() string {
	return "XXGoroutine"
}

func (g *XXGoroutine) Run() {
	log.Printf("%v Run", g.Name())
	for i := 0; i < 1000; i++ {
		go func() {
			time.Sleep(3600 * time.Second)
		}()
	}
}
