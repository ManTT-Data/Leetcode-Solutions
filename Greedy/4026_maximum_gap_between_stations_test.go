package greedy

import "testing"

func TestMaxGapBetweenStations(t *testing.T) {
	tests := []struct {
		name    string
		skill   string
		station string
		want    int
	}{
		{
			name:    "Example 1",
			skill:   "aa",
			station: "aaaa",
			want:    3,
		},
		{
			name:    "Example 2",
			skill:   "xyz",
			station: "xyzz",
			want:    2,
		},
		{
			name:    "Single worker - gap is 0",
			skill:   "a",
			station: "a",
			want:    0,
		},
		{
			name:    "Single worker with multiple options",
			skill:   "b",
			station: "abbbb",
			want:    0,
		},
		{
			name:    "Exact match length",
			skill:   "abc",
			station: "abc",
			want:    1,
		},
		{
			name:    "Max gap between endpoints",
			skill:   "ab",
			station: "azzzzb",
			want:    5, // station 0 ('a') and station 5 ('b') -> 5 - 0 = 5
		},
		{
			name:    "Alternating pattern",
			skill:   "aba",
			station: "ababa",
			want:    3, // leftMatch = [0, 1, 2], rightMatch = [2, 3, 4] -> max gap = 3
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxGapBetweenStations(tt.skill, tt.station)
			if got != tt.want {
				t.Errorf("maxGapBetweenStations(%q, %q) = %d; want %d", tt.skill, tt.station, got, tt.want)
			}
		})
	}
}
