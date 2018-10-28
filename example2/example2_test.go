package example2

import (
	"testing"
)

/*
$ go test -test.bench=".*" ./example2
*/

func BenchmarkGetProfile1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetProfile1()
	}
}

func BenchmarkGetProfile2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = GetProfile2()
	}
}
