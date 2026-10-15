package tree

// Problem: 222. Count Complete Tree Nodes
// Difficulty: Medium
// Link: https://leetcode.com/problems/count-complete-tree-nodes/
//
// Time Complexity: O((log n)^2) - Vòng lặp duyệt qua tối đa h = log n tầng, tại mỗi tầng tính chiều cao mất O(log n).
// Space Complexity: O(1) - Cài đặt lặp tối ưu, chỉ sử dụng con trỏ và biến đếm, không tốn thêm bộ nhớ.

// TreeNode định nghĩa một nút trong cây nhị phân.
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// countNodes đếm tổng số nút của cây nhị phân gần hoàn chỉnh (Complete Binary Tree).
// Yêu cầu đề bài: Thiết kế thuật toán có độ phức tạp thời gian nhỏ hơn O(n).
//
// Ý tưởng giải thuật:
// Trong Complete Binary Tree:
// - Mọi tầng đều được lấp đầy, ngoại trừ có thể là tầng cuối cùng (các nút dồn hết về bên trái).
// - Nếu ta biết một cây con là Perfect Binary Tree (cây nhị phân hoàn hảo) có chiều cao h,
//   số nút của nó có thể tính ngay lập tức trong O(1) theo công thức: 2^h - 1.
//
// Thuật toán:
// 1. Với nút hiện tại `curr`:
//    - Đo chiều cao `leftH` của cây con trái (bằng cách đi theo nhánh trái nhất: curr.Left -> Left -> ...).
//    - Đo chiều cao `rightH` của cây con phải (cũng đi theo nhánh trái nhất: curr.Right -> Left -> ...).
//
// 2. So sánh leftH và rightH:
//    - Trường hợp 1: leftH == rightH
//      Nghĩa là tầng cuối cùng đã lấp đầy toàn bộ cây con bên trái và đang tràn sang cây con bên phải.
//      Do đó, cây con bên trái chắc chắn là Perfect Binary Tree có chiều cao leftH.
//      Số nút của cây con trái (2^leftH - 1) cộng thêm nút `curr` (1 nút) = 2^leftH (tức 1 << leftH).
//      Ta cộng (1 << leftH) vào tổng số nút và tiếp tục duyệt sang cây con bên phải: `curr = curr.Right`.
//
//    - Trường hợp 2: leftH > rightH
//      Tầng cuối cùng chưa chạm tới cây con bên phải (đang dừng ở cây con bên trái).
//      Điều này chứng tỏ cây con bên phải là Perfect Binary Tree có chiều cao rightH = leftH - 1.
//      Số nút của cây con phải (2^rightH - 1) cộng thêm nút `curr` (1 nút) = 2^rightH (tức 1 << rightH).
//      Ta cộng (1 << rightH) vào tổng số nút và tiếp tục duyệt sang cây con bên trái: `curr = curr.Left`.
//
// 3. Lặp lại cho đến khi `curr == nil`.
func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}

	count := 0
	curr := root

	for curr != nil {
		leftH := getLeftHeight(curr.Left)
		rightH := getLeftHeight(curr.Right)

		if leftH == rightH {
			// Cây con trái là Perfect Binary Tree chiều cao leftH
			// Số nút của cây trái + node curr = (2^leftH - 1) + 1 = 1 << leftH
			count += 1 << leftH
			curr = curr.Right
		} else {
			// Cây con phải là Perfect Binary Tree chiều cao rightH
			// Số nút của cây phải + node curr = (2^rightH - 1) + 1 = 1 << rightH
			count += 1 << rightH
			curr = curr.Left
		}
	}

	return count
}

// getLeftHeight tính chiều cao của cây bằng cách đi dọc theo cạnh trái nhất.
func getLeftHeight(node *TreeNode) int {
	h := 0
	for node != nil {
		h++
		node = node.Left
	}
	return h
}
