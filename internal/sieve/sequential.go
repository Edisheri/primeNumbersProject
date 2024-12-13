package sieve

func SequentialSieve(n int) []int {
	isPrime := make([]bool, n+1) // Булевый массив для пометки чисел
	for i := 2; i <= n; i++ {
		isPrime[i] = true // Все числа изначально помечены как простые
	}

	// Убираем составные числа
	for p := 2; p*p <= n; p++ {
		if isPrime[p] {
			for multiple := p * p; multiple <= n; multiple += p {
				isPrime[multiple] = false // Помечаем число как составное
			}
		}
	}

	// Сохраняем простые числа в срез
	primes := []int{}
	for i := 2; i <= n; i++ {
		if isPrime[i] {
			primes = append(primes, i) // Добавляем число в список простых
		}
	}
	return primes
}
