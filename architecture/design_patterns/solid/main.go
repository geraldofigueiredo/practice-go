package main

import "fmt"

func main() {
	products := []Product{
		{"apple", red, small},
		{"alexa", blue, small},
		{"house", brown, large},
		{"earth", blue, large},
		{"earth2", blue, large},
	}

	filter := BetterFilter{}

	filtered := filter.Filter(products, &AndSpecification{&SizeSpecification{large}, &ColorSpecification{blue}})

	for _, v := range filtered {
		fmt.Printf("%+v\n", *v)
	}
}
