// Package pet provides interface for working with pets.
package pet

// Vet defines interface for pets.
type Vet interface {
	HumanYears() int
}

// Pet defines pet type. At the moment 2 pets are implemented:
//	1. Dog
//	2. Cat
type Pet struct {
	Name  string
	Years int
	Vet
}

// Dog defines type of a dog.
type Dog Pet

// Cat defines type of a cat.
type Cat Pet

// HumanYears converts human years to dog years.
func (d Dog) HumanYears() int {
	return d.Years * 7
}

// HumanYears converts human years to cat years.
func (c Cat) HumanYears() int {
	return c.Years * 10
}
