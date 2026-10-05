package array

// Problem: 977. Squares of a Sorted Array
// Difficulty: Easy
// Link: https://leetcode.com/problems/squares-of-a-sorted-array/
//
// Time Complexity: O(n) - Mỗi phần tử được xét đúng một lần với kỹ thuật hai con trỏ.
// Space Complexity: O(1) - Không tính mảng kết quả trả về.

// sortedSquares trả về mảng bình phương của các phần tử được sắp xếp theo thứ tự không giảm.
// Tiếp cận bằng Two Pointers: Do mảng đã sắp xếp, giá trị bình phương lớn nhất chỉ có thể
// nằm ở đầu mảng (số âm có trị tuyệt đối lớn) hoặc ở cuối mảng (số dương lớn).
// Do đó ta dùng hai con trỏ từ 2 đầu và điền kết quả từ cuối mảng về đầu (từ lớn nhất đến nhỏ nhất).
func sortedSquares(nums []int) []int {
	n := len(nums)
	result := make([]int, n)

	left, right := 0, n-1
	pos := n - 1

	for left <= right {
		leftSquare := nums[left] * nums[left]
		rightSquare := nums[right] * nums[right]

		if leftSquare > rightSquare {
			result[pos] = leftSquare
			left++
		} else {
			result[pos] = rightSquare
			right--
		}
		pos--
	}

	return result
}
