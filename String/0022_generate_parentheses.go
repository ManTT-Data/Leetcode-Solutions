package string_problems

// Problem: 22. Generate Parentheses
// Difficulty: Medium
// Link: https://leetcode.com/problems/generate-parentheses/
//
// Time Complexity: O(4^n / sqrt(n)) - Số lượng tổ hợp ngoặc hợp lệ là số Catalan thứ n: C_n = (1 / (n + 1)) * (2n choose n).
// Space Complexity: O(n) - Độ sâu tối đa của ngăn xếp đệ quy là 2n.

// generateParenthesis sinh ra tất cả các chuỗi ngoặc hợp lệ từ n cặp dấu ngoặc.
// Sử dụng kỹ thuật Quay lui (Backtracking):
// - Ta duy trì biến `open` (số ngoặc mở đã đặt) và `close` (số ngoặc đóng đã đặt).
// - Điều kiện thêm ngoặc mở '(': open < n
// - Điều kiện thêm ngoặc đóng ')': close < open (để đảm bảo ngoặc đóng không vượt quá ngoặc mở)
// - Khi độ dài chuỗi đạt 2 * n, ta thu được một chuỗi ngoặc hợp lệ hoàn chỉnh.
func generateParenthesis(n int) []string {
	var result []string
	current := make([]byte, 0, 2*n)

	var backtrack func(open, close int)
	backtrack = func(open, close int) {
		if len(current) == 2*n {
			result = append(result, string(current))
			return
		}

		// Nếu vẫn còn ngoặc mở để đặt
		if open < n {
			current = append(current, '(')
			backtrack(open+1, close)
			current = current[:len(current)-1] // backtrack
		}

		// Chỉ đặt ngoặc đóng khi số ngoặc đóng ít hơn số ngoặc mở
		if close < open {
			current = append(current, ')')
			backtrack(open, close+1)
			current = current[:len(current)-1] // backtrack
		}
	}

	backtrack(0, 0)
	return result
}
