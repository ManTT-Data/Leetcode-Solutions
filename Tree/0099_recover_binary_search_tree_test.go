package tree

import (
	"reflect"
	"testing"
)

// intPtr helper để tạo con trỏ *int từ số nguyên cho test cases có giá trị null
func intPtr(val int) *int {
	return &val
}

// buildTreeFromLevelOrder dựng cây nhị phân từ mảng level-order có chứa null (*int)
func buildTreeFromLevelOrder(vals []*int) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}

	root := &TreeNode{Val: *vals[0]}
	queue := []*TreeNode{root}
	i := 1

	for len(queue) > 0 && i < len(vals) {
		curr := queue[0]
		queue = queue[1:]

		// Con trái
		if i < len(vals) {
			if vals[i] != nil {
				curr.Left = &TreeNode{Val: *vals[i]}
				queue = append(queue, curr.Left)
			}
			i++
		}

		// Con phải
		if i < len(vals) {
			if vals[i] != nil {
				curr.Right = &TreeNode{Val: *vals[i]}
				queue = append(queue, curr.Right)
			}
			i++
		}
	}

	return root
}

// treeToInOrder chuyển đổi cây thành mảng theo thứ tự duyệt In-order
func treeToInOrder(root *TreeNode) []int {
	var res []int
	var dfs func(node *TreeNode)
	dfs = func(node *TreeNode) {
		if node == nil {
			return
		}
		dfs(node.Left)
		res = append(res, node.Val)
		dfs(node.Right)
	}
	dfs(root)
	return res
}

func TestRecoverTree(t *testing.T) {
	tests := []struct {
		name        string
		input       []*int
		wantInOrder []int
	}{
		{
			name:        "Example 1: [1, 3, null, null, 2]",
			input:       []*int{intPtr(1), intPtr(3), nil, nil, intPtr(2)},
			wantInOrder: []int{1, 2, 3},
		},
		{
			name:        "Example 2: [3, 1, 4, null, null, 2]",
			input:       []*int{intPtr(3), intPtr(1), intPtr(4), nil, nil, intPtr(2)},
			wantInOrder: []int{1, 2, 3, 4},
		},
		{
			name:        "Two adjacent swapped nodes: root and left child",
			input:       []*int{intPtr(2), intPtr(3), intPtr(1)},
			wantInOrder: []int{1, 2, 3},
		},
		{
			name:        "Two non-adjacent swapped nodes: leftmost and rightmost",
			input:       []*int{intPtr(4), intPtr(7), intPtr(6), intPtr(1), intPtr(3), intPtr(5), intPtr(2)},
			wantInOrder: []int{1, 2, 3, 4, 5, 6, 7},
		},
		{
			name:        "Simple two nodes",
			input:       []*int{intPtr(2), intPtr(1)},
			wantInOrder: []int{1, 2},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeFromLevelOrder(tt.input)
			recoverTree(root)
			got := treeToInOrder(root)
			if !reflect.DeepEqual(got, tt.wantInOrder) {
				t.Errorf("recoverTree() resulting in-order = %v; want %v", got, tt.wantInOrder)
			}
		})
	}
}
