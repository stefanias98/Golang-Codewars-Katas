package main

import "testing"

func BenchmarkMain(b *testing.B) {
	m := []int{}

	for i := 1; i < 30; i++ {
		m = append(m, i)
	}

	for n := 0; n < b.N; n++ {
		sumOfDivided(m)
	}
}
