package util

import "testing"

func AssertEq(t *testing.T, a, b interface{}) {
	if a != b {
		t.Fatalf("Assertion failed: %v != %v", a, b)
	}
}
