package twopointer

import "testing"

func TestMinOperations(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		x    int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 1, 4, 2, 3},
			x:    5,
			want: 2, // remove [2, 3] from the right
		},
		{
			name: "Example 2 - Impossible to reduce to zero",
			nums: []int{5, 6, 7, 8, 9},
			x:    4,
			want: -1,
		},
		{
			name: "Example 3 - Remove from both ends",
			nums: []int{3, 2, 20, 1, 1, 3},
			x:    10,
			want: 5, // remove [3, 2] from left and [1, 1, 3] from right
		},
		{
			name: "Total sum exactly equals x",
			nums: []int{1, 2, 3},
			x:    6,
			want: 3,
		},
		{
			name: "Total sum less than x",
			nums: []int{1, 2, 3},
			x:    10,
			want: -1,
		},
		{
			name: "Single element equals x",
			nums: []int{5},
			x:    5,
			want: 1,
		},
		{
			name: "Single element less than x",
			nums: []int{5},
			x:    7,
			want: -1,
		},
		{
			name: "All elements from left side",
			nums: []int{2, 3, 4, 5, 6},
			x:    5,
			want: 2, // [2, 3] from left
		},
		{
			name: "All elements from right side",
			nums: []int{6, 5, 4, 3, 2},
			x:    5,
			want: 2, // [3, 2] from right
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := minOperations(tt.nums, tt.x)
			if got != tt.want {
				t.Errorf("minOperations(%v, %d) = %d; want %d", tt.nums, tt.x, got, tt.want)
			}
		})
	}
}
