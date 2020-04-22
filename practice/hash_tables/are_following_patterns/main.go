package main

import "fmt"

func main() {
	str := []string{"cat", "dog", "dog"}
	p := []string{"a", "b", "b"}
	r := areFollowingPatterns(str, p)

	fmt.Println(r)
}

func areFollowingPatterns(strings []string, patterns []string) bool {
	if len(strings) != len(patterns) {
		return false
	}

	str := make(map[string]string)
	patt := make(map[string]string)

	for i := 0; i < len(strings); i++ {
		str[strings[i]] = patterns[i]
	}

	for j := 0; j < len(patterns); j++ {
		if r, ok := patt[patterns[j]]; ok && r != strings[j] {
			return false
		}

		patt[patterns[j]] = strings[j]
	}

	if len(str) == len(patt) {
		return true
	} else {
		return false
	}
}
