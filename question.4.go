package main

import (
	// "fmt"
	"strconv"
)

func main() {
	p := 0
	for x := 100; x < 999; x++ {
		for y := 100; y < 999; y++ {
			z := x * y
			str := strconv.Itoa(z)

			if str[:len(str)/2] == Reverse(str[len(str)/2:]) {
				if z > p {
					p = z
				}
			}

		}
	}

	println(p)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
