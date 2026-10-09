package array

import "testing"

func TestFindMin(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{3, 4, 5, 1, 2},
			want: 1,
		},
		{
			name: "Example 2",
			nums: []int{4, 5, 6, 7, 0, 1, 2},
			want: 0,
		},
		{
			name: "Example 3",
			nums: []int{11, 13, 15, 17},
			want: 11,
		},
		{
			name: "Single element",
			nums: []int{1},
			want: 1,
		},
		{
			name: "Two elements rotated",
			nums: []int{2, 1},
			want: 1,
		},
		{
			name: "Two elements sorted",
			nums: []int{1, 2},
			want: 1,
		},
		{
			name: "Rotated by 1",
			nums: []int{2, 3, 4, 5, 1},
			want: 1,
		},
		{
			name: "Large rotated array",
			nums: []int{5, 6, 7, 8, 9, 10, 1, 2, 3, 4},
			want: 1,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := findMin(tt.nums)
			if got != tt.want {
				t.Errorf("findMin(%v) = %d; want %d", tt.nums, got, tt.want)
			}
		})
	}
}
