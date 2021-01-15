package main

import "testing"

func BenchmarkCheck1(b *testing.B) {
	d := directionX

	for ii := 0; ii < b.N; ii++ {
		d.Check()
	}
}
