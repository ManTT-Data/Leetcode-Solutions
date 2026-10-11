package dynamicprogramming

// Problem: 70. Climbing Stairs
// Difficulty: Easy
// Link: https://leetcode.com/problems/climbing-stairs/
//
// Time Complexity: O(n) - Duyệt qua một vòng lặp từ 3 đến n.
// Space Complexity: O(1) - Chỉ sử dụng hai biến để lưu trạng thái của hai bước trước đó.

// climbStairs tính số cách khác nhau để leo lên đỉnh của cầu thang gồm n bậc.
// Mỗi lần chỉ có thể leo 1 hoặc 2 bậc.
//
// Ý tưởng quy hoạch động (Dynamic Programming - Fibonacci Sequence):
// - Để bước lên bậc thứ i, ta chỉ có thể:
//   1. Bước 1 bậc từ bậc (i - 1)
//   2. Bước 2 bậc từ bậc (i - 2)
// - Do đó, tổng số cách để lên bậc i là: dp[i] = dp[i-1] + dp[i-2]
// - Cơ sở (Base cases):
//   dp[1] = 1 (1 cách: 1)
//   dp[2] = 2 (2 cách: 1+1, 2)
//
// Tối ưu không gian (Space Optimization):
// - Ta chỉ cần 2 biến `prev1` (đại diện cho dp[i-1]) và `prev2` (đại diện cho dp[i-2])
//   thay vì cấp phát cả mảng kích thước n, giảm bộ nhớ về O(1).
func climbStairs(n int) int {
	if n <= 2 {
		return n
	}

	prev2 := 1 // dp[1]
	prev1 := 2 // dp[2]

	for i := 3; i <= n; i++ {
		current := prev1 + prev2
		prev2 = prev1
		prev1 = current
	}

	return prev1
}
