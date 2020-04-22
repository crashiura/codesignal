package main

import "testing"

func TestPossibleSums(t *testing.T) {
	tCases := []struct {
		coins    []int
		quantity []int
		result   int
	}{
		{
			coins:    []int{10, 50, 100},
			quantity: []int{1, 2, 1},
			result:   9,
		},
		{
			coins:    []int{10, 50, 100, 500},
			quantity: []int{5, 3, 2, 2},
			result:   122,
		},
		{
			coins:    []int{1},
			quantity: []int{5},
			result:   5,
		},
	}

	for _, tc := range tCases {
		res := possibleSums(tc.coins, tc.quantity)
		if res != tc.result {
			t.Errorf("Error expected %d,got %d", tc.result, res)
		}
	}
}
