package twopointer

import (
	"reflect"
	"testing"
)

func TestSortedSquares(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Example 1 - Mixed positive and negative",
			nums: []int{-4, -1, 0, 3, 10},
			want: []int{0, 1, 9, 16, 100},
		},
		{
			name: "Example 2 - Mixed with duplicates in squares",
			nums: []int{-7, -3, 2, 3, 11},
			want: []int{4, 9, 9, 49, 121},
		},
		{
			name: "All negative numbers",
			nums: []int{-5, -3, -2, -1},
			want: []int{1, 4, 9, 25},
		},
		{
			name: "All non-negative numbers",
			nums: []int{1, 2, 3, 4, 5},
			want: []int{1, 4, 9, 16, 25},
		},
		{
			name: "Single negative element",
			nums: []int{-5},
			want: []int{25},
		},
		{
			name: "Single zero element",
			nums: []int{0},
			want: []int{0},
		},
		{
			name: "All identical absolute values",
			nums: []int{-2, -2, 2, 2},
			want: []int{4, 4, 4, 4},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := sortedSquares(tt.nums)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortedSquares(%v) = %v; want %v", tt.nums, got, tt.want)
			}
		})
	}
}
