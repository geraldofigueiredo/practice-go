package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	type args struct {
		arr    []int
		search int
		debug  bool
	}
	tests := []struct {
		name              string
		args              args
		wantPosition      int
		wantNumOperations int
	}{
		{"one arr", args{generateOrderedArr(1), 1, false}, 0, 1},
		{"empty arr", args{generateOrderedArr(0), 1, false}, 0, 0},
		{"best case", args{generateOrderedArr(15), 7, false}, 7, 1},
		{"worst case", args{generateOrderedArr(15), 0, false}, 0, 4},
		{"worst case large size", args{generateOrderedArr(1000000000), 0, false}, 0, 29},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPosition, gotNumOperations := BinarySearch(tt.args.arr, tt.args.search, tt.args.debug)
			if gotPosition != tt.wantPosition {
				t.Errorf("BinarySearch() gotPosition = %v, want %v", gotPosition, tt.wantPosition)
			}
			if gotNumOperations != tt.wantNumOperations {
				t.Errorf("BinarySearch() gotNumOperations = %v, want %v", gotNumOperations, tt.wantNumOperations)
			}
		})
	}
}

func generateOrderedArr(size int) []int {
	arr := make([]int, size)
	for i := range arr {
		arr[i] = i
	}

	return arr
}
