package pet

// Dog type defines structure of a dog.
type Dog struct {
	Name  string
	Years int
}

// HumanYears calculates human years to dog years.
func (d Dog) HumanYears() int {
	return d.Years * 7
}
