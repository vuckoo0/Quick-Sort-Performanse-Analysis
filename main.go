package main

import (
	qs "IstrazivackiRadQuickSort/quicksort"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

const (
	numbetOfCycles = 1
)

var (
	err        error
	slice      qs.Slice
	cpuProfile = flag.String("cpuprofile", "", "Write cpu profile to file")
	memProfile = flag.String("memprofile", "", "Write memory profile to file")
)

func runCycle(cycle int) {

	slice = qs.GenerateSlice()
	err = slice.QuickSort(qs.PivotLast)
	if err != nil {
		fmt.Printf("[Cycle: %d] Error: %v", cycle, err)
		os.Exit(0)
	}

	fmt.Printf("[Cycle: %d]\tSuccsesful cycle...\n", cycle)
}

func main() {

	flag.Parse()

	if *cpuProfile != "" {
		cpuFile, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal(err)
		}

		pprof.StartCPUProfile(cpuFile)
		defer pprof.StopCPUProfile()
	}

	if *memProfile != "" {
		memFile, err := os.Create(*memProfile)
		if err != nil {
			log.Fatal(err)
		}

		pprof.WriteHeapProfile(memFile)
		defer memFile.Close()
	}

	cycle := 0
	for cycle < numbetOfCycles {
		runCycle(cycle)
		cycle++
	}
}
