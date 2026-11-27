package tree

import "testing"

func TestKthSmallest(t *testing.T) {
	tests := []struct {
		name  string
		input []*int
		k     int
		want  int
	}{
		{
			name:  "Example 1: [3, 1, 4, null, 2], k = 1",
			input: []*int{intPtr(3), intPtr(1), intPtr(4), nil, intPtr(2)},
			k:     1,
			want:  1,
		},
		{
			name:  "Example 2: [5, 3, 6, 2, 4, null, null, 1], k = 3",
			input: []*int{intPtr(5), intPtr(3), intPtr(6), intPtr(2), intPtr(4), nil, nil, intPtr(1)},
			k:     3,
			want:  3,
		},
		{
			name:  "Single node tree, k = 1",
			input: []*int{intPtr(42)},
			k:     1,
			want:  42,
		},
		{
			name:  "Largest element in BST",
			input: []*int{intPtr(4), intPtr(2), intPtr(6), intPtr(1), intPtr(3), intPtr(5), intPtr(7)},
			k:     7,
			want:  7,
		},
		{
			name:  "Right-skewed BST",
			input: []*int{intPtr(1), nil, intPtr(2), nil, intPtr(3), nil, intPtr(4)},
			k:     3,
			want:  3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			root := buildTreeFromLevelOrder(tt.input)
			got := kthSmallest(root, tt.k)
			if got != tt.want {
				t.Errorf("kthSmallest(%v, %d) = %d; want %d", tt.input, tt.k, got, tt.want)
			}
		})
	}
}
