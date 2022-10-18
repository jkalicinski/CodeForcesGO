package main

import "fmt"

func main() {
	var wordsCnt int
	fmt.Scanln(&wordsCnt)

	words := make([]string, int(wordsCnt))

	for i := 0; i < wordsCnt; i++ {
		fmt.Scanln(&words[i])
	}

	for i := 0; i < wordsCnt; i++ {
		if len(words[i]) <= 10 {
			fmt.Println(words[i])
			continue
		}

		var line = fmt.Sprint(words[i][0:1], len(words[i])-2, words[i][len(words[i])-1:])
		fmt.Println(line)
	}
}
