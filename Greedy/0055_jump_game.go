package greedy

// Problem: 55. Jump Game
// Difficulty: Medium
// Link: https://leetcode.com/problems/jump-game/
//
// Time Complexity: O(n) - Duyệt mảng một lần duy nhất, dừng sớm nếu đã chạm tới đích.
// Space Complexity: O(1) - Chỉ sử dụng một biến `maxReach` để theo dõi chỉ số xa nhất có thể nhảy tới.

// canJump kiểm tra xem có thể nhảy từ chỉ số đầu tiên (0) đến chỉ số cuối cùng (len(nums)-1) hay không.
//
// Ý tưởng giải thuật (Tham lam - Greedy):
// - Duy trì biến `maxReach` là vị trí xa nhất có thể đi tới được tính đến thời điểm hiện tại.
// - Duyệt từng chỉ số `i` từ 0 đến `len(nums)-1`:
//   + Nếu `i > maxReach`: Không thể bước tới được vị trí `i`, trả về `false`.
//   + Cập nhật `maxReach = max(maxReach, i + nums[i])`.
//   + Nếu `maxReach >= len(nums)-1`: Chắc chắn tới được đích, trả về `true` sớm.
func canJump(nums []int) bool {
	maxReach := 0
	lastIdx := len(nums) - 1

	for i := 0; i <= lastIdx; i++ {
		if i > maxReach {
			return false
		}
		if reach := i + nums[i]; reach > maxReach {
			maxReach = reach
		}
		if maxReach >= lastIdx {
			return true
		}
	}

	return true
}
