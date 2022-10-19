package main

import "fmt"

func main() {
	games := 0
	scores := ""

	fmt.Scanln(&games)
	fmt.Scanln(&scores)
	a := 0
	d := 0
	winMin := games/2 + 1

	for i := 0; i < games; i++ {
		if scores[i] == 65 {
			a++
		} else if scores[i] == 68 {
			d++
		}

		if a >= winMin {
			fmt.Println("Anton")
			return
		}

		if d >= winMin {
			fmt.Println("Danik")
			return
		}
	}

	if a > d {
		fmt.Println("Anton")
		return
	}

	if a > d {
		fmt.Println("Danik")
		return
	}

	fmt.Println("Friendship")
}
