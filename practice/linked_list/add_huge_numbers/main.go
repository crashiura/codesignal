package main

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func main() {
}

func addTwoHugeNumbers(a *ListNode, b *ListNode) *ListNode {
	aList := valToArr(a)
	bList := valToArr(b)
	if len(aList) == 0 {
		return b
	}
	if len(bList) == 0 {
		return a
	}

	i := len(aList) - 1
	j := len(bList) - 1

	var next *ListNode
	var node *ListNode
	var remainder int

	for 0 <= i || 0 <= j {
		var sum int
		if 0 <= i && 0 <= j {
			sum = aList[i] + bList[j]
		} else if 0 > i && 0 <= j {
			sum = bList[j]
		} else if 0 > j && 0 <= i {
			sum = aList[i]
		}

		if remainder != 0 {
			sum++
			remainder = 0
		}
		if sum > 9999 {
			sum -= 10000
			remainder++
		}

		node = &ListNode{
			Value: sum,
			Next:  nil,
		}

		if next != nil {
			node.Next = next
		}

		next = node
		i--
		j--
	}

	if remainder != 0 {
		return &ListNode{
			Value: remainder,
			Next:  node,
		}
	}

	return node
}

func valToArr(a *ListNode) []int {
	arr := make([]int, 0, 4)
	for a != nil {
		v, _ := a.Value.(int)
		arr = append(arr, v)
		a = a.Next
	}

	return arr
}
