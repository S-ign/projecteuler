// By listing the first six prime numbers:
// 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.
// What is the 10 001st prime number?

package main

import "fmt"

func isPrime(n int) bool {
	if n == 1 {
		return false
	}
	factors := []int{}
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
		if len(factors) > 2 {
			return false
		}
	}
	return true
}

func findNPrime(n int) (nPrime int, primes []int) {
	lookingForPrime := true
	next := 1

	primes = []int{}

	for lookingForPrime {
		if isPrime(next) {
			primes = append(primes, next)
		}
		if len(primes) == n {
			break
		}
		next++
	}
	return primes[len(primes)-1], primes
}

func main() {
	fmt.Println(findNPrime(10001))
}
