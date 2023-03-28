// A palindromic number reads the same both ways. The largest
// palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.

package main

import "fmt"

func isPalindromic(num int) bool {
	res := 0
	input := num
	for num > 0 {
		remainder := num % 10
		res = (res * 10) + remainder
		num /= 10
	}
	if input == res {
		return true
	}
	return false
}

func findPalindromes(n int) (palindromes []int) {

	for a := n; a < n+100; a++ {
		for b := n; b < n+100; b++ {
			if isPalindromic(a * b) {
				fmt.Println(a * b)
				palindromes = append(palindromes, a*b)
			}
		}
	}

	return
}

func main() {
	fmt.Println(findPalindromes(900))
}
