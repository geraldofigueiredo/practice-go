package main

import (
	"fmt"

	linkedlist "github.com/geraldofigueiredo/practice-go/data_structures/linked_list"
)

func main() {
	element := linkedlist.NewEmptyListElement()
	element.SetValue(100)
	element.SetNext(linkedlist.NewEmptyListElement())

	fmt.Println(element, element.GetNext())
}
