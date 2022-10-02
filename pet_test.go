package pet

import (
	"fmt"
	"testing"
)

type testpaircat struct {
	Cat
	answer int
}
type testpairdog struct {
	Dog
	answer int
}

func TestHumanYears(t *testing.T) {
	cats := []testpaircat{
		{
			Cat: Cat{
				Name:  "Barsik",
				Years: 3,
			},
			answer: 30,
		},
		{
			Cat: Cat{
				Name:  "Murka",
				Years: 10,
			},
			answer: 100,
		},
	}

	dogs := []testpairdog{
		{
			Dog: Dog{
				Name:  "Bobik",
				Years: 3,
			},
			answer: 21,
		},
		{
			Dog: Dog{
				Name:  "Barboss",
				Years: 7,
			},
			answer: 49,
		},
	}

	for _, c := range cats {
		if c.HumanYears() != c.answer {
			t.Error(
				"for", c.Cat,
				"expected", c.answer,
				"got", c.HumanYears(),
			)
		}
	}
	for _, d := range dogs {
		if d.HumanYears() != d.answer {
			t.Error(
				"for", d.Dog,
				"expected", d.answer,
				"got", d.HumanYears(),
			)
		}
	}
}

func ExampleCat() {
	cat := Cat{
		Name:  "Barsik",
		Years: 2,
	}
	fmt.Println(cat.HumanYears())
	// Output: 20
}
func ExampleDog() {
	dog := Dog{
		Name:  "Barboss",
		Years: 3,
	}
	fmt.Println(dog.HumanYears())
	// Output: 21
}
