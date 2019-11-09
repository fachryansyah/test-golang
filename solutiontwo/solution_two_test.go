package solutiontwo

import "testing"

func TestSortWord(t *testing.T) {
	result := SortWord("omama")
	if result != "aaomm" {
		t.Errorf("SOrtWord function was error, want: %s got: %s", "aaomm", result)
	}
}
