package test

import (
	"primeNumbersProject/internal/sieve"
	"testing"
)

func BenchmarkSequentialSieve(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sieve.SequentialSieve(100000)
	}
}

func BenchmarkParallelSieveByData(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sieve.ParallelSieveByData(100000, 4)
	}
}

func BenchmarkParallelSieveWithThreadPool(b *testing.B) {
	for i := 0; i < b.N; i++ {
		sieve.ParallelSieveWithThreadPool(100000, 4)
	}
}
