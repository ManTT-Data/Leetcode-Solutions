package twopointer

import "sort"

// Problem: 16. 3Sum Closest
// Difficulty: Medium
// Link: https://leetcode.com/problems/3sum-closest/
//
// Time Complexity: O(n^2) - O(n log n) for sorting and O(n^2) for the two-pointer search.
// Space Complexity: O(1) or O(log n) auxiliary space used by sorting.

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// threeSumClosest finds three integers whose sum is closest to target.
func threeSumClosest(nums []int, target int) int {
	sort.Ints(nums)
	n := len(nums)
	closestSum := nums[0] + nums[1] + nums[2]

	for i := 0; i < n-2; i++ {
		// Bỏ qua giá trị trùng lặp ở vị trí phần tử thứ nhất để tối ưu
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		left, right := i+1, n-1
		for left < right {
			sum := nums[i] + nums[left] + nums[right]

			// Nếu tìm thấy tổng bằng chính xác target, đây là khoảng cách gần nhất (bằng 0)
			if sum == target {
				return target
			}

			// Cập nhật kết quả nếu khoảng cách tới target nhỏ hơn
			if abs(sum-target) < abs(closestSum-target) {
				closestSum = sum
			}

			if sum < target {
				left++
			} else {
				right--
			}
		}
	}

	return closestSum
}
