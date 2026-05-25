package main

import (
	qs "IstrazivackiRadQuickSort/quicksort"
	"flag"
	"fmt"
	"log"
	"os"
	"runtime/pprof"
)

var (
	err        error
	slice      qs.Slice
	cpuProfile = flag.String("cpuprofile", "", "write cpu profile to file")
)

func runCycle() {

	slice.GenerateSlice()
	err = slice.QuickSort(qs.PivotLast)
	if err != nil {
		fmt.Println(err)
	}
}

func main() {

	flag.Parse()
	if *cpuProfile != "" {
		file, err := os.Create(*cpuProfile)
		if err != nil {
			log.Fatal(err)
		}

		pprof.StartCPUProfile(file)
		defer pprof.StopCPUProfile()
	}

}
