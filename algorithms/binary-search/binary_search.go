package main

func BinarySearch(arr []int, search int) (position, numOperations int) {
	low := 0
	high := len(arr) - 1
	numOperations = 0

	for low <= high {
		numOperations++
		mid := (low + high) / 2
		guess := arr[mid]
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
