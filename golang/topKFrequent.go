package main

/*

https://leetcode.com/problems/top-k-frequent-elements/

Given an integer array nums and an integer k, return the k most frequent elements. You may return the answer in any order.

Example 1:

Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]

Example 2:

Input: nums = [1], k = 1
Output: [1]

Constraints:
    1 <= nums.length <= 105
    -104 <= nums[i] <= 104
    k is in the range [1, the number of unique elements in the array].
    It is guaranteed that the answer is unique.

Follow up: Your algorithm's time complexity must be better than O(n log n), where n is the array's size.
*/

type Element struct {
	Value int
	Freq  int
}

// FIXME: I have no idea why it is failing on the third test case
func topKFrequent(nums []int, k int) []int {
	freq := make(map[int]int, len(nums))
	hist := make(map[int][]int, len(nums))
	var topKElements []*Element
	for i := 0; i < k; i++ {
		topKElements = append(topKElements, &Element{})
	}
	topK := make([]int, k)
	for _, num := range nums {
		freq[num]++
		f := freq[num]
		hist[num] = append(hist[num], f)
		for i, n := range topKElements {
			if f > n.Freq {
				n.Value = num
				n.Freq = f
				topK[i] = num
				break
			}
		}
	}
	return topK
}
