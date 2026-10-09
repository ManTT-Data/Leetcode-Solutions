package twopointer

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name   string
		prices []int
		want   int
	}{
		{
			name:   "Example 1",
			prices: []int{7, 1, 5, 3, 6, 4},
			want:   5,
		},
		{
			name:   "Example 2 - Decreasing prices, no profit",
			prices: []int{7, 6, 4, 3, 1},
			want:   0,
		},
		{
			name:   "Single day price",
			prices: []int{5},
			want:   0,
		},
		{
			name:   "Strictly increasing prices",
			prices: []int{1, 2, 3, 4, 5},
			want:   4,
		},
		{
			name:   "All identical prices",
			prices: []int{3, 3, 3, 3},
			want:   0,
		},
		{
			name:   "Fluctuating with new minimum appearing later",
			prices: []int{2, 4, 1, 7},
			want:   6,
		},
		{
			name:   "Minimum at the very end",
			prices: []int{3, 8, 2, 5, 1},
			want:   5,
		},
		{
			name:   "Max constraint difference",
			prices: []int{0, 10000},
			want:   10000,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.prices)
			if got != tt.want {
				t.Errorf("maxProfit(%v) = %d; want %d", tt.prices, got, tt.want)
			}
		})
	}
}
