package greedy

import "testing"

func TestCanCompleteCircuit(t *testing.T) {
	tests := []struct {
		name string
		gas  []int
		cost []int
		want int
	}{
		{
			name: "Example 1",
			gas:  []int{1, 2, 3, 4, 5},
			cost: []int{3, 4, 5, 1, 2},
			want: 3,
		},
		{
			name: "Example 2 - Impossible",
			gas:  []int{2, 3, 4},
			cost: []int{3, 4, 3},
			want: -1,
		},
		{
			name: "Single station - Enough gas",
			gas:  []int{2},
			cost: []int{2},
			want: 0,
		},
		{
			name: "Single station - Not enough gas",
			gas:  []int{2},
			cost: []int{3},
			want: -1,
		},
		{
			name: "Start station at index 0",
			gas:  []int{3, 1, 1},
			cost: []int{1, 2, 2},
			want: 0,
		},
		{
			name: "Start station at the last index",
			gas:  []int{1, 1, 1, 5},
			cost: []int{2, 2, 2, 1},
			want: 3,
		},
		{
			name: "Exact total gas equals total cost",
			gas:  []int{5, 8, 2, 8},
			cost: []int{6, 5, 6, 6},
			want: 3,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := canCompleteCircuit(tt.gas, tt.cost)
			if got != tt.want {
				t.Errorf("canCompleteCircuit(%v, %v) = %d; want %d", tt.gas, tt.cost, got, tt.want)
			}
		})
	}
}
