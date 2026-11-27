package greedy

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "Example 1: reachable [2, 3, 1, 1, 4]",
			nums: []int{2, 3, 1, 1, 4},
			want: true,
		},
		{
			name: "Example 2: stuck at zero [3, 2, 1, 0, 4]",
			nums: []int{3, 2, 1, 0, 4},
			want: false,
		},
		{
			name: "Single element [0]",
			nums: []int{0},
			want: true,
		},
		{
			name: "Starts with zero and length > 1",
			nums: []int{0, 2, 3},
			want: false,
		},
		{
			name: "Jump right over zero [2, 0, 0]",
			nums: []int{2, 0, 0},
			want: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canJump(tt.nums)
			if got != tt.want {
				t.Errorf("canJump(%v) = %v; want %v", tt.nums, got, tt.want)
			}
		})
	}
}
