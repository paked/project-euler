package main

import (
	"fmt"
)

func main() {
	x := 1

	for {
		found := true
		for y := 1; y <= 20; y++ {
			// fmt.Println(x, y)
			// fmt.Printf("%v is divisable by %v", x, y)
			if x%y != 0 {
				found = false
			}
		}

		if found {
			fmt.Println(x)
			break
		}
		x += 1
	}
}
