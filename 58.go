package main

import (
	"fmt"
	"strings"
)

func lengthOfLastWord(s string) int {
	words := strings.Fields(s)

	if len(words) == 0 {
		return 0
	}

	return len(words[len(words)-1])
}

func main() {
	fmt.Println(lengthOfLastWord("ppp ss lllk"))
}
