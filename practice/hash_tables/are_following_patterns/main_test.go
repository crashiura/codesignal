package main

import "testing"

func TestAreFollowingPatterns(t *testing.T) {
	tCases := []struct {
		str     []string
		pattern []string
		result  bool
	}{
		{
			str:     []string{"cat", "dog", "dog"},
			pattern: []string{"a", "b", "b"},
			result:  true,
		},
		{
			str:     []string{"cat", "dog", "doggy"},
			pattern: []string{"a", "b", "b"},
			result:  false,
		},
		{
			str:     []string{"cat", "dog", "dog"},
			pattern: []string{"a", "b", "c"},
			result:  false,
		},
		{
			str:     []string{"aaa"},
			pattern: []string{"aaa"},
			result:  true,
		},
		{
			str:     []string{"aaa", "aaa", "aaa"},
			pattern: []string{"aaa", "bbb", "aaa"},
			result:  false,
		},
	}

	for _, tc := range tCases {
		result := areFollowingPatterns(tc.str, tc.pattern)
		if result != tc.result {
			t.Errorf("Error expected %t, got %t, strings: %+v, pattern: %+v", result, tc.result, tc.str, tc.pattern)
		}
	}
}
