package greedy

import "testing"

func TestLargestNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want string
	}{
		{
			name: "Example 1",
			nums: []int{10, 2},
			want: "210",
		},
		{
			name: "Example 2",
			nums: []int{3, 30, 34, 5, 9},
			want: "9534330",
		},
		{
			name: "All zeros",
			nums: []int{0, 0},
			want: "0",
		},
		{
			name: "Single zero",
			nums: []int{0},
			want: "0",
		},
		{
			name: "Single non-zero element",
			nums: []int{1},
			want: "1",
		},
		{
			name: "Same prefix with different lengths",
			nums: []int{111311, 1113},
			want: "1113111311",
		},
		{
			name: "Numbers with common digits",
			nums: []int{8308, 830},
			want: "8308830",
		},
		{
			name: "Large number with zeros",
			nums: []int{1000000000, 1000},
			want: "10001000000000",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := largestNumber(tt.nums)
			if got != tt.want {
				t.Errorf("largestNumber(%v) = %q; want %q", tt.nums, got, tt.want)
			}
		})
	}
}
