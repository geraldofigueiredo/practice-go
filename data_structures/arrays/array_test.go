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
		{"capacity multiplied by growth rate (32)", args{capacity: 17}, vector{capacity: 32, size: 0, data: make([]int, 32)}, false},
		{"capacity multiplied by growth rate (64)", args{capacity: 33}, vector{capacity: 64, size: 0, data: make([]int, 32)}, false},
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

func Test_vector_GetValueAt(t *testing.T) {
	defaultVector := vector{
		size:     10,
		capacity: 16,
		data:     []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
	}

	type fields struct {
		size     int
		capacity int
		data     []int
	}
	type args struct {
		index int
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    int
		wantErr bool
	}{
		{name: "negative index", fields: fields(defaultVector), args: args{-1}, want: 0, wantErr: true},
		{name: "index larger than size", fields: fields(defaultVector), args: args{defaultVector.size}, want: 0, wantErr: true},
		{name: "valid index", fields: fields(defaultVector), args: args{5}, want: 5, wantErr: false},
		{name: "first index", fields: fields(defaultVector), args: args{0}, want: 0, wantErr: false},
		{name: "last index", fields: fields(defaultVector), args: args{defaultVector.size - 1}, want: 9, wantErr: false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			v := vector{
				size:     tt.fields.size,
				capacity: tt.fields.capacity,
				data:     tt.fields.data,
			}
			got, err := v.GetValueAt(tt.args.index)
			if (err != nil) != tt.wantErr {
				t.Errorf("vector.GetValueAt() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("vector.GetValueAt() = %v, want %v", got, tt.want)
			}
		})
	}
}
