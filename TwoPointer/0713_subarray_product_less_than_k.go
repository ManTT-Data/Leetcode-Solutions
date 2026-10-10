package twopointer

// Problem: 713. Subarray Product Less Than K
// Difficulty: Medium
// Link: https://leetcode.com/problems/subarray-product-less-than-k/
//
// Time Complexity: O(n) - Cả hai con trỏ left và right đều duyệt qua mảng tối đa 1 lần.
// Space Complexity: O(1) - Chỉ sử dụng các biến con trỏ, tích và biến đếm cố định.

// numSubarrayProductLessThanK đếm số lượng mảng con liên tiếp có tích các phần tử nhỏ hơn k.
// Sử dụng kỹ thuật Sliding Window (Hai con trỏ):
// - Vì mọi phần tử nums[i] >= 1, tích của bất kỳ mảng con không rỗng nào cũng luôn >= 1.
//   Do đó, nếu k <= 1 thì không thể có mảng con nào có tích strictly less than k (< k) -> return 0.
// - Con trỏ `right` mở rộng cửa sổ về bên phải và nhân dồn nums[right] vào `product`.
// - Khi `product >= k`, ta liên tục thu hẹp cửa sổ từ bên trái bằng con trỏ `left` (chia dần nums[left]).
// - Với mỗi vị trí `right`, số lượng mảng con hợp lệ kết thúc tại `right` chính là: (right - left + 1).
func numSubarrayProductLessThanK(nums []int, k int) int {
	if k <= 1 {
		return 0
	}

	result := 0
	product := 1
	left := 0

	for right := 0; right < len(nums); right++ {
		product *= nums[right]

		// Thu hẹp cửa sổ khi tích vượt quá hoặc bằng k
		for product >= k && left <= right {
			product /= nums[left]
			left++
		}

		// Số lượng mảng con hợp lệ kết thúc tại right là (right - left + 1)
		result += right - left + 1
	}

	return result
}
