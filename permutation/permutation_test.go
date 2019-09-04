package permutation_test

import (
	"fmt"
	"stringPermutations/permutation"
	"testing"
)

type permutationsTestCase struct {
	given string
	want []string
}

func (p permutationsTestCase) String() string {
	return fmt.Sprintf("given %s, %v is wanted", p.given, p.want)
}

func TestCalculatePermutation(t *testing.T) {
	cases := []permutationsTestCase {
		{given: "a", want: []string{"a"}},
		{given: "b", want: []string{"b"}},
		{given: "ab", want: []string{"ab", "ba"}},
		{given: "abc", want: []string{"abc", "bac", "bca", "acb", "cab", "cba"}},
	}

	t.Run("given a word get all its permutations", func(t *testing.T) {
		for _, c := range cases {
			got := permutation.Calculate(c.given)
			t.Run(c.String(), func(t *testing.T) {
				for _, wantPer := range c.want {
					if notIn(wantPer, got) {
						t.Errorf("%s permutation is missing", wantPer)
					}
				}
			})
		}
	})
}

func notIn(word string, words []string) bool {
	for _, w := range words {
		if w == word {
			return false
		}
	}

	return true
}