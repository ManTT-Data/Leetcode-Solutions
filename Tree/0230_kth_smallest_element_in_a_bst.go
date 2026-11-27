package tree

// Problem: 230. Kth Smallest Element in a BST
// Difficulty: Medium
// Link: https://leetcode.com/problems/kth-smallest-element-in-a-bst/
//
// Time Complexity: O(H + k) - Đi xuống nhánh trái nhất mất O(H) (với H là chiều cao cây), sau đó duyệt tối đa k nút.
// Space Complexity: O(H) - Ngăn xếp lưu đường đi từ gốc xuống lá, tối đa bằng chiều cao cây H.

// kthSmallest tìm phần tử nhỏ thứ k (1-indexed) trong cây nhị phân tìm kiếm (BST).
//
// Ý tưởng giải thuật (Iterative In-order Traversal):
// - Thứ tự duyệt In-order (Trái -> Gốc -> Phải) trên BST cho ra các giá trị tăng dần.
// - Thay vì duyệt toàn bộ cây, dùng ngăn xếp (stack) mô phỏng In-order:
//   1. Đẩy toàn bộ các nút dọc theo nhánh trái của `curr` vào stack.
//   2. Lấy nút đỉnh stack ra xét, giảm `k` đi 1.
//   3. Nếu `k == 0`, trả về ngay `node.Val` (dừng sớm).
//   4. Ngược lại, chuyển `curr = node.Right` và lặp lại.
func kthSmallest(root *TreeNode, k int) int {
	var stack []*TreeNode
	curr := root

	for curr != nil || len(stack) > 0 {
		for curr != nil {
			stack = append(stack, curr)
			curr = curr.Left
		}

		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		k--
		if k == 0 {
			return node.Val
		}

		curr = node.Right
	}

	return 0
}
