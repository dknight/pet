package pet

import "testing"

type testpairdog struct {
	Dog
	answer int
}

func TestingCatYears(t *testing.T) {
	pairs := []testpairdog{
		{
			Dog: Dog{
				Name:  "Bobik",
				Years: 4,
			},
			answer: 28,
		},
		{
			Dog: Dog{
				Name:  "Barboss",
				Years: 10,
			},
			answer: 70,
		},
	}

	for _, p := range pairs {
		if p.HumanYears() != p.answer {
			t.Error(
				"for", p.Dog,
				"expected", p.answer,
				"got", p.HumanYears(),
			)
		}
	}
}
