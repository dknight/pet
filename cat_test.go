package pet

import "testing"

type testpaircat struct {
	Cat
	answer int
}

func TestingDogYears(t *testing.T) {
	pairs := []testpaircat{
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

	for _, p := range pairs {
		if p.HumanYears() != p.answer {
			t.Error(
				"for", p.Cat,
				"expected", p.answer,
				"got", p.HumanYears(),
			)
		}
	}
}
