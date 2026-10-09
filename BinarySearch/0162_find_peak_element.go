package binarysearch

// Problem: 162. Find Peak Element
// Difficulty: Medium
// Link: https://leetcode.com/problems/find-peak-element/
//
// Time Complexity: O(log n) - Thuật toán tìm kiếm nhị phân chia đôi không gian tìm kiếm ở mỗi bước.
// Space Complexity: O(1) - Sử dụng bộ nhớ hằng số O(1).

// findPeakElement tìm chỉ số của một phần tử đỉnh (peak element) bất kỳ trong mảng.
// Phần tử đỉnh là phần tử lớn hơn các phần tử kề cạnh nó (với giả định nums[-1] = nums[n] = -∞).
func findPeakElement(nums []int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := left + (right-left)/2

		// So sánh phần tử mid với phần tử kế tiếp mid + 1
		if nums[mid] > nums[mid+1] {
			// Sườn dốc đang đi xuống về phía bên phải.
			// Do nums[-1] = -∞, chắc chắn tồn tại ít nhất một đỉnh ở nửa bên trái hoặc chính là mid.
			right = mid
		} else {
			// Sườn dốc đang đi lên về phía bên phải (nums[mid] < nums[mid+1]).
			// Do nums[n] = -∞, chắc chắn tồn tại ít nhất một đỉnh ở nửa bên phải (từ mid + 1 trở đi).
			left = mid + 1
		}
	}

	return left
}
