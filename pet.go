// Package pet provides interface for working with pets.
package pet

// Pet defines pet type.
type Pet struct {
	Name  string
	Years int
}

// Vet defines interface for pets.
type Vet interface {
	HumanYears() int
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
