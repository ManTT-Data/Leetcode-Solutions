package twopointer

import (
	"math"
	"sort"
)

// Problem: 209. Minimum Size Subarray Sum
// Difficulty: Medium
// Link: https://leetcode.com/problems/minimum-size-subarray-sum/
//
// Time Complexity: O(n) - Mỗi phần tử được duyệt qua tối đa 2 lần (bởi con trỏ right và left).
// Space Complexity: O(1) - Chỉ sử dụng một vài biến con trỏ và biến tổng cố định.

// minSubArrayLen tìm độ dài nhỏ nhất của mảng con liên tiếp có tổng >= target.
// Sử dụng kỹ thuật Sliding Window (Hai con trỏ):
// - Con trỏ `right` mở rộng cửa sổ về bên phải để cộng dồn các phần tử vào `currentSum`.
// - Khi `currentSum >= target`, ta liên tục thu hẹp cửa sổ từ bên trái bằng con trỏ `left`
//   để tìm độ dài ngắn nhất có thể thỏa mãn điều kiện, sau đó trừ dần `nums[left]`.
func minSubArrayLen(target int, nums []int) int {
	n := len(nums)
	minLength := math.MaxInt
	currentSum := 0
	left := 0

	for right := 0; right < n; right++ {
		currentSum += nums[right]
		// Khi tổng cửa sổ đạt yêu cầu, cố gắng thu nhỏ cửa sổ từ bên trái
		for currentSum >= target {
			currentLength := right - left + 1
			if currentLength < minLength {
				minLength = currentLength
			}
			currentSum -= nums[left]
			left++
		}
	}

	if minLength == math.MaxInt {
		return 0
	}
	return minLength
}

// Follow-up: Giải thuật với độ phức tạp O(n log n) sử dụng Prefix Sum + Binary Search.
// Do các phần tử nums[i] đều là số nguyên dương, mảng prefix sum là dãy tăng nghiêm ngặt.
// Với mỗi vị trí i, ta tìm vị trí j nhỏ nhất sao cho:
//   prefix[j] - prefix[i] >= target  <=>  prefix[j] >= target + prefix[i]
// Ta có thể áp dụng tìm kiếm nhị phân để tìm j trong O(log n).
func minSubArrayLenBinarySearch(target int, nums []int) int {
	n := len(nums)
	minLength := math.MaxInt

	// Xây dựng mảng prefix sum với prefix[0] = 0
	prefix := make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = prefix[i] + nums[i]
	}

	for i := 0; i < n; i++ {
		toFind := target + prefix[i]
		// Tìm vị trí j nhỏ nhất trong prefix sao cho prefix[j] >= toFind
		j := sort.Search(len(prefix), func(idx int) bool {
			return prefix[idx] >= toFind
		})

		if j < len(prefix) {
			currentLength := j - i
			if currentLength < minLength {
				minLength = currentLength
			}
		}
	}

	if minLength == math.MaxInt {
		return 0
	}
	return minLength
}
