package main

import (
	"fmt"
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
			k:  3,
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
			in: []int{1},
			k:  1,
			result: &ListNode{
				Value: 1,
				Next:  nil,
			},
		},
		{
			in: []int{2, 1},
			k:  1,
			result: &ListNode{
				Value: 1,
				Next: &ListNode{
					Value: 2,
					Next:  nil,
				},
			},
		},
		{
			in: []int{1, 2, 3, 4},
			k:  2,
			result: &ListNode{
				Value: 3,
				Next: &ListNode{
					Value: 4,
					Next: &ListNode{
						Value: 1,
						Next: &ListNode{
							Value: 2,
							Next:  nil,
						},
					},
				},
			},
		},
		{
			in: []int{1, 2, 3, 4, 5, 6},
			k:  3,
			result: &ListNode{
				Value: 4,
				Next: &ListNode{
					Value: 5,
					Next: &ListNode{
						Value: 6,
						Next: &ListNode{
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
				},
			},
		},
		{
			in: []int{1, 2, 3},
			k:  2,
			result: &ListNode{
				Value: 2,
				Next: &ListNode{
					Value: 3,
					Next: &ListNode{
						Value: 1,
						Next:  nil,
					},
				},
			},
		},
		{
			in: []int{1, 2, 3, 4, 5, 6, 7},
			k:  4,
			result: &ListNode{
				Value: 4,
				Next: &ListNode{
					Value: 5,
					Next: &ListNode{
						Value: 6,
						Next: &ListNode{
							Value: 7,
							Next: &ListNode{
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
					},
				},
			},
		},
		{
			in: []int{1, 2, 3},
			k:  1,
			result: &ListNode{
				Value: 3,
				Next: &ListNode{
					Value: 1,
					Next: &ListNode{
						Value: 2,
						Next:  nil,
					},
				},
			},
		},
		{
			in: []int{1000000000, -849483855, -1000000000, 376365554, -847904832},
			k:  4,
			result: &ListNode{
				Value: -849483855,
				Next: &ListNode{
					Value: -1000000000,
					Next: &ListNode{
						Value: 376365554,
						Next: &ListNode{
							Value: -847904832,
							Next: &ListNode{
								Value: 1000000000,
								Next:  nil,
							},
						},
					},
				},
			},
		},
		{
			in: []int{1, 2, 3, 4, 5, 6, 7, 8},
			k:  5,
			result: &ListNode{
				Value: 4,
				Next: &ListNode{
					Value: 5,
					Next: &ListNode{
						Value: 6,
						Next: &ListNode{
							Value: 7,
							Next: &ListNode{
								Value: 8,
								Next: &ListNode{
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
						},
					},
				},
			},
		},
	}

	for _, tc := range testCases {
		result := rearrangeLastN(buildList(tc.in), tc.k)
		if result == nil {
			t.Errorf("result is nil")
		}
		for result != nil {
			v, _ := result.Value.(int)
			v2, _ := tc.result.Value.(int)
			fmt.Println(result)
			if v != v2 {
				t.Errorf("Error expected %d, got %d", v, v2)
			}
			result = result.Next
			tc.result = tc.result.Next
		}
	}
}
