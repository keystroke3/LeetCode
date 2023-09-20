package main

func containsDuplicate(nums []int) (map[int]bool, bool) {
	seen := make(map[int]bool)
	for _, n := range nums {
		if seen[n] {
			return seen, true
		}
		seen[n] = true
	}
	return seen, false
}
