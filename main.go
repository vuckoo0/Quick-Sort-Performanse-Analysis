package main

import (
	qs "IstrazivackiRadQuickSort/quicksort"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"runtime/pprof"
	"time"
)

const (
	numbetOfCycles = 1
)

var (
	err   error
	slice qs.Slice

	cpuProfileFlag = flag.String("cpuprofile", "", "Write cpu profile to file")
	memProfileFlag = flag.String("memprofile", "", "Write memory profile to file")

	pivotPosition = "last"
	sliceOrder    = "block"
)

func runCycle(cycle int, resultsFile *os.File) {
	var memStatsBefore, memStatsAfter runtime.MemStats

	runtime.ReadMemStats(&memStatsBefore)
	startTime := time.Now()

	slice, err = qs.GenerateSlice(sliceOrder)
	if err != nil {
		fmt.Printf("[Cycle: %d] Error: %v\n", cycle, err)
		os.Exit(1)
	}

	err = slice.QuickSort(pivotPosition)
	if err != nil {
		fmt.Printf("[Cycle: %d] Error: %v\n", cycle, err)
		os.Exit(1)
	}

	elapsed := time.Since(startTime)
	runtime.ReadMemStats(&memStatsAfter)

	memBefore := memStatsBefore.Alloc / 1024
	memAfter := memStatsAfter.Alloc / 1024
	memDiff := int64(memAfter) - int64(memBefore)

	result := fmt.Sprintf("[Cycle: %2d] Time: %v | Memory Before: %d KB | Memory After: %d KB | Diff: %d KB\n",
		cycle, elapsed, memBefore, memAfter, memDiff)

	fmt.Print(result)
	if resultsFile != nil {
		resultsFile.WriteString(fmt.Sprintf("%v,", elapsed))
	}
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

	resultsFile, err := os.OpenFile("results.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer resultsFile.Close()

	header := fmt.Sprintf("Starting %d cycles with pivot='%s', sliceOrder='%s'\n\n", numbetOfCycles, pivotPosition, sliceOrder)
	fmt.Print(header)

	for cycle := 1; cycle <= numbetOfCycles; cycle++ {
		runCycle(cycle, resultsFile)
	}

	footer := fmt.Sprintf("\nCompleted %d cycles. Results saved to results.log\n", numbetOfCycles)
	fmt.Print(footer)
}
