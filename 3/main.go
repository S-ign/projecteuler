// The prime factors of 13195 are 5, 7, 13 and 29.
// What is the largest prime factor of the number 600851475143 ?

package main

import "fmt"

func getFactors(n int) (factors []int) {
	for i := 1; i < n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return
}

func getPrime(n int) bool {
	if n == 1 {
		return false
	}
	factors := []int{}
	for i := 1; i < n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
		if len(factors) > 2 {
			return false
		}
	}
	return true
}

func getPrimes(factors []int) (primes []int) {
	for _, v := range factors {
		if getPrime(v) {
			primes = append(primes, v)
		}
	}
	return
}

func main() {
	fmt.Println(getPrimes(getFactors(600851475143)))
}
