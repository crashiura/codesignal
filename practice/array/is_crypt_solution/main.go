package main

import "strconv"

func main() {
	crypt := []string{"SEND", "MORE", "MONEY"}
	solution := [][]string{
		{"O", "0"},
		{"M", "1"},
		{"Y", "2"},
		{"E", "5"},
		{"N", "6"},
		{"D", "7"},
		{"R", "8"},
		{"S", "9"},
	}

	isCryptSolution(crypt, solution)
}

func isCryptSolution(crypt []string, solution [][]string) bool {
	m := make(map[string]string, len(solution))
	for _, v := range solution {
		m[v[0]] = v[1]
	}

	strNums := [3]string{"", "", ""}
	for i, word := range crypt {
		for p, j := range word {
			strVal := m[string(j)]
			if p == 0 && strVal == "0" && len(word) > 1 {
				return false
			}
			strNums[i] = strNums[i] + strVal
		}
	}

	num1, _ := strconv.Atoi(strNums[0])
	num2, _ := strconv.Atoi(strNums[1])
	num3, _ := strconv.Atoi(strNums[2])

	return num1+num2 == num3
}
