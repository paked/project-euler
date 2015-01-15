package main

import (
	"fmt"
)

func main() {
	sumOfSquares := 0
	sum := 0

	for i := 0; i <= 100; i++ {
		sum += i
		sumOfSquares += i * i
	}
	squareOfSum := sum * sum

	fmt.Printf("difference is %v\n", squareOfSum-sumOfSquares)
}
