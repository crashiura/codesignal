package main

import "testing"

func TestSwapLexOrder(t *testing.T) {
	tCases := []struct {
		str    string
		pairs  [][]int
		result string
	}{
		{
			str:    "abdc",
			pairs:  [][]int{{1, 4}, {3, 4}},
			result: "dbca",
		},
		{
			str:    "abcdefgh",
			pairs:  [][]int{{1, 4}, {7, 8}},
			result: "dbcaefhg",
		},
		{
			str:    "acxrabdz",
			pairs:  [][]int{{1, 3}, {6, 8}, {3, 8}, {2, 7}},
			result: "zdxrabca",
		},
		{
			str:    "z",
			pairs:  [][]int{},
			result: "z",
		},
		{
			str:    "dznsxamwoj",
			pairs:  [][]int{{1, 3}, {3, 4}, {6, 5}, {8, 10}},
			result: "zdsnxamwoj",
		},
		{
			str:    "fixmfbhyutghwbyezkveyameoamqoi",
			pairs:  [][]int{{8, 5}, {10, 8}, {4, 18}, {20, 12}, {5, 2}, {17, 2}, {13, 25}, {29, 12}, {22, 2}, {17, 11}},
			result: "fzxmybhtuigowbyefkvhyameoamqei",
		},
	}

	for _, c := range tCases {
		result := swapLexOrder(c.str, c.pairs)
		if result != t.result {
			t.Errorf("expected %s, got %s", result, c.result)
		}
	}
}
