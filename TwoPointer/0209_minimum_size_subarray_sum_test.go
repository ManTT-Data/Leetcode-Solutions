package twopointer

import "testing"

func TestMinSubArrayLen(t *testing.T) {
	tests := []struct {
		name   string
		target int
		nums   []int
		want   int
	}{
		{
			name:   "Example 1",
			target: 7,
			nums:   []int{2, 3, 1, 2, 4, 3},
			want:   2,
		},
		{
			name:   "Example 2 - Single element satisfies target",
			target: 4,
			nums:   []int{1, 4, 4},
			want:   1,
		},
		{
			name:   "Example 3 - No subarray reaches target",
			target: 11,
			nums:   []int{1, 1, 1, 1, 1, 1, 1, 1},
			want:   0,
		},
		{
			name:   "Single element equals target",
			target: 5,
			nums:   []int{5},
			want:   1,
		},
		{
			name:   "Single element less than target",
			target: 5,
			nums:   []int{3},
			want:   0,
		},
		{
			name:   "All elements needed",
			target: 15,
			nums:   []int{1, 2, 3, 4, 5},
			want:   5,
		},
		{
			name:   "Single large element in the middle",
			target: 6,
			nums:   []int{1, 2, 6, 1},
			want:   1,
		},
		{
			name:   "Multiple valid windows of different sizes",
			target: 15,
			nums:   []int{5, 1, 3, 5, 10, 7, 4, 9, 2, 8},
			want:   2, // [10, 7] or [7, 9] or [9, 8] sum >= 15 with length 2
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+"_SlidingWindow", func(t *testing.T) {
			got := minSubArrayLen(tt.target, tt.nums)
			if got != tt.want {
				t.Errorf("minSubArrayLen(%d, %v) = %d; want %d", tt.target, tt.nums, got, tt.want)
			}
		})

		t.Run(tt.name+"_BinarySearch", func(t *testing.T) {
			got := minSubArrayLenBinarySearch(tt.target, tt.nums)
			if got != tt.want {
				t.Errorf("minSubArrayLenBinarySearch(%d, %v) = %d; want %d", tt.target, tt.nums, got, tt.want)
			}
		})
	}
}
