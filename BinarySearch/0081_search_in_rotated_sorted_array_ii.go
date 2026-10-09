package binarysearch

// Problem: 81. Search in Rotated Sorted Array II
// Difficulty: Medium
// Link: https://leetcode.com/problems/search-in-rotated-sorted-array-ii/
//
// Time Complexity: O(log n) on average, O(n) in the worst case (when many duplicates exist).
// Space Complexity: O(1) - Constant auxiliary space.

// searchII determines if target exists in a rotated sorted array with duplicates.
// (On LeetCode, this function is named `search`).
func searchII(nums []int, target int) bool {
	left, right := 0, len(nums)-1

	for left <= right {
		mid := left + (right-left)/2

		if nums[mid] == target {
			return true
		}

		// Khi nums[left] == nums[mid], ta không thể xác định nửa nào được sắp xếp chuẩn
		// Do nums[mid] != target, nên nums[left] chắc chắn != target, an toàn bỏ qua left++
		if nums[left] == nums[mid] {
			left++
			continue
		}

		// Nhánh 1: Nửa bên trái [left..mid] có thứ tự tăng dần nghiêm ngặt
		if nums[left] < nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else { // Nhánh 2: Nửa bên phải [mid..right] có thứ tự tăng dần
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}

	return false
}
