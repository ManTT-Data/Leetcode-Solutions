package binarysearch

import "testing"

func TestSearchII(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		{
			name:   "Example 1",
			nums:   []int{2, 5, 6, 0, 0, 1, 2},
			target: 0,
			want:   true,
		},
		{
			name:   "Example 2 - Target not found",
			nums:   []int{2, 5, 6, 0, 0, 1, 2},
			target: 3,
			want:   false,
		},
		{
			name:   "Duplicates with target on left side",
			nums:   []int{1, 0, 1, 1, 1},
			target: 0,
			want:   true,
		},
		{
			name:   "Duplicates with target on right side",
			nums:   []int{1, 1, 1, 0, 1},
			target: 0,
			want:   true,
		},
		{
			name:   "Single element found",
			nums:   []int{1},
			target: 1,
			want:   true,
		},
		{
			name:   "Single element not found",
			nums:   []int{1},
			target: 0,
			want:   false,
		},
		{
			name:   "All identical elements, not target",
			nums:   []int{2, 2, 2, 2, 2},
			target: 5,
			want:   false,
		},
		{
			name:   "Target at pivot position",
			nums:   []int{3, 1, 2, 3, 3, 3, 3},
			target: 1,
			want:   true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := searchII(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("searchII(%v, %d) = %v; want %v", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
