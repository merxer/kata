package test

import (
	"testing"
)

func TestAdd(t *testing.T) {
	if Add(1,2) != 3 {
		t.Error("expected 3")
	}

	if Add(5,3) != 8 {
		t.Error("expect 8")
	}
}
