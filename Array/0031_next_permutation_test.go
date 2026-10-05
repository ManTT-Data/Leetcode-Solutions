package array

import (
	"reflect"
	"testing"
)

func TestNextPermutation(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want []int
	}{
		{
			name: "Example 1",
			nums: []int{1, 2, 3},
			want: []int{1, 3, 2},
		},
		{
			name: "Example 2 - Descending order",
			nums: []int{3, 2, 1},
			want: []int{1, 2, 3},
		},
		{
			name: "Example 3 - With duplicates",
			nums: []int{1, 1, 5},
			want: []int{1, 5, 1},
		},
		{
			name: "Single element",
			nums: []int{1},
			want: []int{1},
		},
		{
			name: "Complex case",
			nums: []int{1, 3, 5, 4, 2},
			want: []int{1, 4, 2, 3, 5},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			input := make([]int, len(tt.nums))
			copy(input, tt.nums)
			nextPermutation(input)
			if !reflect.DeepEqual(input, tt.want) {
				t.Errorf("nextPermutation(%v) resulted in %v; want %v", tt.nums, input, tt.want)
			}
		})
	}
}
