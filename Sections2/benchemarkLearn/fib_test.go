package main

import "testing"

func BenchmarkFib(b *testing.B) {
	for i := 0; i < b.N; i++ {
		b.StopTimer()
		fib(30)
		b.StartTimer()
	}
}
