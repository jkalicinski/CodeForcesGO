package main

import "fmt"

func main() {
	input := 0
	fmt.Scanln(&input)

	if input <= 3 {
		fmt.Println("NO")
		return
	}

	if input%2 == 0 {
		fmt.Println("YES")
		return
	}

	fmt.Println("NO")
}
