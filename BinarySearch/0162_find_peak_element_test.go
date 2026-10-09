package binarysearch

import "testing"

func isPeak(nums []int, idx int) bool {
	n := len(nums)
	if idx < 0 || idx >= n {
		return false
	}
	leftOK := idx == 0 || nums[idx] > nums[idx-1]
	rightOK := idx == n-1 || nums[idx] > nums[idx+1]
	return leftOK && rightOK
}

func TestFindPeakElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3, 1},
		},
		{
			name: "Example 2 - Multiple peaks",
			nums: []int{1, 2, 1, 3, 5, 6, 4},
		},
		{
			name: "Single element",
			nums: []int{1},
		},
		{
			name: "Two elements strictly increasing",
			nums: []int{1, 2},
		},
		{
			name: "Two elements strictly decreasing",
			nums: []int{2, 1},
		},
		{
			name: "Strictly increasing array",
			nums: []int{1, 2, 3, 4, 5},
		},
		{
			name: "Strictly decreasing array",
			nums: []int{5, 4, 3, 2, 1},
		},
		{
			name: "Alternating elements",
			nums: []int{1, 3, 2, 4, 3, 5, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findPeakElement(tt.nums)
			if !isPeak(tt.nums, got) {
				t.Errorf("findPeakElement(%v) returned index %d (value %d), which is not a valid peak", tt.nums, got, tt.nums[got])
			}
		})
	}
}
