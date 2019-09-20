package main

import "testing"

func TestIsCrypt(t *testing.T) {
	tcs := []struct {
		crypt    []string
		solution [][]string
		result   bool
	}{
		{
			crypt: []string{"SEND", "MORE", "MONEY"},
			solution: [][]string{
				{"O", "0"},
				{"M", "1"},
				{"Y", "2"},
				{"E", "5"},
				{"N", "6"},
				{"D", "7"},
				{"R", "8"},
				{"S", "9"},
			},
			result: true,
		},

		{
			crypt: []string{"TEN", "TWO", "ONE"},
			solution: [][]string{
				{"O", "1"},
				{"T", "0"},
				{"W", "9"},
				{"E", "5"},
				{"N", "4"},
			},
			result: false,
		},
		{
			crypt: []string{"A", "A", "A"},
			solution: [][]string{
				{"A", "0"},
			},
			result: true,
		},
	}

	for _, tc := range tcs {
		if result := isCryptSolution(tc.crypt, tc.solution); result != tc.result {
			t.Errorf("Error expected %t, got %t, test data %v", result, tc.result, tc)
		}
	}

}
