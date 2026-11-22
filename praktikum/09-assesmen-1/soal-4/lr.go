package main

import (
	"fmt"
)

func isPrime(n int) bool {
	if n < 2 {
		return false
	}
	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}

func isNonDecreasing(n int) bool {
	if n < 10 {
		return true
	}
	lastDigit := 10
	for n > 0 {
		currentDigit := n % 10
		if currentDigit > lastDigit {
			return false
		}
		lastDigit = currentDigit
		n /= 10
	}
	return true
}

func main() {
	var L, R int
	if _, err := fmt.Scan(&L, &R); err != nil {
		return
	}

	count := 0

	for i := L; i <= R; i++ {
		sumDigits := 0
		temp := i
		for temp > 0 {
			sumDigits += temp % 10
			temp /= 10
		}

		if isNonDecreasing(i) && isPrime(sumDigits) {
			count++
		}
	}

	fmt.Println(count)
}
