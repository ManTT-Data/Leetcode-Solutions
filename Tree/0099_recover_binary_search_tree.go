package tree

// Problem: 99. Recover Binary Search Tree
// Difficulty: Medium
// Link: https://leetcode.com/problems/recover-binary-search-tree/
//
// Time Complexity: O(n) - Duyệt cây theo Morris In-order Traversal, mỗi cạnh được đi qua tối đa 3 lần.
// Space Complexity: O(1) - Sử dụng Morris Traversal (luồng liên kết tạm thời), không tốn bộ nhớ ngăn xếp đệ quy O(h).

// recoverTree khôi phục cây nhị phân tìm kiếm (BST) khi có đúng 2 nút bị đổi chỗ nhầm.
// Yêu cầu: Không làm thay đổi cấu trúc cây, chỉ đổi lại giá trị 2 nút, bộ nhớ phụ tối ưu O(1).
//
// Ý tưởng giải thuật (Morris In-order Traversal):
// 1. Đặc điểm BST:
//    - Thứ tự duyệt In-order (Trái -> Gốc -> Phải) của một BST hợp lệ luôn là một dãy số tăng dần nghiêm ngặt.
//    - Khi 2 nút bị hoán đổi giá trị, dãy in-order sẽ xuất hiện sự nghịch thế (prev.Val > curr.Val):
//      + Nếu 2 nút KHÔNG kề nhau: Sẽ có 2 lần nghịch thế.
//        Lần 1: prev.Val > curr.Val -> Nút thứ nhất bị sai là `first = prev` (giá trị lớn hơn).
//        Lần 2: prev.Val > curr.Val -> Nút thứ hai bị sai là `second = curr` (giá trị nhỏ hơn).
//      + Nếu 2 nút KỀ nhau: Chỉ có 1 lần nghịch thế duy nhất.
//        `first = prev` và `second = curr`.
//
// 2. Morris Traversal để đạt O(1) Space:
//    - Thay vì dùng đệ quy hay Stack tốn O(h) bộ nhớ (với h là chiều cao cây, worst case O(n)),
//      Morris Traversal tận dụng con trỏ phải còn trống của nút tiền nhiệm (predecessor)
//      để trỏ ngược về nút gốc hiện tại (tạo cầu nối tạm thời).
//    - Khi duyệt xong cây con trái, cầu nối được xóa đi để phục hồi cấu trúc cây ban đầu.
func recoverTree(root *TreeNode) {
	var first, second, prev *TreeNode
	curr := root

	for curr != nil {
		if curr.Left == nil {
			// Xử lý nút curr hiện tại
			if prev != nil && prev.Val > curr.Val {
				if first == nil {
					first = prev
				}
				second = curr
			}
			prev = curr
			curr = curr.Right
		} else {
			// Tìm nút tiền nhiệm predecessor (nút phải cùng của cây con bên trái)
			pred := curr.Left
			for pred.Right != nil && pred.Right != curr {
				pred = pred.Right
			}

			if pred.Right == nil {
				// Tạo cầu nối (thread) từ predecessor về curr
				pred.Right = curr
				curr = curr.Left
			} else {
				// Đã duyệt xong cây con trái, hủy bỏ cầu nối
				pred.Right = nil

				// Xử lý nút curr hiện tại
				if prev != nil && prev.Val > curr.Val {
					if first == nil {
						first = prev
					}
					second = curr
				}
				prev = curr
				curr = curr.Right
			}
		}
	}

	// Hoán đổi lại giá trị của hai nút bị đặt sai vị trí
	if first != nil && second != nil {
		first.Val, second.Val = second.Val, first.Val
	}
}
