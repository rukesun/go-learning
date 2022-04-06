package main

import (
	"fmt"
	"log"
	"net/http"
	_ "net/http/pprof"

	"github.com/rukesun/go-learning/examples/pprof/internal/simulator"
	_ "github.com/rukesun/go-learning/examples/pprof/internal/xxsignal"
)

var Simulators = []simulator.Simulator{}

func RunHandler(w http.ResponseWriter, req *http.Request) {
	log.Printf("handle run, url=%v", req.URL)
	for _, v := range Simulators {
		go v.Run()
	}
	fmt.Fprintf(w, "OK")
}

func main() {
	log.Printf("pprof demo is running")
	// init simulators.
	Simulators = append(Simulators, simulator.NewXXCPU())
	Simulators = append(Simulators, simulator.NewXXHeap())
	Simulators = append(Simulators, simulator.NewXXGoroutine())
	// init routers
	http.HandleFunc("/run", RunHandler)
	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}
