// 2520 is the smallest number that can be divided by each
// of the numbers from 1 to 10 without any remainder.
// What is the smallest positive number that is evenly
// divisible by all of the numbers from 1 to 20?

package main

import "fmt"

func isDivisibleByAll(divisors []int, number int) bool {
	counter := 0
	for _, i := range divisors {
		if number%i == 0 {
			counter++
		} else {
			return false
		}
	}
	return len(divisors) == counter
}

func findGoodNumber() (theGoodNumber int) {
	untilTheGoodNumberIsFound := true

	divisors := make([]int, 20)
	for i := range divisors {
		divisors[i] = i + 1
	}

	badNumber := 2520

	for untilTheGoodNumberIsFound {
		if isDivisibleByAll(divisors, badNumber) {
			theGoodNumber = badNumber
			return
		}
		badNumber++
	}
	return
}

func main() {
	fmt.Println(findGoodNumber())
}
