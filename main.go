package main

import "fmt"

const count = 100

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
	if num == 1 || num == 2 {
		return true
	}
	for p := 2; p <= num-1; p++ {
		if num%p == 0 {
			return false
		}
	}
	return true
}
