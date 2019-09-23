package main

import (
	"testing"
)

func TestAddHugeNumbers(t *testing.T) {
	tcs := []struct {
		a      []int
		b      []int
		result *ListNode
	}{
		{
			a: []int{9876, 5432, 1999},
			b: []int{1, 8001},
			result: &ListNode{
				Value: 9876,
				Next: &ListNode{
					Value: 5434,
					Next: &ListNode{
						Value: 0000,
						Next:  nil,
					},
				},
			},
		},
		{
			a: []int{123, 4, 5},
			b: []int{100, 100, 100},
			result: &ListNode{
				Value: 223,
				Next: &ListNode{
					Value: 104,
					Next: &ListNode{
						Value: 105,
						Next:  nil,
					},
				},
			},
		},
		{
			a: []int{1},
			b: []int{9999, 9999, 9999, 9999, 9999, 9999},
			result: &ListNode{
				Value: 1,
				Next: &ListNode{
					Value: 0,
					Next: &ListNode{
						Value: 0,
						Next: &ListNode{
							Value: 0,
							Next: &ListNode{
								Value: 0,
								Next: &ListNode{
									Value: 0,
									Next: &ListNode{
										Value: 0,
										Next:  nil,
									},
								},
							},
						},
					},
				},
			},
		},
		{
			a: []int{1},
			b: []int{9999},
			result: &ListNode{
				Value: 1,
				Next: &ListNode{
					Value: 0,
					Next:  nil,
				},
			},
		},
		{
			a: []int{},
			b: []int{9999, 99},
			result: &ListNode{
				Value: 9999,
				Next: &ListNode{
					Value: 99,
					Next:  nil,
				},
			},
		},
		{
			a: []int{9999, 77},
			b: []int{9999, 99},
			result: &ListNode{
				Value: 1,
				Next: &ListNode{
					Value: 9998,
					Next: &ListNode{
						Value: 176,
						Next:  nil,
					},
				},
			},
		},
	}

	for _, tc := range tcs {
		aL := buildList(tc.a)
		bL := buildList(tc.b)
		result := addTwoHugeNumbers(aL, bL)
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
