package array

// Problem: 153. Find Minimum in Rotated Sorted Array
// Difficulty: Medium
// Link: https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/
//
// Time Complexity: O(log n) - Thuật toán tìm kiếm nhị phân chia đôi không gian tìm kiếm ở mỗi bước.
// Space Complexity: O(1) - Không sử dụng thêm bộ nhớ phụ thuộc vào kích thước mảng.

// findMin tìm phần tử nhỏ nhất trong mảng đã sắp xếp và bị xoay.
func findMin(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		// Nếu phần tử ở giữa lớn hơn phần tử ngoài cùng bên phải,
		// thì điểm cực tiểu (minimum) chắc chắn nằm ở nửa bên phải (từ mid + 1 đến right).
		if nums[mid] > nums[right] {
			left = mid + 1
		} else {
			// Ngược lại, điểm cực tiểu nằm ở nửa bên trái hoặc chính là mid (từ left đến mid).
			right = mid
		}
	}

	return nums[left]
}
