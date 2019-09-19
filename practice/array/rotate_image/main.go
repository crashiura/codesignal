package main

import "fmt"

func main() {
	expected := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	r := rotateImage(expected)
	fmt.Println(r)
}

func rotateImage(a [][]int) [][]int {
	result := make([][]int, len(a), len(a))
	for i := 0; i < len(a); i++ {
		for j := len(a) - 1; j >= 0; j-- {
			result[i] = append(result[i], a[j][i])
		}
	}
	return result
}
