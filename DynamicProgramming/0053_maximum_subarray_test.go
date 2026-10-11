package dynamicprogramming

import "testing"

func TestMaxSubArray(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{-2, 1, -3, 4, -1, 2, 1, -5, 4},
			want: 6, // subarray [4, -1, 2, 1]
		},
		{
			name: "Example 2: Single element positive",
			nums: []int{1},
			want: 1,
		},
		{
			name: "Example 3: All positive with small negative",
			nums: []int{5, 4, -1, 7, 8},
			want: 23,
		},
		{
			name: "All negative numbers",
			nums: []int{-5, -2, -8, -1},
			want: -1, // largest is -1
		},
		{
			name: "Single negative element",
			nums: []int{-3},
			want: -3,
		},
		{
			name: "All identical positive numbers",
			nums: []int{2, 2, 2, 2},
			want: 8,
		},
		{
			name: "Two elements, larger is first",
			nums: []int{2, -1},
			want: 2,
		},
		{
			name: "Two elements, larger is second",
			nums: []int{-1, 2},
			want: 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name+"_Kadane", func(t *testing.T) {
			got := maxSubArray(tt.nums)
			if got != tt.want {
				t.Errorf("maxSubArray(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})

		t.Run(tt.name+"_DivideAndConquer", func(t *testing.T) {
			got := maxSubArrayDivideAndConquer(tt.nums)
			if got != tt.want {
				t.Errorf("maxSubArrayDivideAndConquer(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}
