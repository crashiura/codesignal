package main

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	tcs := []struct {
		input  []int
		result bool
	}{
		{
			input:  []int{0, 1, 0},
			result: true,
		},
		{
			input:  []int{0, 1, 2, 3},
			result: false,
		},
		{
			input:  []int{0},
			result: true,
		},
		{
			input:  []int{0, 1},
			result: false,
		},
	}

	for _, tc := range tcs {
		l := buildList(tc.input)
		if res := isPalindrome(l); res != tc.result {
			t.Errorf("Error expected %t, got %t", res, tc.result)
			t.Failed()
		}
	}

}

func buildList(a []int) *ListNode {
	if len(a) == 0 {
		return &ListNode{
			Value: nil,
			Next:  nil,
		}
	}
	var prev *ListNode
	var root *ListNode

	for i := 0; i < len(a); i++ {
		node := &ListNode{
			Value: a[i],
			Next:  nil,
		}
		if prev != nil {
			prev.Next = node
		}

		if i == 0 {
			root = node
		}

		prev = node
	}

	return root
}
