package main

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	type args struct {
		arr    []int
		search int
	}
	tests := []struct {
		name              string
		args              args
		wantPosition      int
		wantNumOperations int
	}{
		{"one arr", args{[]int{1}, 1}, 0, 1},
		{"empty arr", args{[]int{}, 1}, 0, 0},
		{"best case", args{[]int{1, 3, 6, 8, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}, 13}, 7, 1},
		{"worst case", args{[]int{1, 3, 6, 8, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21}, 1}, 0, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotPosition, gotNumOperations := BinarySearch(tt.args.arr, tt.args.search)
			if gotPosition != tt.wantPosition {
				t.Errorf("BinarySearch() gotPosition = %v, want %v", gotPosition, tt.wantPosition)
			}
			if gotNumOperations != tt.wantNumOperations {
				t.Errorf("BinarySearch() gotNumOperations = %v, want %v", gotNumOperations, tt.wantNumOperations)
			}
		})
	}
}
