package main

import (
	"fmt"
)

func main() {
	t := [][]string{
		{".", ".", ".", "1", "4", ".", ".", "2", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", ".", ".", ".", ".", ".", ".", "."},
		{".", ".", "1", ".", ".", ".", ".", ".", "."},
		{".", "6", "7", ".", ".", ".", ".", ".", "9"},
		{".", ".", ".", ".", ".", ".", "8", "1", "."},
		{".", "3", ".", ".", ".", ".", ".", ".", "6"},
		{".", ".", ".", ".", ".", "7", ".", ".", "."},
		{".", ".", ".", "5", ".", ".", ".", "7", "."},
	}

	r := sudoku2(t)
	fmt.Println(r)
}

func sudoku2(grid [][]string) bool {
	l := len(grid)
	m := make(map[string]map[string]struct{}, len(grid))

	for i := 0; i < l; i++ {
		tmp := make(map[string]struct{}, len(grid[i]))
		for j := 0; j < l; j++ {
			currentPointVal := grid[j][i]
			// check row
			for p := 0; p < len(grid[j]); p++ {
				val := grid[j][p]
				if val == "." {
					continue
				}
				if val == currentPointVal && p != i {
					return false
				}
			}

			if currentPointVal == "." {
				continue
			}

			// check column
			_, found := tmp[currentPointVal]
			if found {
				return false
			}
			tmp[currentPointVal] = struct{}{}

			// check row 3x3
			currentMapKey := fmt.Sprintf("%d-%d", i/3, j/3)
			currentMap := m[currentMapKey]
			if currentMap == nil {
				currentMap = map[string]struct{}{}
				m[currentMapKey] = currentMap
			}

			if _, found := currentMap[currentPointVal]; found == true {
				return false
			}
			currentMap[currentPointVal] = struct{}{}
		}
	}

	return true
}
