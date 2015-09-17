package main

import "testing"

func TestFibonnaci(t *testing.T) {
	result := fibonacci(2)
	if result != 1 {
		t.Errorf("The fibonacci is incorrect %d", result)
	}
}
