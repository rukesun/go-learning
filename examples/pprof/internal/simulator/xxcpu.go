package simulator

import (
	"log"
)

type XXCPU struct {
}

func NewXXCPU() *XXCPU {
	return &XXCPU{}
}

func (c *XXCPU) Name() string {
	return "XXCPU"
}

func (c *XXCPU) Run() {
	log.Printf("%v Run", c.Name())
	loop := 1000000000
	for i := 0; i < loop; i++ {
		// do nothing
	}
}
