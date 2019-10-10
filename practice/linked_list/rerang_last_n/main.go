package main

import "fmt"

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func main() {
	data := []int{2, 1}
	k := 1
	res := rearrangeLastN(buildList(data), k)
	fmt.Println(listToArr(res))
}

func rearrangeLastN(head *ListNode, n int) *ListNode {
	newTail := head
	tail := head
	count := 0

	for current := head; current != nil; current = current.Next {
		tail = current
		if count != n+1 {
			count++
			continue
		}

		newTail = newTail.Next
	}

	if n == 0 || count <= n {
		return head
	}

	tmp := newTail.Next
	newTail.Next = nil
	tail.Next = head

	return tmp
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

func listToArr(l *ListNode) []int {
	arr := make([]int, 0)

	for l != nil {
		v, _ := l.Value.(int)
		arr = append(arr, v)
		l = l.Next
	}

	return arr
}
