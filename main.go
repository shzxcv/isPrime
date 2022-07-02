package main

import (
	"fmt"
	"math"
)

const count = 10000000

func main() {
	for n := 1; n <= count; n++ {
		isPrime := isPrimeGo(n)
		if isPrime {
			fmt.Printf("%d is prime\n", n)
		} else {
			fmt.Printf("%d is not prime\n", n)
		}
	}
}

func isPrimeRegex(num int) bool {
	return true
}

func isPrimeGo(num int) bool {
	if num == 1 {
		return false
	}
	if num == 2 {
		return true
	}
	if num%2 == 0 {
		return false
	}
	root := int(math.Sqrt(float64(num)))
	for p := 3; p <= root+1; p += 2 {
		if num%p == 0 {
			return false
		}
	}
	return true
}
