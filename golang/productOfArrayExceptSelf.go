package main

import "math"

/*
https://leetcode.com/problems/product-of-array-except-self/

Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.



Example 1:

Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Example 2:

Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]



Constraints:

    2 <= nums.length <= 105
    -30 <= nums[i] <= 30
    The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.



Follow up: Can you solve the problem in O(1) extra space complexity? (The output array does not count as extra space for space complexity analysis.)
*/

func productExceptSelf(nums []int) []int {
	firstZeroIndex := -1
	totalProd := 1
	for i, n := range nums {
		if n != 0 {
			totalProd *= n
			continue
		}
		if firstZeroIndex != -1 {
			return make([]int, len(nums))
		}
		firstZeroIndex = i
	}
	products := make([]int, len(nums))
	if firstZeroIndex != -1 {
		products[firstZeroIndex] = totalProd
		return products
	}
	for i, n := range nums {
		products[i] = int(float64(totalProd) * math.Pow(float64(n), -1))
	}
	return products
}
