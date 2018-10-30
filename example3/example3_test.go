package example3_test

import (
	"testing"

	"../example3"
)

func TestSum1(t *testing.T) {
	if example3.Sum(1, 2) != 3 {
		t.Error("間違い！")
	}
}
