package main

import (
	"testing"
)

func TestReverseNodeInKGroups(t *testing.T) {
	testCases := []struct {
		in     []int
		k      int
		result *ListNode
	}{
		{
			in: []int{1, 2, 3},
			k:  2,
			result: &ListNode{
				Value: 2,
				Next: &ListNode{
					Value: 1,
					Next: &ListNode{
						Value: 3,
						Next:  nil,
					},
				},
			},
		},
		{
			in: []int{1, 2, 3},
			k:  1,
			result: &ListNode{
				Value: 1,
				Next: &ListNode{
					Value: 2,
					Next: &ListNode{
						Value: 3,
						Next:  nil,
					},
				},
			},
		},
		{
			in: []int{1, 2, 3},
			k:  3,
			result: &ListNode{
				Value: 3,
				Next: &ListNode{
					Value: 2,
					Next: &ListNode{
						Value: 1,
						Next:  nil,
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		result := reverseNodesInKGroups(buildList(tc.in), tc.k)
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
