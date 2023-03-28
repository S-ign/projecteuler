// Multiples of 3 or 5
// If we list all the natural numbers below 10 that are multiples
// of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
// Find the sum of all the multiples of 3 or 5 below 1000.

package main

import (
	"fmt"
)

func getMultiples(multipleOf int, bellow int) (multiples []int) {
	for i := 1; i < bellow; i++ {
		if i%multipleOf == 0 {
			multiples = append(multiples, i)
		}
	}
	return
}

// getListMultiples provide a list to get the multiples of
// all numbers bellow provided number.
// l = {3, 5}, bellow = 100
// Get a combined list without duplicates of the multiples
// of 3 or 5 of all number bellow 100
func getListMultiples(l []int, bellow int) (multiples []int) {
	m := [][]int{}
	// get multiples for each number in l, append to m
	for _, v := range l {
		m = append(m, getMultiples(v, bellow))
	}

	multiples = removeDuplicates(combineIntSlice(m))

	return
}

func sumSlice(s []int) (sum int) {
	for _, v := range s {
		sum += v
	}
	return
}

func combineIntSlice(a [][]int) (c []int) {
	for _, v := range a {
		c = append(c, v...)
	}
	return
}

func removeDuplicates(a []int) (noDuplicates []int) {
	unique := make(map[int]bool)
	for _, v := range a {
		unique[v] = true
	}

	for k := range unique {
		noDuplicates = append(noDuplicates, k)
	}

	return
}

func main() {
	multiplesOf := []int{3, 5}
	fmt.Println(sumSlice(getListMultiples(multiplesOf, 1000)))
}
