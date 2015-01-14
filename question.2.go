package main

func main() {
	a := 1
	b := 1
	count := 0

	for {
		num := a + b
		a = b
		b = num

		if num > 4000000 {
			break
		}

		if (num % 2) == 0 {
			count += num
		}

		// println(a)
	}
	println(count)
	// println(count)
}
