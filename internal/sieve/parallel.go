package sieve

import (
	"sync"
)

// Декомпозиция по данным
func ParallelSieveByData(n, numThreads int) []int {
	basePrimes := SequentialSieve(int(float64(n) * 0.5)) // Находим базовые простые
	isPrime := make([]bool, n+1)
	for i := range isPrime {
		isPrime[i] = true
	}

	var wg sync.WaitGroup
	chunkSize := (n + numThreads - 1) / numThreads

	for t := 0; t < numThreads; t++ {
		start := t * chunkSize
		end := (t + 1) * chunkSize
		if end > n {
			end = n
		}

		wg.Add(1)
		go func(start, end int) {
			defer wg.Done()
			for _, p := range basePrimes {
				for multiple := max(p*p, (start+(p-1))/p*p); multiple < end; multiple += p {
					isPrime[multiple] = false
				}
			}
		}(start, end)
	}

	wg.Wait()
	primes := []int{}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}

// Пул потоков
func ParallelSieveWithThreadPool(n, numThreads int) []int {
	basePrimes := SequentialSieve(int(float64(n) * 0.5))
	isPrime := make([]bool, n+1)
	for i := range isPrime {
		isPrime[i] = true
	}

	tasks := make(chan int, len(basePrimes))
	var wg sync.WaitGroup
	for t := 0; t < numThreads; t++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for p := range tasks {
				for multiple := p * p; multiple <= n; multiple += p {
					isPrime[multiple] = false
				}
			}
		}()
	}

	for _, p := range basePrimes {
		tasks <- p
	}
	close(tasks)
	wg.Wait()

	primes := []int{}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i)
		}
	}
	return primes
}
