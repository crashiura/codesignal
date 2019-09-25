package main

import (
	"testing"
)

func TestMergeTwoLinkedList(t *testing.T) {
	tcs := []struct {
		a      []int
		b      []int
		result *ListNode
	}{
		{
			a: []int{1, 2, 3},
			b: []int{4, 5, 6},
			result: &ListNode{
				Value: 1,
				Next: &ListNode{
					Value: 2,
					Next: &ListNode{
						Value: 3,
						Next: &ListNode{
							Value: 4,
							Next: &ListNode{
								Value: 5,
								Next: &ListNode{
									Value: 6,
									Next:  nil,
								},
							},
						},
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		aL := buildList(tc.a)
		bL := buildList(tc.b)

		result := mergeTwoLinkedLists(aL, bL)
		for result != nil {
			if result.Value != tc.result.Value {
				t.Errorf("Error expected %d, got %d", result.Value, tc.result.Value)
			}
			result = result.Next
			tc.result = tc.result.Next
		}
	}
}

func buildList(a []int) *ListNode {
	if len(a) == 0 {
		return &ListNode{
			Value: 0,
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
