package main

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func main() {
}

func mergeTwoLinkedLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var head, node *ListNode

	v1, _ := l1.Value.(int)
	v2, _ := l2.Value.(int)

	if v1 > v2 {
		head = l2
		node = l2
		l2 = l2.Next
	} else {
		head = l1
		node = l1
		l1 = l1.Next
	}

	for l1 != nil && l2 != nil {

		v1, _ := l1.Value.(int)
		v2, _ := l2.Value.(int)

		if v1 > v2 {
			node.Next = l2
			node = node.Next
			l2 = l2.Next
		} else {
			node.Next = l1
			node = node.Next
			l1 = l1.Next
		}
	}

	if l1 != nil {
		node.Next = l1
	}

	if l2 != nil {
		node.Next = l2
	}

	return head
}
