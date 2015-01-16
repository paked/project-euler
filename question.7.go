package main

import (
	"fmt"
)

func main() {
	count := 0
	x := 3
	for {
		if isPrime(x) {
			count += 1
			fmt.Println(x, "is a prime number")
		}

		if count == 10001 {
			fmt.Println(x)
			break
		}

		x += 1
	}

	// fmt.Println(isPrime(5), isPrime(7), isPrime(6))

}

func isPrime(n int) bool {
	for i := 3; i < n; i++ {
		// println(i)
		if n%i == 0 {
			// println(i)
			return false
		}

	}

	return true
}
