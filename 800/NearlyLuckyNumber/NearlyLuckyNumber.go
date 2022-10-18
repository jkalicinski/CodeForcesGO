package main

import "fmt"

func main() {
	input := ""
	fmt.Scanln(&input)

	runes := []rune(input)
	luckyCount := 0

	for i := 0; i < len(runes); i++ {
		if runes[i] == 52 || runes[i] == 55 {
			luckyCount++
		}
	}

	if luckyCount == 7 || luckyCount == 4 {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
