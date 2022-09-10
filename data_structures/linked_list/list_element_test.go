package linkedlist

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

type someStruct struct {
	x int
	y int
}

func TestNewEmptyListElement(t *testing.T) {
	got := NewEmptyListElement[int]()
	want := &ListElement[int]{}

	if equal := reflect.DeepEqual(got, want); !equal {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestNewListElement(t *testing.T) {
	strData := "some string"
	strListElement := NewListElement[string](strData, nil)
	assert.IsType(t, "", strListElement.data)
	assert.Equal(t, strData, strListElement.GetValue())

	intData := 99
	intListElement := NewListElement[int](intData, nil)
	assert.IsType(t, 0, intListElement.data)
	assert.Equal(t, intData, intListElement.GetValue())

	someStructData := someStruct{}
	customTypeListElement := NewListElement[someStruct](someStructData, nil)
	assert.IsType(t, someStructData, customTypeListElement.data)
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
			element := NewEmptyListElement()
			element.SetValue(tt.elementValue)

			if equal := reflect.DeepEqual(element.value, tt.result); !equal {
				t.Errorf("SetValue(%v) = %d, want %d", element, element.value, tt.result)
			}
		})
	}
}

//
//func TestGetValue(t *testing.T) {
//	tests := []struct {
//		name         string
//		elementValue int64
//		result       int64
//	}{
//		{"get value: 0", 0, 0},
//		{"get value: -1", -1, -1},
//		{"get value: 1", 1, 1},
//		{"get value: 99999", 99999, 99999},
//		{"get value: -99999", -99999, -99999},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			element := NewListElement(tt.elementValue, nil)
//			value := element.GetValue()
//			if equal := reflect.DeepEqual(value, tt.result); !equal {
//				t.Errorf("%v.GetValue() = %d, want %d", element, value, tt.result)
//			}
//		})
//	}
//}
//
//func TestSetNext(t *testing.T) {
//	next := NewListElement(0, nil)
//
//	tests := []struct {
//		name        string
//		nextElement *ListElement
//		result      *ListElement
//	}{
//		{"next element nil", nil, nil},
//		{"next element equals", next, next},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//
//			element := NewEmptyListElement()
//			element.SetNext(tt.nextElement)
//
//			if equal := reflect.DeepEqual(element.next, tt.result); !equal {
//				t.Errorf("element.SetNext(%v) = %v, want %v", tt.nextElement, element, tt.result)
//			}
//		})
//	}
//}
//
//func TestGetNext(t *testing.T) {
//	next := NewListElement(0, nil)
//
//	tests := []struct {
//		name        string
//		nextElement *ListElement
//		result      *ListElement
//	}{
//		{"next element nil", nil, nil},
//		{"next element equals", next, next},
//	}
//
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//
//			element := NewListElement(0, tt.nextElement)
//
//			if equal := reflect.DeepEqual(element.GetNext(), tt.result); !equal {
//				t.Errorf("element.GetNext() = %v, want %v", element.next, tt.result)
//			}
//		})
//	}
//}
