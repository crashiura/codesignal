package main

import "fmt"

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func main() {
	data := []int{1, 2, 3, 4, 5}
	//data := []int{1, 2, 3}
	k := 2

	r := reverseNodesInKGroups(buildList(data), k)

	fmt.Println(listToArr(r))
}

func reverseNodesInKGroups(head *ListNode, k int) *ListNode {
	if k == 1 || head.Next == nil || head == nil {
		return head
	}

	tempHead := &ListNode{Value: 0, Next: head}
	current := tempHead

	for current.Next != nil {
		start := current.Next
		end := current.Next

		for i := 1; i < k; i++ {
			end = end.Next
		}

		if end == nil {
			break
		}

		next := end.Next
		end.Next = nil

		//reverse k'th list
		t := start
		var prev *ListNode
		for t != nil {
			next := t.Next
			t.Next = prev
			prev = t
			t = next
		}

		current.Next = end
		start.Next = next

		current = start
	}

	return tempHead.Next
}

func listToArr(l *ListNode) []int {
	arr := make([]int, 0)

	for l != nil {
		v, _ := l.Value.(int)
		arr = append(arr, v)
		l = l.Next
	}

	return arr
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
