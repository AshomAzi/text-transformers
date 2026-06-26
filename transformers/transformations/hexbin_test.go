package transformations

import "testing"

func TestHexBIn(t *testing.T) {
	got := HexBin("42 (hex) 1010 (bin)")
	expected := "66 10"
	if got != expected {
		t.Errorf("HexBin() = %q, want %q", got, expected)
	}
}

