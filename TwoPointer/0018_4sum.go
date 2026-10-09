package twopointer

import "sort"

// Problem: 18. 4Sum
// Difficulty: Medium
// Link: https://leetcode.com/problems/4sum/
//
// Time Complexity: O(n^3) - Two nested loops with a two-pointer inner scan.
// Space Complexity: O(1) or O(log n) auxiliary space used by sorting.

// fourSum returns all unique quadruplets [nums[a], nums[b], nums[c], nums[d]] that sum up to target.
func fourSum(nums []int, target int) [][]int {
	var result [][]int
	n := len(nums)
	if n < 4 {
		return result
	}

	sort.Ints(nums)
	t := int64(target)

	for i := 0; i < n-3; i++ {
		// Nhánh cắt tỉa 1: tổng 4 số nhỏ nhất khả dĩ lớn hơn target -> không thể có nghiệm nào nữa
		if int64(nums[i])+int64(nums[i+1])+int64(nums[i+2])+int64(nums[i+3]) > t {
			break
		}

		// Nhánh cắt tỉa 2: tổng của nums[i] với 3 số lớn nhất vẫn nhỏ hơn target -> tăng nums[i]
		if int64(nums[i])+int64(nums[n-3])+int64(nums[n-2])+int64(nums[n-1]) < t {
			continue
		}

		// Bỏ qua giá trị trùng lặp ở vị trí số thứ nhất
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

		for j := i + 1; j < n-2; j++ {
			// Nhánh cắt tỉa tương tự cho vị trí số thứ hai
			if int64(nums[i])+int64(nums[j])+int64(nums[j+1])+int64(nums[j+2]) > t {
				break
			}
			if int64(nums[i])+int64(nums[j])+int64(nums[n-2])+int64(nums[n-1]) < t {
				continue
			}

			// Bỏ qua giá trị trùng lặp ở vị trí số thứ hai
			if j > i+1 && nums[j] == nums[j-1] {
				continue
			}

			left, right := j+1, n-1
			for left < right {
				sum := int64(nums[i]) + int64(nums[j]) + int64(nums[left]) + int64(nums[right])

				switch {
				case sum == t:
					result = append(result, []int{nums[i], nums[j], nums[left], nums[right]})

					// Bỏ qua các phần tử trùng lặp ở left và right
					for left < right && nums[left] == nums[left+1] {
						left++
					}
					for left < right && nums[right] == nums[right-1] {
						right--
					}

					left++
					right--
				case sum < t:
					left++
				default: // sum > t
					right--
				}
			}
		}
	}

	return result
}
