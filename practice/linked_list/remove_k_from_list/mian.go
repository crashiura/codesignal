package remove_k_from_list

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func removeKFromList(l *ListNode, k int) *ListNode {
	if l == nil {
		return l
	}

	v, _ := l.Value.(int)
	if v == k {
		l = l.Next
	}

	node := l
	for node != nil && node.Next != nil {
		if v, _ := node.Next.Value.(int); v == k {
			node.Next = node.Next.Next
		} else {
			node = node.Next
		}
	}

	if node != nil {
		if v, _ = node.Value.(int); v == k {
			l = l.Next
		}
	}

	return l
}
