package main

import "testing"

func benchmarkPartitions(k, n int, b *testing.B) {

	for n := 0; n < b.N; n++ {
		doubles(k, n)
	}

}

func BenchmarkPartitions10(b *testing.B)    { benchmarkPartitions(10, 10, b) }
func BenchmarkPartitions20(b *testing.B)    { benchmarkPartitions(20, 20, b) }
func BenchmarkPartitions50(b *testing.B)    { benchmarkPartitions(50, 50, b) }
func BenchmarkPartitions100(b *testing.B)   { benchmarkPartitions(100, 100, b) }
func BenchmarkPartitions200(b *testing.B)   { benchmarkPartitions(200, 200, b) }
func BenchmarkPartitions500(b *testing.B)   { benchmarkPartitions(500, 500, b) }
func BenchmarkPartitions1000(b *testing.B)  { benchmarkPartitions(1000, 1000, b) }
func BenchmarkPartitions2000(b *testing.B)  { benchmarkPartitions(2000, 2000, b) }
func BenchmarkPartitions5000(b *testing.B)  { benchmarkPartitions(5000, 5000, b) }
func BenchmarkPartitions10000(b *testing.B) { benchmarkPartitions(10000, 10000, b) }
