package solutiontwo

import "testing"

func TestSortWord(t *testing.T) {
	result := SortWord("osama")
	if result != "aaoms" {
		t.Errorf("SortWord function was error, want: %s got: %s", "aaoms", result)
	}
}
