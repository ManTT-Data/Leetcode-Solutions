package array

import "testing"

func TestThreeSumClosest(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   int
	}{
		{
			name:   "Example 1",
			nums:   []int{-1, 2, 1, -4},
			target: 1,
			want:   2,
		},
		{
			name:   "Example 2",
			nums:   []int{0, 0, 0},
			target: 1,
			want:   0,
		},
		{
			name:   "Exact match",
			nums:   []int{1, 1, 1, 0},
			target: -100,
			want:   2,
		},
		{
			name:   "Negative target",
			nums:   []int{4, 0, 5, -5, 3, 3, 0, -4, -5},
			target: -2,
			want:   -2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSumClosest(tt.nums, tt.target)
			if got != tt.want {
				t.Errorf("threeSumClosest(%v, %d) = %d; want %d", tt.nums, tt.target, got, tt.want)
			}
		})
	}
}
