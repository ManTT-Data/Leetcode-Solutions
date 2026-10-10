package twopointer

import "testing"

func TestNumSubarrayProductLessThanK(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{10, 5, 2, 6},
			k:    100,
			want: 8,
		},
		{
			name: "Example 2 - k is 0",
			nums: []int{1, 2, 3},
			k:    0,
			want: 0,
		},
		{
			name: "k is 1 (no product strictly less than 1)",
			nums: []int{1, 2, 3},
			k:    1,
			want: 0,
		},
		{
			name: "All subarrays valid",
			nums: []int{1, 2, 3},
			k:    10,
			want: 6, // [1], [2], [3], [1, 2], [2, 3], [1, 2, 3]
		},
		{
			name: "No element less than k",
			nums: []int{5, 6, 7},
			k:    5,
			want: 0,
		},
		{
			name: "Single element less than k",
			nums: []int{3},
			k:    5,
			want: 1,
		},
		{
			name: "Single element equal to k",
			nums: []int{5},
			k:    5,
			want: 0,
		},
		{
			name: "Array with duplicate ones",
			nums: []int{1, 1, 1},
			k:    2,
			want: 6, // product is always 1 < 2
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numSubarrayProductLessThanK(tt.nums, tt.k)
			if got != tt.want {
				t.Errorf("numSubarrayProductLessThanK(%v, %d) = %d; want %d", tt.nums, tt.k, got, tt.want)
			}
		})
	}
}
