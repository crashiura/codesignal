package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(firstNotRepeatingCharacter("abacabad"))
}

func firstNotRepeatingCharacter(s string) string {
	for _, v := range s {
		if strings.Count(s, string(v)) == 1 {
			return string(v)
		}
	}

	return "_"
}
