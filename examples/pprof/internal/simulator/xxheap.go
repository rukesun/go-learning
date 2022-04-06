package simulator

import (
	"log"
)

type XXHeap struct {
	data [][1024 * 1024]byte
}

func NewXXHeap() *XXHeap {
	return &XXHeap{}
}

func (h *XXHeap) Name() string {
	return "XXHeap"
}

func (h *XXHeap) Run() {
	log.Printf("%v Run", h.Name())
	for i := 0; i < 10; i++ {
		h.data = append(h.data, [1024 * 1024]byte{})
	}
	log.Printf("After %v Run, len:%v, cap:%v", h.Name(), len(h.data), cap(h.data))
}
