package main

import (
	"fmt"
)

func twoSum(nums []int, target int) (a int, b int) {
	diffs := make(map[int]int)
	for i, n := range nums {
		inverse := target - n
		if p, seen := diffs[n]; seen {
			return i, p
		}
		diffs[inverse] = i
	}
	fmt.Println(diffs)
	return
}
