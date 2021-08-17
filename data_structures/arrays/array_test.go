package arrays

import (
	"reflect"
	"testing"
)

func TestNewVector(t *testing.T) {
	type args struct {
		capacity int
	}
	tests := []struct {
		name    string
		args    args
		want    VectorInterface
		wantErr bool
	}{
		{"invalid capacity", args{capacity: 0}, vector{}, true},
		{"min capacity", args{capacity: 1}, vector{capacity: 16, size: 0, data: make([]int, 16)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewVector(tt.args.capacity)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewVector() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewVector() = %v, want %v", got, tt.want)
			}
		})
	}
}
