package main

import "fmt"

func main() {
	word := ""
	transl := ""

	fmt.Scanln(&word)
	fmt.Scanln(&transl)

	if len(word) != len(transl) {
		fmt.Println("NO")
		return
	}

	for i := 0; i < len(word); i++ {
		if word[i] != transl[len(word)-1-i] {
			fmt.Println("NO")
			return
		}
	}

	fmt.Println("YES")
}
