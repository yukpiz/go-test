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

func TestSum3(t *testing.T) {
	if Sum(1, 2) != 3 {
		t.Error("e", "r", "r")
	}
}

func TestSum4(t *testing.T) {
	if Sum(1, 2) != 3 {
		t.Errorf("%s", "error")
	}
}

func TestSum5(t *testing.T) {
	if Sum(1, 2) != 3 {
		t.FailNow()
	}
}
