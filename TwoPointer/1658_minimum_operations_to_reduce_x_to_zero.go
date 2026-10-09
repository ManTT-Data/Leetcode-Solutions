package twopointer

// Problem: 1658. Minimum Operations to Reduce X to Zero
// Difficulty: Medium
// Link: https://leetcode.com/problems/minimum-operations-to-reduce-x-to-zero/
//
// Time Complexity: O(n) - Mỗi phần tử được duyệt qua tối đa 2 lần bởi hai con trỏ left và right.
// Space Complexity: O(1) - Chỉ sử dụng các biến con trỏ và biến tổng cố định.

// minOperations trả về số thao tác ít nhất để giảm x về 0 bằng cách loại bỏ các phần tử ở 2 đầu mảng.
// Ý tưởng đổi góc nhìn (Complementary Sliding Window):
// - Việc loại bỏ tiền tố và hậu tố có tổng bằng x với số phần tử ít nhất
//   tương đương với việc tìm mảng con liên tiếp ở giữa có tổng bằng:
//     target = totalSum - x
//   sao cho độ dài của mảng con này là LỚN NHẤT (maxLen).
// - Kết quả sẽ là: len(nums) - maxLen.
func minOperations(nums []int, x int) int {
	totalSum := 0
	for _, num := range nums {
		totalSum += num
	}

	target := totalSum - x

	// Trường hợp tổng cả mảng nhỏ hơn x: không thể giảm x về 0
	if target < 0 {
		return -1
	}

	// Trường hợp tổng cả mảng đúng bằng x: phải xóa toàn bộ mảng
	if target == 0 {
		return len(nums)
	}

	n := len(nums)
	maxLen := -1
	currentSum := 0
	left := 0

	for right := 0; right < n; right++ {
		currentSum += nums[right]

		// Khi tổng cửa sổ vượt quá target, thu hẹp từ bên trái
		for currentSum > target && left <= right {
			currentSum -= nums[left]
			left++
		}

		// Khi tìm thấy mảng con có tổng đúng bằng target
		if currentSum == target {
			currentLength := right - left + 1
			if currentLength > maxLen {
				maxLen = currentLength
			}
		}
	}

	if maxLen == -1 {
		return -1
	}

	return n - maxLen
}
