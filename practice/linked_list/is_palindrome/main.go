package main

import "fmt"

type ListNode struct {
	Value interface{}
	Next  *ListNode
}

func main() {
	fmt.Println()
}
func isPalindrome(l *ListNode) bool {
	if l == nil {
		return true
	}

	if l.Next == nil {
		return true
	}

	numbers := make([]int, 0, 10)
	for l != nil {
		v, _ := l.Value.(int)
		numbers = append(numbers, v)
		l = l.Next
	}

	length := len(numbers) - 1
	for i, j := 0, length; i < len(numbers); i++ {
		if numbers[i] != numbers[j] {
			return false
		}
		j--
	}
	return true
}
