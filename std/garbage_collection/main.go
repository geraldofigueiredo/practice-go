package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	var mem runtime.MemStats
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 500000000)
		if s == nil {
			fmt.Println("Operation Failed!")
		}
	}
	printStats(mem)

	for i := 0; i < 10; i++ {
		s := make([]byte, 500000000)
		if s == nil {
			fmt.Println("Operation Failed!")
		}
		time.Sleep(time.Second * 5)
	}
	printStats(mem)
}

func printStats(mem runtime.MemStats) {
	runtime.ReadMemStats(&mem)
	fmt.Println("mem.Alloc:", mem.Alloc)
	fmt.Println("mem.TotalAlloc:", mem.TotalAlloc)
	fmt.Println("mem.HeapAlloc:", mem.HeapAlloc)
	fmt.Println("mem.NumGC:", mem.NumGC)
	fmt.Println("---")
}
