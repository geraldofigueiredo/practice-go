package main

import (
	"fmt"

	"github.com/geraldofigueiredo/practice-go/data_structures/arrays"
)

func main() {
	// element := linkedlist.NewEmptyListElement()
	// element.SetValue(100)
	// element.SetNext(linkedlist.NewEmptyListElement())

	// fmt.Println(element, element.GetNext())

	vector, err := arrays.NewVector(17)
	if err != nil {
		panic(err.Error())
	}
	fmt.Printf("vector: %+v\n", vector)
	fmt.Println(vector.Len())
}
