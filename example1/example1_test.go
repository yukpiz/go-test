package example1

import (
	"testing"
)

func TestSum1(t *testing.T) {
	// Pass
}

func TestSum2(t *testing.T) {
	if Sum(1, 2) != 3 {
		t.Fail()
	}
}
