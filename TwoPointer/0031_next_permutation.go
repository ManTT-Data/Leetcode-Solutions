package twopointer

// Problem: 31. Next Permutation
// Difficulty: Medium
// Link: https://leetcode.com/problems/next-permutation/
//
// Time Complexity: O(n) - At most two passes over the array and one reverse operation.
// Space Complexity: O(1) - In-place modification with constant extra space.

// nextPermutation rearranges numbers into the lexicographically next greater permutation in-place.
func nextPermutation(nums []int) {
	n := len(nums)
	if n <= 1 {
		return
	}

	// Bước 1: Tìm chỉ số i lớn nhất sao cho nums[i] < nums[i+1] (quét từ phải sang trái)
	i := n - 2
	for i >= 0 && nums[i] >= nums[i+1] {
		i--
	}

	// Bước 2: Nếu tìm thấy i, tìm chỉ số j lớn nhất từ phải sang sao cho nums[j] > nums[i]
	if i >= 0 {
		j := n - 1
		for nums[j] <= nums[i] {
			j--
		}
		// Đổi chỗ nums[i] và nums[j]
		nums[i], nums[j] = nums[j], nums[i]
	}

	// Bước 3: Đảo ngược đoạn từ i+1 đến cuối mảng để đạt thứ tự nhỏ nhất
	reverse(nums, i+1, n-1)
}

func reverse(nums []int, start, end int) {
	for start < end {
		nums[start], nums[end] = nums[end], nums[start]
		start++
		end--
	}
}
