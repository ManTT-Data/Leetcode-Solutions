package twopointer

// Problem: 11. Container With Most Water
// Difficulty: Medium
// Link: https://leetcode.com/problems/container-with-most-water/
//
// Time Complexity: O(n) - Mỗi phần tử được duyệt tối đa một lần bởi hai con trỏ left và right.
// Space Complexity: O(1) - Chỉ sử dụng các biến con trỏ và biến lưu diện tích lớn nhất.

// maxArea tìm lượng nước tối đa mà bình chứa có thể tích trữ được.
// Sử dụng kỹ thuật Hai con trỏ (Two Pointers):
// - Đặt hai con trỏ ở hai đầu mảng: `left = 0`, `right = len(height) - 1` (chiều rộng tối đa).
// - Diện tích nước được tính bằng: (right - left) * min(height[left], height[right]).
// - Chiến lược di chuyển con trỏ:
//   Lượng nước bị giới hạn bởi cột thấp hơn. Khi ta thu hẹp khoảng cách (right - left giảm),
//   nếu giữ lại cột thấp hơn và dịch chuyển cột cao hơn, diện tích chắc chắn sẽ giảm.
//   Do đó, cơ hội duy nhất để tìm diện tích lớn hơn là luôn dịch chuyển con trỏ ở CỘT THẤP HƠN
//   với kỳ vọng tìm được cột cao hơn thay thế.
func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxWater := 0

	for left < right {
		width := right - left
		h := height[left]
		if height[right] < h {
			h = height[right]
		}

		currentWater := width * h
		if currentWater > maxWater {
			maxWater = currentWater
		}

		// Dịch con trỏ ở cột thấp hơn
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}

	return maxWater
}
