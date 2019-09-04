package permutation_test

import (
	"fmt"
	"stringPermutations/permutation"
	"testing"
)

type permutationsTestCase struct {
	given string
	want  []string
}

func (p permutationsTestCase) String() string {
	return fmt.Sprintf("given %s, %v is wanted", p.given, p.want)
}

func TestCalculatePermutation(t *testing.T) {
	cases := []permutationsTestCase{
		{given: "a", want: []string{"a"}},
		{given: "b", want: []string{"b"}},
		{given: "ab", want: []string{"ab", "ba"}},
		{given: "abc", want: []string{"abc", "bac", "bca", "acb", "cab", "cba"}},
	}

	for _, testCase := range cases {
		got := permutation.Calculate(testCase.given)
		assertCalculatedArrayContainsAllItens(t, testCase.want, got)
	}
}

func assertCalculatedArrayContainsAllItens(t *testing.T, want, got []string) {
	t.Helper()
	for _, wantPer := range want {
		if notIn(wantPer, got) {
			t.Errorf("%s permutation is missing", wantPer)
		}
	}

}

func notIn(word string, words []string) bool {
	for _, w := range words {
		if w == word {
			return false
		}
	}

	return true
}
