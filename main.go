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
	err   error
	slice qs.Slice

	cpuProfileFlag = flag.String("cpuprofile", "", "Write cpu profile to file")
	memProfileFlag = flag.String("memprofile", "", "Write memory profile to file")
	pivotFlag      = flag.String("pivot", "", "Enter the pivot position")
	sliceOrderFlag = flag.String("slice", "", "Determents the order state of the slice")
)

func runCycle(cycle int) {

	slice, err = qs.GenerateSlice(*sliceOrderFlag)
	if err != nil {
		fmt.Printf("[Cycle: %d] Error: %v", cycle, err)
		os.Exit(0)
	}

	err = slice.QuickSort(*pivotFlag)
	if err != nil {
		fmt.Printf("[Cycle: %d] Error: %v", cycle, err)
		os.Exit(0)
	}

	fmt.Printf("[Cycle: %d]\tSuccsesful cycle...\n", cycle)
}

func main() {

	flag.Parse()

	if *cpuProfileFlag != "" {
		cpuFile, err := os.Create(*cpuProfileFlag)
		if err != nil {
			log.Fatal(err)
		}

		pprof.StartCPUProfile(cpuFile)
		defer pprof.StopCPUProfile()
	}

	if *memProfileFlag != "" {
		memFile, err := os.Create(*memProfileFlag)
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
