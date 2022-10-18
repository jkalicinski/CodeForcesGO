package main

import (
	"fmt"
	"math"
	"strings"
	"unicode"
)

func main() {
	input := ""
	fmt.Scanln(&input)
	runes := []rune(input)

	if len(runes) == 1 {
		fmt.Print(input)
		return
	}

	minUpper := 0

	if len(runes)%2 == 0 {
		minUpper = (len(input) / 2) + 1
	} else {
		minUpper = int(math.Ceil(float64(len(runes)) / 2.0))
	}

	upperCount := 0

	for i := 0; i < len(runes); i++ {
		if unicode.IsUpper(runes[i]) {
			upperCount++
		}

		if upperCount >= int(minUpper) {
			fmt.Print(strings.ToUpper(input))
			return
		}
	}

	fmt.Print(strings.ToLower(input))
}
