package twopointer

// Problem: 287. Find the Duplicate Number
// Difficulty: Medium
// Link: https://leetcode.com/problems/find-the-duplicate-number/
//
// Time Complexity: O(n) - Thuật toán Floyd phát hiện chu trình tìm điểm gặp nhau và điểm bắt đầu chu trình trong thời gian tuyến tính.
// Space Complexity: O(1) - Không sử dụng thêm mảng phụ hay thay đổi mảng ban đầu, chỉ dùng hai con trỏ rùa và thỏ.

// findDuplicate tìm số bị lặp lại trong mảng gồm n + 1 số nguyên trong khoảng [1, n].
// Thỏa mãn các ràng buộc khắt khe của đề bài:
// - Không thay đổi mảng nums.
// - Chỉ dùng bộ nhớ O(1).
//
// Ý tưởng giải thuật: Floyd's Tortoise and Hare (Phát hiện chu trình bằng Hai con trỏ nhanh - chậm):
// 1. Chuyển đổi mô hình:
//    - Coi mảng nums như một danh sách liên kết: mỗi chỉ số i trỏ tới nút có chỉ số là nums[i].
//    - Vì các giá trị đều nằm trong đoạn [1, n], nên chỉ số 0 không bao giờ là đích đến của bất kỳ nút nào.
//    - Do có n + 1 nút mà chỉ có n giá trị đích, theo nguyên lý Dirichlet (Pigeonhole Principle),
//      chắc chắn có ít nhất hai nút trỏ tới cùng một nút đích -> tạo thành chu trình (cycle).
//    - Điểm bắt đầu của chu trình (Cycle Entrance) chính là số bị lặp lại!
//
// 2. Giai đoạn 1 (Tìm điểm gặp nhau trong chu trình):
//    - Con trỏ `slow` di chuyển 1 bước: slow = nums[slow]
//    - Con trỏ `fast` di chuyển 2 bước: fast = nums[nums[fast]]
//    - Lặp lại cho đến khi slow == fast.
//
// 3. Giai đoạn 2 (Tìm điểm bắt đầu chu trình - số lặp):
//    - Đặt lại `slow = 0`.
//    - Cả hai con trỏ `slow` và `fast` cùng di chuyển 1 bước mỗi lần cho đến khi gặp nhau.
//    - Điểm gặp nhau chính là số bị lặp lại.
func findDuplicate(nums []int) int {
	slow := nums[0]
	fast := nums[nums[0]]

	// Giai đoạn 1: Tìm điểm giao nhau trong chu trình
	for slow != fast {
		slow = nums[slow]
		fast = nums[nums[fast]]
	}

	// Giai đoạn 2: Tìm điểm bắt đầu chu trình
	slow = 0
	for slow != fast {
		slow = nums[slow]
		fast = nums[fast]
	}

	return slow
}
