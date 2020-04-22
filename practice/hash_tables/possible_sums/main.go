package main

import (
	"fmt"
	"io"
	"os"
)

func main() {

	f, _ := os.Create(".test.txt")
	f.Write([]byte("asdf"))
	f.Truncate(60)

	buf := make([]byte, 2)
	if _, err := io.ReadFull(f, buf); err != nil {
		panic(err)
	}

	fmt.Println(buf)
	io.MultiReader()

	defer f.Close()
}

func possibleSums(coins []int, quantity []int) int {
	m := make(map[int]struct{})

	for i := 0; i < len(coins); i++ {
		tmpMap := make(map[int]struct{})
		for k := range m {
			tmpMap[k] = struct{}{}
		}
		for j := 1; j <= quantity[i]; {
			for v := range tmpMap {
				m[coins[i]*j+v] = struct{}{}
			}
			j++
			m[coins[i]*j-coins[i]] = struct{}{}
		}
	}
	return len(m)
}
