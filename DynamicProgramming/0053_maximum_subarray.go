package dynamicprogramming

// Problem: 53. Maximum Subarray
// Difficulty: Medium
// Link: https://leetcode.com/problems/maximum-subarray/
//
// Time Complexity: O(n) - Duyệt qua mảng một lần với thuật toán Kadane.
// Space Complexity: O(1) - Chỉ sử dụng hai biến lưu tổng hiện tại và tổng lớn nhất.

// maxSubArray tìm mảng con liên tiếp có tổng lớn nhất.
// Sử dụng thuật toán Kadane (Quy hoạch động tối ưu không gian):
// - Tại mỗi vị trí i, ta quyết định:
//   1. Bắt đầu một mảng con mới tại nums[i]
//   2. Hay cộng nums[i] vào mảng con trước đó (currentSum + nums[i])
// - Công thức: currentSum = max(nums[i], currentSum + nums[i])
// - maxSubArray = max(maxSubArray, currentSum)
func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := nums[0]

	for i := 1; i < len(nums); i++ {
		// Nếu currentSum < 0, bắt đầu mảng con mới từ nums[i] sẽ có lợi hơn
		if currentSum < 0 {
			currentSum = nums[i]
		} else {
			currentSum += nums[i]
		}

		if currentSum > maxSum {
			maxSum = currentSum
		}
	}

	return maxSum
}

// Follow-up: Giải thuật Chia để trị (Divide and Conquer).
// Độ phức tạp thời gian: O(n) tương tự cấu trúc Segment Tree (hoặc O(n log n) nếu chia đơn giản).
// Mỗi đoạn [l, r] duy trì 4 thông tin:
// - totalSum: tổng tất cả các phần tử trong đoạn
// - maxPrefix: tổng tiền tố lớn nhất
// - maxSuffix: tổng hậu tố lớn nhất
// - maxSub: tổng mảng con lớn nhất trong đoạn
type segment struct {
	totalSum  int
	maxPrefix int
	maxSuffix int
	maxSub    int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func pushUp(left, right segment) segment {
	totalSum := left.totalSum + right.totalSum
	maxPrefix := max(left.maxPrefix, left.totalSum+right.maxPrefix)
	maxSuffix := max(right.maxSuffix, right.totalSum+left.maxSuffix)
	maxSub := max(max(left.maxSub, right.maxSub), left.maxSuffix+right.maxPrefix)

	return segment{
		totalSum:  totalSum,
		maxPrefix: maxPrefix,
		maxSuffix: maxSuffix,
		maxSub:    maxSub,
	}
}

func divideAndConquer(nums []int, l, r int) segment {
	if l == r {
		return segment{
			totalSum:  nums[l],
			maxPrefix: nums[l],
			maxSuffix: nums[l],
			maxSub:    nums[l],
		}
	}

	mid := l + (r-l)/2
	leftSeg := divideAndConquer(nums, l, mid)
	rightSeg := divideAndConquer(nums, mid+1, r)

	return pushUp(leftSeg, rightSeg)
}

// maxSubArrayDivideAndConquer giải bài toán bằng phương pháp chia để trị O(n).
func maxSubArrayDivideAndConquer(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	return divideAndConquer(nums, 0, len(nums)-1).maxSub
}
