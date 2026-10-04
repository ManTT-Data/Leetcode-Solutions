package array

// Problem: 1. Two Sum
// Difficulty: Easy
// Link: https://leetcode.com/problems/two-sum/
//
// Time Complexity: O(n) - Single pass through the array.
// Space Complexity: O(n) - Hash map storing up to n elements.

// twoSum returns indices of the two numbers such that they add up to target.
func twoSum(nums []int, target int) []int {
	numToIndex := make(map[int]int, len(nums))

	for i, num := range nums {
		diff := target - num
		if idx, found := numToIndex[diff]; found {
			return []int{idx, i}
		}
		numToIndex[num] = i
	}

	return nil
}
