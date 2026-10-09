package twopointer

import "sort"

// Problem: 15. 3Sum
// Difficulty: Medium
// Link: https://leetcode.com/problems/3sum/
//
// Time Complexity: O(n^2) - O(n log n) for sorting and O(n^2) for the two-pointer traversal.
// Space Complexity: O(1) or O(log n) depending on the sorting implementation (excluding result storage).

// threeSum returns all unique triplets [nums[i], nums[j], nums[k]] that sum up to 0.
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	var result [][]int
	n := len(nums)

	for i := 0; i < n-2; i++ {
		// Nếu phần tử nhỏ nhất trong bộ 3 > 0, tổng 3 số chắc chắn > 0 (vì mảng đã sắp xếp tăng dần)
		if nums[i] > 0 {
			break
		}

		// Bỏ qua các giá trị trùng lặp ở vị trí phần tử thứ nhất
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			switch {
			case sum == 0:
				result = append(result, []int{nums[i], nums[left], nums[right]})

				// Bỏ qua các phần tử trùng lặp ở vị trí left và right
				for left < right && nums[left] == nums[left+1] {
					left++
				}
				for left < right && nums[right] == nums[right-1] {
					right--
				}

				left++
				right--
			case sum < 0:
				left++
			default: // sum > 0
				right--
			}
		}
	}

	return result
}
