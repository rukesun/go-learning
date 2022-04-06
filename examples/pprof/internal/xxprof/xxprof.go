/*
It's just an example, thread safety is not guaranteed.
*/
package xxprof

import (
	"fmt"
	"log"
	"os"
	"path"
	"runtime/pprof"
	"syscall"
	"time"
)

const (
	timeFormat = "0102150405"
)

type Profile struct {
	cpuFile  *os.File
	heapFile *os.File
}

func NewProfile() *Profile {
	return &Profile{}
}

func createDumpFile(kind string) string {
	command := path.Base(os.Args[0])
	pid := syscall.Getpid()
	return path.Join(os.TempDir(), fmt.Sprintf("%s-%d-%s-%s.pprof", command, pid, kind, time.Now().Format(timeFormat)))
}

func (p *Profile) startCpuProfile() {
	fn := createDumpFile("cpu")
	f, err := os.Create(fn)
	if err != nil {
		log.Printf("profile: could not create cpu profile %v, error: %v", fn, err)
		return
	}
	p.cpuFile = f

	log.Printf("profile: start cpu profile, file=%v", fn)
	pprof.StartCPUProfile(f)
}

func (p *Profile) stopCpuProfile() {
	log.Printf("profile: stop cpu profile")
	pprof.StopCPUProfile()
	if p.cpuFile != nil {
		p.cpuFile.Close()
		p.cpuFile = nil
	}
}

func (p *Profile) startHeapProfile() {
	fn := createDumpFile("heap")
	f, err := os.Create(fn)
	if err != nil {
		log.Printf("profile: could not create heap profile %v, error: %v", fn, err)
		return
	}
	p.heapFile = f

	log.Printf("profile: start heap profile, file=%v", fn)
}

func (p *Profile) stopHeapProfile() {
	log.Printf("profile: stop heap profile")
	if p.heapFile != nil {
		pprof.WriteHeapProfile(p.heapFile)
		p.heapFile.Close()
		p.heapFile = nil
	}
}

func (p *Profile) Start() {
	p.startCpuProfile()
	p.startHeapProfile()
}

func (p *Profile) Stop() {
	p.stopCpuProfile()
	p.stopHeapProfile()
}
