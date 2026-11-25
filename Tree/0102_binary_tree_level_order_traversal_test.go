package tree

import (
	"reflect"
	"testing"
)

func TestLevelOrder(t *testing.T) {
	tests := []struct {
		name  string
		input []*int
		want  [][]int
	}{
		{
			name:  "Example 1: [3, 9, 20, null, null, 15, 7]",
			input: []*int{intPtr(3), intPtr(9), intPtr(20), nil, nil, intPtr(15), intPtr(7)},
			want:  [][]int{{3}, {9, 20}, {15, 7}},
		},
		{
			name:  "Example 2: single node [1]",
			input: []*int{intPtr(1)},
			want:  [][]int{{1}},
		},
		{
			name:  "Example 3: empty tree []",
			input: []*int{},
			want:  [][]int{},
		},
		{
			name:  "Left-skewed tree",
			input: []*int{intPtr(1), intPtr(2), nil, intPtr(3)},
			want:  [][]int{{1}, {2}, {3}},
		},
		{
			name:  "Complete binary tree",
			input: []*int{intPtr(1), intPtr(2), intPtr(3), intPtr(4), intPtr(5), intPtr(6), intPtr(7)},
			want:  [][]int{{1}, {2, 3}, {4, 5, 6, 7}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeFromLevelOrder(tt.input)
			got := levelOrder(root)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("levelOrder(%v) = %v; want %v", tt.input, got, tt.want)
			}
		})
	}
}
