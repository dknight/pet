package pet

import "testing"

type testpaircat struct {
	Vet
	answer int
}

func TestingHumanYears(t *testing.T) {
	pairs := []testpaircat{
		{
			Vet: Cat{
				Name:  "Barsik",
				Years: 3,
			},
			answer: 30,
		},
		{
			Vet: Cat{
				Name:  "Murka",
				Years: 10,
			},
			answer: 100,
		},
		{
			Vet: Dog{
				Name:  "Bobik",
				Years: 5,
			},
			answer: 35,
		},
		{
			Vet: Dog{
				Name:  "Barboss",
				Years: 9,
			},
			answer: 63,
		},
	}

	for _, p := range pairs {
		if p.HumanYears() != p.answer {
			t.Error(
				"for", p.Vet,
				"expected", p.answer,
				"got", p.HumanYears(),
			)
		}
	}
}
