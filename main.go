package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/scorpionknifes/go-pcre"
)

const (
	count = 100000000
	char  = "x"
)

var re = pcre.MustCompile(fmt.Sprintf(`^(?!(%s%s+)\1+$|^%s$)`, char, char, char), 0)

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
	chars := strings.Repeat(char, num)
	isPrime := re.Matcher([]byte(chars), 0).Matches()
	return isPrime
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
