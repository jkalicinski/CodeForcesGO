package main

import "fmt"

func main() {
	input := 0
	fmt.Scanln(&input)

	for {

		input++

		if IsDistinct(input) {
			fmt.Println(input)
			return
		}
	}
}

func IsDistinct(year int) bool {
	var numbers = []rune(fmt.Sprint(year))

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			if numbers[i] == numbers[j] {
				return false
			}
		}
	}

	return true
}
