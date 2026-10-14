package twopointer

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{
			name:   "Example 1",
			height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			want:   49,
		},
		{
			name:   "Example 2: Minimum length with identical heights",
			height: []int{1, 1},
			want:   1,
		},
		{
			name:   "Two elements of different heights",
			height: []int{4, 3},
			want:   3,
		},
		{
			name:   "Strictly increasing heights",
			height: []int{1, 2, 3, 4, 5},
			want:   6, // indices 1 and 4: min(2, 5) * 3 = 6
		},
		{
			name:   "Strictly decreasing heights",
			height: []int{5, 4, 3, 2, 1},
			want:   6,
		},
		{
			name:   "Taller bars at the boundaries",
			height: []int{10, 1, 1, 1, 10},
			want:   40,
		},
		{
			name:   "Very tall bars close together in the center",
			height: []int{1, 2, 100, 100, 2, 1},
			want:   100, // indices 2 and 3: min(100, 100) * 1 = 100
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxArea(tt.height)
			if got != tt.want {
				t.Errorf("maxArea(%v) = %d; want %d", tt.height, got, tt.want)
			}
		})
	}
}
