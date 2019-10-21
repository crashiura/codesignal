package gropingDishes

import (
	"sort"
)

func groupingDishes(dishes [][]string) [][]string {
	m := make(map[string][]string)
	keys := make([]string, 0)

	for i := range dishes {
		dishName := dishes[i][0]
		for j := 1; j < len(dishes[i]); j++ {
			ingredient, ok := m[dishes[i][j]]
			if !ok {
				keys = append(keys, dishes[i][j])
				d := make([]string, 0)
				d = append(d, dishName)
				m[dishes[i][j]] = d
				continue
			}
			ingredient = append(ingredient, dishName)
			m[dishes[i][j]] = ingredient
		}
	}

	sort.Strings(keys)
	result := make([][]string, 0)
	for _, ingName := range keys {
		data := m[ingName]
		if len(data) <= 1 {
			continue
		}
		sort.Strings(data)
		res := make([]string, 0, len(data)+1)
		res = append(res, ingName)
		res = append(res, data...)
		result = append(result, res)
	}

	return result
}
