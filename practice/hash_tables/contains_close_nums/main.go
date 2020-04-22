package main

import "fmt"

func main() {
	in := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9}

	k := 3
	fmt.Println(containsCloseNums(in, k))
}

func containsCloseNums(nums []int, k int) bool {
	if len(nums) == 0 || k == 0 {
		return false
	}

	m := make(map[int]int)
	k++

	var startIndex int
	for _, v := range nums {
		val, found := m[v]
		if !found {
			m[v] = 0
		} else {
			m[v] = val + 1
		}

		value, found := m[v]
		if found && value > 0 {
			return true
		}

		if len(m) == k {
			delete(m, nums[startIndex])
			startIndex++
		}
	}

	return false
}
