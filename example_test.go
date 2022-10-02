package pet

import "fmt"

func ExampleCat_HumanYears() {
	cat := Cat{
		Name:  "Barsik",
		Years: 2,
	}
	fmt.Println(cat.HumanYears())
	// Output: 20
}
func ExampleDog_HumanYears() {
	dog := Dog{
		Name:  "Barboss",
		Years: 3,
	}
	fmt.Println(dog.HumanYears())
	// Output: 21
}
