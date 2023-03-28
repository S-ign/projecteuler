// The sum of the squares of the first ten natural numbers is,
// 1^2 + 2^2 + ... + 10^2 = 385

// The square of the sum of the first ten natural numbers is,
// (1 + 2 + ... + 10)^2 = 55^2 = 3025

// Hence the difference between the sum of the squares of the
// first ten natural numbers and the square of the sum is 3025
// 3025 - 385 = 2640

// Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.

package main

import (
	"fmt"
	"math"
)

func sumOfSquares(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += int(math.Pow(float64(i), 2))
	}
	return
}

func squareOfSum(n int) (sum int) {
	for i := 1; i <= n; i++ {
		sum += i
	}
	return int(math.Pow(float64(sum), 2))
}

func main() {
	fmt.Println(squareOfSum(100) - sumOfSquares(100))
}
