package transformations

import "testing"

func TestVowels(t *testing.T) {
	got := Vowels("A egg, a apple")
	expected := "An egg, an apple"

	if got != expected {
		t.Errorf("got %q, expected %v", got, expected)
	}
}