package main

import (
	qs "IstrazivackiRadQuickSort/quicksort"

	"flag"
	"fmt"
	"log"
	"os"
	"runtime"
	"syscall"
	"time"

	"os/signal"
	"runtime/debug"
	"runtime/pprof"
)

var (
	numbetOfCycles = 30
	err            error
	slice          qs.Slice

	piv qs.PivotPos
	slc qs.SliceOrd

	cpuProfileFlag = flag.String("cpuprofile", "", "Write cpu profile to file")
	memProfileFlag = flag.String("memprofile", "", "Write memory profile to file")

	pivotPosition qs.PivotPos
	sliceOrder    qs.SliceOrd
)

func runCycle(cycle int, slc qs.SliceOrd, pivot qs.PivotPos, timePerCel, memoryPerCel *string) {
	var memStatsBefore, memStatsAfter runtime.MemStats

	runtime.ReadMemStats(&memStatsBefore)
	startTime := time.Now()

	slice, err = qs.GenerateSlice(slc)
	if err != nil {
		fmt.Printf("[Cycle: %d] Error: %v\n", cycle, err)
		os.Exit(1)
	}

	err = slice.QuickSort(pivot)
	if err != nil {
		fmt.Printf("[Cycle: %d] Error: %v\n", cycle, err)
		os.Exit(1)
	}

	elapsed := time.Since(startTime).Milliseconds()
	runtime.ReadMemStats(&memStatsAfter)

	memDiff := int64(memStatsAfter.TotalAlloc/1024) - int64(memStatsBefore.TotalAlloc/1024)

	fmt.Printf("[%13s | %13s | Cycle: %2d] Time: %4v | Memory: %d KB\n", pivot.ToString(), slc.ToString(), cycle, elapsed, memDiff)

	*timePerCel += fmt.Sprintf("%v, ", elapsed)
	*memoryPerCel += fmt.Sprintf("%d, ", memDiff)

	// if resultsFile != nil {
	// 	resultsFile.WriteString(forFile)
	// }
}

func runTest(rf *os.File) {

	fmt.Println("Starting Quick sort analysis...")

	for piv = 0; piv < 4; piv++ {
		for slc = 0; slc < 4; slc++ {

			time := ""
			memory := ""

			for cycle := 1; cycle <= numbetOfCycles; cycle++ {
				runCycle(cycle, slc, piv, &time, &memory)
			}

			if rf != nil {
				rf.WriteString(fmt.Sprintf("[%s %s]: %s\n[%s %s]: %s\n", piv.ToString(), slc.ToString(), time[:len(time)-2], piv.ToString(), slc.ToString(), memory[:len(memory)-2]))
			}
		}
	}

	fmt.Println("Analysis complete! Check result.log for analysis data.")
}

func main() {

	flag.Parse()
	debug.SetGCPercent(-1)

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, os.Interrupt, syscall.SIGTERM)

	go func() {
		<-sigChan

		fmt.Println("Analysis interupted! Program force closed.")
		os.Exit(0)
	}()

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

	runTest(resultsFile)
}
