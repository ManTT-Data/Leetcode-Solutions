package twopointer

import "testing"

func TestFindDuplicate(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{1, 3, 4, 2, 2},
			want: 2,
		},
		{
			name: "Example 2",
			nums: []int{3, 1, 3, 4, 2},
			want: 3,
		},
		{
			name: "Example 3 - All elements identical",
			nums: []int{3, 3, 3, 3, 3},
			want: 3,
		},
		{
			name: "Minimal case n=1",
			nums: []int{1, 1},
			want: 1,
		},
		{
			name: "Duplicate at beginning",
			nums: []int{2, 2, 1},
			want: 2,
		},
		{
			name: "Duplicate at ends",
			nums: []int{2, 1, 3, 2},
			want: 2,
		},
		{
			name: "More than two duplicates",
			nums: []int{2, 5, 9, 6, 9, 3, 8, 9, 7, 1},
			want: 9,
		},
		{
			name: "Consecutive duplicates at end",
			nums: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 9},
			want: 9,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findDuplicate(tt.nums)
			if got != tt.want {
				t.Errorf("findDuplicate(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}
