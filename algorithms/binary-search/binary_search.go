package main

import "fmt"

func BinarySearch(arr []int, search int) (position, numOperations int) {
	low := 0
	high := len(arr) - 1
	numOperations = 0

	for low <= high {
		numOperations++
		mid := (low + high) / 2
		guess := arr[mid]
		VisualizeBinarySearch(arr, low, high, mid)
		if guess == search {
			return mid, numOperations
		}
		if guess > search {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	return 0, numOperations
}

const (
	colorReset  = "\033[0m"
	colorRed    = "\033[31m"
	colorGreen  = "\033[32m"
	colorYellow = "\033[33m"
	colorBlue   = "\033[34m"
	colorPurple = "\033[35m"
	colorCyan   = "\033[36m"
	colorWhite  = "\033[37m"
)

func VisualizeBinarySearch(arr []int, low, high, mid int) {
	for i, el := range arr {
		if i == low {
			fmt.Print(string(colorBlue), el, string(colorReset), " ")
			continue
		}
		if i == high {
			fmt.Print(string(colorCyan), el, string(colorReset), " ")
			continue
		}
		if i == mid {
			fmt.Print(string(colorRed), el, string(colorReset), " ")
			continue
		}
		fmt.Print(el, " ")
	}
	fmt.Println("")
}
