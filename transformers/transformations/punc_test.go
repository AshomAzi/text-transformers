package transformations

import (
	"testing"
)

func TestPunc(t *testing.T) {
	got := Punc("Welcome back home brother , see you soon !")
	expected := "Welcome back home brother, see you soon!"
	if got != expected {
		t.Errorf("Expected %q, got %q", expected, got)
	}
}

func TestMPunc(t *testing.T) {
	got := MPunc(`He said: " I love you man ", don't hurt a brother ' ok '`)
	expected := `He said: "I love you man", don't hurt a brother 'ok'`
	if got != expected {
		t.Errorf("Expected %q, got %q", expected, got)
	}
}
