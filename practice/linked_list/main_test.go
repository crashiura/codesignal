package main

import (
	"testing"
)

func TestRemoveKFromList(t *testing.T) {
	tcs := []struct {
		input  []int
		k      int
		result *ListNode
	}{
		{
			input: []int{3, 1, 2, 4},
			k:     3,
			result: &ListNode{
				Value: 1,
				Next: &ListNode{
					Value: 2,
					Next: &ListNode{
						Value: 4,
						Next:  nil,
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		l := buildList(tc.input)
		result := removeKFromList(l, tc.k)
		for result != nil {
			v, _ := result.Value.(int)
			v2, _ := tc.result.Value.(int)
			if v != v2 {
				t.Errorf("Error expected %d, got %d", v, v2)
			}
			result = result.Next
			tc.result = tc.result.Next
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
