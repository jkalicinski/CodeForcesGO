package main

import "fmt"

func main() {
	dest := 1
	fmt.Scanln(&dest)
	steps := 0
	left := dest

	for {
		if left <= 5 {
			steps = steps + 1
			break
		} else {
			if left >= 5 {
				left = left - 5
				steps = steps + 1
			} else if left >= 4 {
				left = left - 4
				steps = steps + 1
			} else if left >= 3 {
				left = left - 3
				steps = steps + 1
			} else if left >= 2 {
				left = left - 2
				steps = steps + 1
			} else if left >= 1 {
				left = left - 1
				steps = steps + 1
			}
		}
	}

	fmt.Print(steps)
}
