package linkedlist

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewEmptyNode(t *testing.T) {
	got := NewEmptyNode[int]()
	want := &Node[int]{}

	if equal := reflect.DeepEqual(got, want); !equal {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestNewNode(t *testing.T) {
	type someStruct struct {
		x int
		y int
	}
	strData := "some string"
	strNode := NewNode[string](strData, nil, nil)
	assert.IsType(t, "", strNode.data)
	assert.Equal(t, strData, strNode.GetValue())

	intData := 99
	intNode := NewNode[int](intData, nil, nil)
	assert.IsType(t, 0, intNode.data)
	assert.Equal(t, intData, intNode.GetValue())

	someStructData := someStruct{}
	customTypeNode := NewNode[someStruct](someStructData, nil, nil)
	assert.IsType(t, someStructData, customTypeNode.data)
}

func TestSetValue(t *testing.T) {
	tests := []struct {
		name         string
		elementValue int64
		result       int64
	}{
		{"set value: 0", 0, 0},
		{"set value: -1", -1, -1},
		{"set value: 1", 1, 1},
		{"set value: 99999", 99999, 99999},
		{"set value: -99999", -99999, -99999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			element := NewEmptyNode[int64]()
			element.SetValue(tt.elementValue)

			if equal := reflect.DeepEqual(element.GetValue(), tt.result); !equal {
				t.Errorf("SetValue(%v) = %d, want %d", element, element.GetValue(), tt.result)
			}
		})
	}
}

func TestGetValue(t *testing.T) {
	element := NewNode[string]("some string", nil, nil)
	assert.Equal(t, "some string", element.GetValue())
}

func TestSetNext(t *testing.T) {
	element := NewEmptyNode[string]()
	element.SetNext(NewNode[string]("some string", nil, nil))
	assert.Equal(t, "some string", element.GetNext().GetValue())
}

func TestGetNext(t *testing.T) {
	element := NewNode[int64](0, NewNode[int64](1, nil, nil), nil)
	assert.Equal(t, int64(1), element.GetNext().GetValue())
}

func TestGetNextWithNil(t *testing.T) {
	element := NewEmptyNode[int64]()
	assert.Nil(t, element.GetNext())
}

func TestGetNextWithNilNext(t *testing.T) {
	element := NewNode[int64](0, nil, nil)
	assert.Nil(t, element.GetNext())
}

func TestNodeWithSomeStruct(t *testing.T) {
	type someStruct struct {
		x int
		y int
	}
	// Create test data
	data1 := someStruct{x: 42, y: 24}
	data2 := someStruct{x: 100, y: 200}

	// Create linked elements
	element2 := NewNode[someStruct](data2, nil, nil)
	element1 := NewNode[someStruct](data1, element2, nil)

	// Test first element
	assert.Equal(t, data1, element1.GetValue())

	// Test linked element
	nextElement := element1.GetNext()
	assert.NotNil(t, nextElement)
	assert.Equal(t, data2, nextElement.GetValue())
}
