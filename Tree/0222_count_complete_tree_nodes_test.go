package tree

import "testing"

// buildCompleteTree tạo Complete Binary Tree từ lát cắt số nguyên theo thứ tự level-order.
func buildCompleteTree(vals []int) *TreeNode {
	if len(vals) == 0 {
		return nil
	}

	nodes := make([]*TreeNode, len(vals))
	for i, v := range vals {
		nodes[i] = &TreeNode{Val: v}
	}

	for i := 0; i < len(vals); i++ {
		leftIdx := 2*i + 1
		rightIdx := 2*i + 2

		if leftIdx < len(vals) {
			nodes[i].Left = nodes[leftIdx]
		}
		if rightIdx < len(vals) {
			nodes[i].Right = nodes[rightIdx]
		}
	}

	return nodes[0]
}

func TestCountNodes(t *testing.T) {
	tests := []struct {
		name string
		vals []int
		want int
	}{
		{
			name: "Example 1 - 6 nodes",
			vals: []int{1, 2, 3, 4, 5, 6},
			want: 6,
		},
		{
			name: "Example 2 - Empty tree",
			vals: []int{},
			want: 0,
		},
		{
			name: "Example 3 - Single node",
			vals: []int{1},
			want: 1,
		},
		{
			name: "Two nodes - Root and Left child",
			vals: []int{1, 2},
			want: 2,
		},
		{
			name: "Three nodes - Perfect tree depth 2",
			vals: []int{1, 2, 3},
			want: 3,
		},
		{
			name: "Four nodes - Complete tree depth 3",
			vals: []int{1, 2, 3, 4},
			want: 4,
		},
		{
			name: "Seven nodes - Perfect tree depth 3",
			vals: []int{1, 2, 3, 4, 5, 6, 7},
			want: 7,
		},
		{
			name: "Large tree - 100 nodes",
			vals: func() []int {
				arr := make([]int, 100)
				for i := range arr {
					arr[i] = i + 1
				}
				return arr
			}(),
			want: 100,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildCompleteTree(tt.vals)
			got := countNodes(root)
			if got != tt.want {
				t.Errorf("countNodes() = %d; want %d", got, tt.want)
			}
		})
	}
}
