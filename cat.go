package pet

// Cat type defines structure of a cat.
type Cat struct {
	Name  string
	Years int
}

// HumanYears calculates human years to cat years.
func (c Cat) HumanYears() int {
	return c.Years * 9
}
