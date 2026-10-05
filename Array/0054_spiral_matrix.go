package array

// Problem: 54. Spiral Matrix
// Difficulty: Medium
// Link: https://leetcode.com/problems/spiral-matrix/
//
// Time Complexity: O(m * n) - Mỗi phần tử trong ma trận được duyệt đúng một lần.
// Space Complexity: O(1) - Không tính bộ nhớ lưu mảng kết quả trả về.

// spiralOrder trả về tất cả các phần tử của ma trận theo thứ tự xoắn ốc (spiral order).
func spiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return []int{}
	}

	m, n := len(matrix), len(matrix[0])
	result := make([]int, 0, m*n)

	top, bot := 0, m-1
	left, right := 0, n-1

	for top <= bot && left <= right {
		// 1. Duyệt từ trái sang phải trên hàng top
		for col := left; col <= right; col++ {
			result = append(result, matrix[top][col])
		}
		top++

		// 2. Duyệt từ trên xuống dưới trên cột right
		for row := top; row <= bot; row++ {
			result = append(result, matrix[row][right])
		}
		right--

		// 3. Duyệt từ phải sang trái trên hàng bot
		// Cần kiểm tra top <= bot để tránh duyệt lặp lại khi chỉ còn 1 hàng duy nhất
		if top <= bot {
			for col := right; col >= left; col-- {
				result = append(result, matrix[bot][col])
			}
			bot--
		}

		// 4. Duyệt từ dưới lên trên trên cột left
		// Cần kiểm tra left <= right để tránh duyệt lặp lại khi chỉ còn 1 cột duy nhất
		if left <= right {
			for row := bot; row >= top; row-- {
				result = append(result, matrix[row][left])
			}
			left++
		}
	}

	return result
}
