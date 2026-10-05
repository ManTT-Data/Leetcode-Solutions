package array

// Problem: 33. Search in Rotated Sorted Array
// Difficulty: Medium
// Link: https://leetcode.com/problems/search-in-rotated-sorted-array/
//
// Time Complexity: O(log n) - Modified binary search dividing the search space by half each iteration.
// Space Complexity: O(1) - Constant auxiliary space.

// search finds the index of target in a rotated sorted array, or returns -1 if not found.
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return mid
		}

		// Nhánh 1: Nửa bên trái [left..mid] được sắp xếp tăng dần
		if nums[left] <= nums[mid] {
			// Kiểm tra xem target có nằm trong khoảng nửa trái hay không
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // Nhánh 2: Nửa bên phải [mid..right] được sắp xếp tăng dần
			// Kiểm tra xem target có nằm trong khoảng nửa phải hay không
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return -1
}
