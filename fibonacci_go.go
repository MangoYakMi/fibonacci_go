package main

import "fmt"

// Функция для вычисления числа Фибоначчи
func calculateFibonacci(n int) int {
	if n <= 1 {
		return n
	}

	previous, current := 0, 1

	for i := 2; i <= n; i++ {
		next := previous + current
		previous = current
		current = next
	}

	return current
}

func main() {
	number := 10 // Вычисляем число Фибоначчи для number

	fibonacciNumber := calculateFibonacci(number)
	fmt.Printf("Число Фибоначчи для %d: %d\n", number, fibonacciNumber)
}
