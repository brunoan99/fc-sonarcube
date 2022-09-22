package main

import "testing"

func TestSum(t *testing.T) {
	got := Sum(10, 10)
	expected := 20
	if got != expected {
		t.Errorf("expected %v got %v", expected, got)
	}
}
