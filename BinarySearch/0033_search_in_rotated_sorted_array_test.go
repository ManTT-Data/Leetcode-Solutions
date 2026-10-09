package binarysearch

import "testing"

func TestSearchInRotatedSortedArray(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "Example 1",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 0,
			want:   4,
		},
		{
			name:   "Example 2 - Target not found",
			nums:   []int{4, 5, 6, 7, 0, 1, 2},
			target: 3,
			want:   -1,
		},
		{
			name:   "Example 3 - Single element not found",
			nums:   []int{1},
			target: 0,
			want:   -1,
		},
		{
			name:   "Single element found",
			nums:   []int{1},
			target: 1,
			want:   0,
		},
		{
			name:   "Rotated at index 1",
			nums:   []int{3, 1},
			target: 1,
			want:   1,
		},
		{
			name:   "Target at right side",
			nums:   []int{5, 1, 3},
			target: 3,
			want:   2,
		},
		{
			name:   "Unrotated sorted array",
			nums:   []int{1, 2, 3, 4, 5},
			target: 4,
			want:   3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := search(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("search(%v, %d) = %d; want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
