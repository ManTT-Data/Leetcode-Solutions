package dynamicprogramming

import "testing"

func TestRob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1: [1, 2, 3, 1]",
			nums: []int{1, 2, 3, 1},
			want: 4,
		},
		{
			name: "Example 2: [2, 7, 9, 3, 1]",
			nums: []int{2, 7, 9, 3, 1},
			want: 12,
		},
		{
			name: "Skip two houses in the middle: [2, 1, 1, 2]",
			nums: []int{2, 1, 1, 2},
			want: 4,
		},
		{
			name: "Single house",
			nums: []int{5},
			want: 5,
		},
		{
			name: "Two houses",
			nums: []int{4, 9},
			want: 9,
		},
		{
			name: "Empty slice",
			nums: []int{},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := rob(tt.nums)
			if got != tt.want {
				t.Errorf("rob(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}
