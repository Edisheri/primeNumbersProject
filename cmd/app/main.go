package main

import (
	"fmt"
	"primeNumbersProject/internal/sieve"
	"time"
)

func main() {
	var n, numThreads int
	fmt.Print("Введите верхнюю границу (N): ")
	fmt.Scanln(&n)
	fmt.Print("Введите количество потоков (M): ")
	fmt.Scanln(&numThreads)

	// Последовательный алгоритм
	fmt.Println("\n--- Последовательный алгоритм ---")
	start := time.Now()
	sequentialPrimes := sieve.SequentialSieve(n)
	seqDuration := time.Since(start)
	fmt.Printf("Найдено %d простых чисел. Время выполнения: %v\n", len(sequentialPrimes), seqDuration)

	// Параллельный алгоритм (декомпозиция по данным)
	fmt.Println("\n--- Параллельный алгоритм: Декомпозиция по данным ---")
	start = time.Now()
	parallelPrimesByData := sieve.ParallelSieveByData(n, numThreads)
	parDurationData := time.Since(start)
	fmt.Printf("Найдено %d простых чисел. Время выполнения: %v\n", len(parallelPrimesByData), parDurationData)

	// Расчет ускорения и эффективности для декомпозиции по данным
	speedupData := seqDuration.Seconds() / parDurationData.Seconds()
	efficiencyData := speedupData / float64(numThreads)
	fmt.Printf("Ускорение: %.2f, Эффективность: %.2f%%\n", speedupData, efficiencyData*100)

	// Параллельный алгоритм (пул потоков)
	fmt.Println("\n--- Параллельный алгоритм: Пул потоков ---")
	start = time.Now()
	parallelPrimesWithPool := sieve.ParallelSieveWithThreadPool(n, numThreads)
	parDurationPool := time.Since(start)
	fmt.Printf("Найдено %d простых чисел. Время выполнения: %v\n", len(parallelPrimesWithPool), parDurationPool)

	// Расчет ускорения и эффективности для пула потоков
	speedupPool := seqDuration.Seconds() / parDurationPool.Seconds()
	efficiencyPool := speedupPool / float64(numThreads)
	fmt.Printf("Ускорение: %.2f, Эффективность: %.2f%%\n", speedupPool, efficiencyPool*100)

	// Выводы
	fmt.Println("\n--- Выводы ---")
	if efficiencyData < 0.5 || efficiencyPool < 0.5 {
		fmt.Println("❌ Параллельный алгоритм неэффективен при текущих настройках. Возможные причины:")
		if n < 10000 {
			fmt.Println("   - Слишком маленькое значение N, затраты на управление потоками превышают выигрыш.")
		}
		if numThreads > 10 {
			fmt.Println("   - Слишком большое количество потоков M относительно объема работы.")
		}
	} else {
		fmt.Println("✅ Параллельный алгоритм показал хорошие результаты для текущих параметров.")
		if speedupData > 1.5 || speedupPool > 1.5 {
			fmt.Println("   - Ускорение значительное, параллельная обработка эффективна.")
		} else {
			fmt.Println("   - Ускорение присутствует, но незначительное. Рекомендуется увеличить значение N.")
		}
	}
	fmt.Println("\nСинхронизация потоков необходима для предотвращения гонок данных.")
	fmt.Println("Использование критических секций или потокобезопасных структур данных гарантирует корректность.")
}
