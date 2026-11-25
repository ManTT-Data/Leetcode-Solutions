package dynamicprogramming

// Problem: 198. House Robber
// Difficulty: Medium
// Link: https://leetcode.com/problems/house-robber/
//
// Time Complexity: O(n) - Duyệt qua mảng một lần duy nhất.
// Space Complexity: O(1) - Tối ưu hóa không gian từ bảng DP 1 chiều xuống 2 biến lưu trạng thái trước đó.

// rob tính tổng số tiền lớn nhất có thể lấy được mà không trộm 2 nhà kề nhau.
//
// Ý tưởng giải thuật (Quy hoạch động - Dynamic Programming):
// Tại ngôi nhà thứ i với số tiền `nums[i]`, có 2 lựa chọn:
// - Không trộm nhà i: tổng tiền giữ nguyên bằng kết quả tối đa đến nhà i-1 (`prev1`).
// - Trộm nhà i: không được trộm nhà i-1, nên tổng tiền là `prev2 + nums[i]`
//   (với `prev2` là kết quả tối đa đến nhà i-2).
//
// Công thức truy hồi:
//   curr = max(prev1, prev2 + nums[i])
// Sau mỗi bước, cập nhật `prev2 = prev1`, `prev1 = curr`.
func rob(nums []int) int {
	prev2, prev1 := 0, 0

	for _, money := range nums {
		take := prev2 + money
		curr := prev1
		if take > curr {
			curr = take
		}
		prev2 = prev1
		prev1 = curr
	}

	return prev1
}
