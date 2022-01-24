package main

// Color

type Color int

const (
	red Color = iota
	green
	blue
	brown
)

// Size

type Size int

const (
	small Size = iota
	medium
	large
)

// Product

type Product struct {
	name  string
	color Color
	size  Size
}

// bad Filter

type Filter struct{}

func (f *Filter) FilterByColor(products []Product, color Color) []*Product {
	filtered := make([]*Product, 0)

	for i, product := range products {
		if product.color == color {
			filtered = append(filtered, &products[i])
		}
	}

	return filtered
}

func (f *Filter) FilterBySize(products []Product, size Size) []*Product {
	filtered := make([]*Product, 0)

	for i, product := range products {
		if product.size == size {
			filtered = append(filtered, &products[i])
		}
	}

	return filtered
}

func (f *Filter) FilterBySizeAndColor(products []Product, size Size, color Color) []*Product {
	filtered := make([]*Product, 0)

	for i, product := range products {
		if product.size == size && product.color == color {
			filtered = append(filtered, &products[i])
		}
	}

	return filtered
}

// Specification

type Specification interface {
	IsSatisfied(p *Product) bool
}

type ColorSpecification struct {
	color Color
}

func (c *ColorSpecification) IsSatisfied(p *Product) bool {
	return c.color == p.color
}

type SizeSpecification struct {
	size Size
}

func (s *SizeSpecification) IsSatisfied(p *Product) bool {
	return s.size == p.size
}

// composition

type AndSpecification struct {
	first, second Specification
}

func (s *AndSpecification) IsSatisfied(p *Product) bool {
	return s.first.IsSatisfied(p) && s.second.IsSatisfied(p)
}

// BetterFilter

type BetterFilter struct{}

func (b *BetterFilter) Filter(products []Product, spec Specification) []*Product {
	filtered := make([]*Product, 0)

	for i, product := range products {
		if spec.IsSatisfied(&product) {
			filtered = append(filtered, &products[i])
		}
	}

	return filtered
}
