package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Алгоритм 1: Пошук через дільники
func Algorytm1(a, b int) int {
	var divisorsA, divisorsB []int
	for i := 1; i <= a; i++ {
		if a%i == 0 {
			divisorsA = append(divisorsA, i)
		}
	}
	for i := 1; i <= b; i++ {
		if b%i == 0 {
			divisorsB = append(divisorsB, i)
		}
	}
	// Знаходимо найбільший спільний дільник
	var result int
	for _, divisorA := range divisorsA {
		for _, divisorB := range divisorsB {
			if divisorA == divisorB && divisorA > result {
				result = divisorA
			}
		}
	}
	return result
}

// Алгоритм 2: Пошук через прості множники
func primeFactors(n int) []int {
	var factors []int
	for i := 2; i*i <= n; i++ {
		for n%i == 0 {
			factors = append(factors, i)
			n /= i
		}
	}
	if n > 1 {
		factors = append(factors, n)
	}
	return factors
}

func Algorytm2(a, b int) int {
	factorsA := primeFactors(a)
	factorsB := primeFactors(b)
	var result int = 1

	// Знаходимо спільні прості множники
	for _, factorA := range factorsA {
		for i, factorB := range factorsB {
			if factorA == factorB {
				result *= factorA
				factorsB = append(factorsB[:i], factorsB[i+1:]...) // Видаляємо цей множник
				break
			}
		}
	}
	return result
}

// Алгоритм 3: Алгоритм Евкліда
func Algorytm3(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	// Генерація випадкових чисел
	// Деякі зміни для гілки 2
	rand.Seed(time.Now().UnixNano())
	nums := make([][2]int, 1000)
	for i := 0; i < 1000; i++ {
		nums[i] = [2]int{rand.Intn(100000) + 1, rand.Intn(100000) + 1}
	}

	// Оцінка часу виконання кожного методу
	start := time.Now()
	for _, pair := range nums {
		Algorytm1(pair[0], pair[1])
	}
	fmt.Println("Час виконання Алгоритму 1:", time.Since(start))

	start = time.Now()
	for _, pair := range nums {
		Algorytm2(pair[0], pair[1])
	}
	fmt.Println("Час виконання Алгоритму 2:", time.Since(start))

	start = time.Now()
	for _, pair := range nums {
		Algorytm3(pair[0], pair[1])
	}
	fmt.Println("Час виконання Алгоритму 3:", time.Since(start))
}
