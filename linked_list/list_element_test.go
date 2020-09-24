package linkedlist

import (
	"reflect"
	"testing"
)

func TestNewEmptyListElement(t *testing.T) {
	got := NewEmptyListElement()
	want := &ListElement{}

	if equal := reflect.DeepEqual(got, want); !equal {
		t.Errorf("got %v want %v", got, want)
	}
}

func TestNewListElement(t *testing.T) {
	tests := []struct {
		name         string
		elementValue int64
		result       *ListElement
	}{
		{"element value: 0", 0, &ListElement{value: 0}},
		{"element value: -1", -1, &ListElement{value: -1}},
		{"element value: 1", 1, &ListElement{value: 1}},
		{"element value: 999999", 999999, &ListElement{value: 999999}},
		{"element value: -999999", -999999, &ListElement{value: -999999}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			element := NewListElement(tt.elementValue, nil)
			if equal := reflect.DeepEqual(element, tt.result); !equal {
				t.Errorf("NewListElement(%d) = %v, want %v", tt.elementValue, element, tt.result)
			}
		})
	}
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

func TestGetValue(t *testing.T) {
	tests := []struct {
		name         string
		elementValue int64
		result       int64
	}{
		{"get value: 0", 0, 0},
		{"get value: -1", -1, -1},
		{"get value: 1", 1, 1},
		{"get value: 99999", 99999, 99999},
		{"get value: -99999", -99999, -99999},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			element := NewListElement(tt.elementValue, nil)
			value := element.GetValue()
			if equal := reflect.DeepEqual(value, tt.result); !equal {
				t.Errorf("%v.GetValue() = %d, want %d", element, value, tt.result)
			}
		})
	}
}

func TestSetNext(t *testing.T) {
	next := NewListElement(0, nil)

	tests := []struct {
		name        string
		nextElement *ListElement
		result      *ListElement
	}{
		{"next element nil", nil, nil},
		{"next element equals", next, next},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			element := NewEmptyListElement()
			element.SetNext(tt.nextElement)

			if equal := reflect.DeepEqual(element.next, tt.result); !equal {
				t.Errorf("element.SetNext(%v) = %v, want %v", tt.nextElement, element, tt.result)
			}
		})
	}
}

func TestGetNext(t *testing.T) {
	next := NewListElement(0, nil)

	tests := []struct {
		name        string
		nextElement *ListElement
		result      *ListElement
	}{
		{"next element nil", nil, nil},
		{"next element equals", next, next},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			element := NewListElement(0, tt.nextElement)

			if equal := reflect.DeepEqual(element.GetNext(), tt.result); !equal {
				t.Errorf("element.GetNext() = %v, want %v", element.next, tt.result)
			}
		})
	}
}
