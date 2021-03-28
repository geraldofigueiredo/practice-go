package main

import (
	"fmt"
	"io/ioutil"
	"sync"
	"time"
)

var bytes []int

func bytesCount(bytesSlice []byte) {
	for _, byte := range bytesSlice {
		bytes[byte]++
	}
}

func bytesCountConcurrent(wg *sync.WaitGroup, bytesSlice []byte) {
	defer wg.Done()
	for _, byte := range bytesSlice {
		bytes[byte]++
	}
}

func concurrentBytesCount(bytesSlice []byte) {
	var wg sync.WaitGroup
	slices := 50
	sliceLen := len(bytesSlice) / slices

	for i := 0; i < slices; i++ {
		wg.Add(1)
		go bytesCountConcurrent(&wg, bytesSlice[sliceLen*i:sliceLen*(i+1)])
	}
	wg.Add(1)
	go bytesCountConcurrent(&wg, bytesSlice[sliceLen*slices:])

	wg.Wait()
}

func main() {
	bytes = make([]int, 256)

	file, err := ioutil.ReadFile("picture.jpg")
	if err != nil {
		panic(err)
	}

	start := time.Now()
	bytesCount(file)
	elapsed := time.Since(start)

	fmt.Println(bytes, "\n\nelapsed time: ", fmt.Sprintf("without goroutine took: %s", elapsed))

	bytes = make([]int, 256)

	start = time.Now()
	concurrentBytesCount(file)
	elapsed = time.Since(start)

	fmt.Println(bytes, "\n\nelapsed time: ", fmt.Sprintf("with goroutines took: %s", elapsed))
}
