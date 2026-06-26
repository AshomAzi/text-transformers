package transformations

import "testing"

func TestCases(t *testing.T) {
	got := Cases("hey!, (up) HOW (low) are you doing (cap) ?")
	expected := "HEY!, how are you Doing ?"
	if got != expected {
		t.Errorf("expected %q, got %q", got, expected)
	}
}

func TestCaseN(t *testing.T) {
	got := CaseN("hey john!, how is everyone (up, 4) AND WHAT (low, 2) are you up to (cap, 3) ?")
	expected := "hey JOHN!, HOW IS EVERYONE and what are You Up To ?"
	if got != expected {
		t.Errorf("expected %q, got %q", got, expected)
	}
}


