package golib

import "testing"

func TestFunc1(t *testing.T) {
	expected := "func1v1"
	result := func1()

	if expected != result {
		t.Errorf("Expected %s, but got %s", expected, result)
	}
}
