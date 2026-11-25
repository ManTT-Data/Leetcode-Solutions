package tree

// Problem: 102. Binary Tree Level Order Traversal
// Difficulty: Medium
// Link: https://leetcode.com/problems/binary-tree-level-order-traversal/
//
// Time Complexity: O(n) - Mỗi nút trong cây nhị phân được đưa vào hàng đợi và xử lý đúng 1 lần.
// Space Complexity: O(n) - Hàng đợi BFS chứa tối đa số nút ở tầng rộng nhất (tối đa n/2 nút).

// levelOrder duyệt cây nhị phân theo từng tầng (Level Order / BFS) từ trái sang phải.
//
// Ý tưởng giải thuật:
// 1. Nếu `root == nil`, trả về mảng rỗng `[][]int{}`.
// 2. Dùng hàng đợi (queue) khởi tạo với nút gốc `root`.
// 3. Trong khi hàng đợi chưa rỗng:
//    - Lưu lại kích thước hiện tại `levelSize = len(queue)` tương ứng với số nút của tầng đang xét.
//    - Khởi tạo slice `level` với độ dài `levelSize`, lần lượt lấy `levelSize` nút đầu hàng đợi:
//      + Ghi nhận `node.Val` vào `level[i]`.
//      + Nếu có con trái/phải thì đưa vào cuối hàng đợi để xét ở tầng kế tiếp.
//    - Thêm `level` vào kết quả `res`.
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	var res [][]int
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelSize := len(queue)
		level := make([]int, levelSize)

		for i := 0; i < levelSize; i++ {
			node := queue[0]
			queue = queue[1:]

			level[i] = node.Val
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}

		res = append(res, level)
	}

	return res
}
