package main

import (
	"testing"
)

func TestPassed(t *testing.T) {
	if passed(80) != true {
		t.Error("Expected true")
	}
}
