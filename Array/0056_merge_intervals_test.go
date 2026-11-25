package array

import (
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {
	tests := []struct {
		name      string
		intervals [][]int
		want      [][]int
	}{
		{
			name:      "Example 1: overlapping intervals",
			intervals: [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			want:      [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		{
			name:      "Example 2: touching intervals",
			intervals: [][]int{{1, 4}, {4, 5}},
			want:      [][]int{{1, 5}},
		},
		{
			name:      "Unsorted intervals with nested interval",
			intervals: [][]int{{1, 4}, {2, 3}, {0, 2}, {6, 8}},
			want:      [][]int{{0, 4}, {6, 8}},
		},
		{
			name:      "Single interval",
			intervals: [][]int{{5, 7}},
			want:      [][]int{{5, 7}},
		},
		{
			name:      "All intervals merge into one",
			intervals: [][]int{{1, 10}, {2, 3}, {4, 8}, {9, 12}},
			want:      [][]int{{1, 12}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := merge(tt.intervals)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("merge(%v) = %v; want %v", tt.intervals, got, tt.want)
			}
		})
	}
}
