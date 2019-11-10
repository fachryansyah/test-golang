package solutionone

import "testing"

func TestCountWord(test *testing.T) {
	vocal, consonant := CountWord("omama")
	if vocal != 2 && consonant != 2 {
		test.Errorf("Count function was error, want(consonant: %d, vocal: %d) got(consonant: %d, vocal: %d)", 2, 2, consonant, vocal)
	}
}
