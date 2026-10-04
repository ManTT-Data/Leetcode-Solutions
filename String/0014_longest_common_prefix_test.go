package string_problems

import (
	"testing"
)

func TestLongestCommonPrefix(t *testing.T) {
	tests := []struct {
		name string
		strs []string
		want string
	}{
		{
			name: "Example 1",
			strs: []string{"flower", "flow", "flight"},
			want: "fl",
		},
		{
			name: "Example 2",
			strs: []string{"dog", "racecar", "car"},
			want: "",
		},
		{
			name: "Single string",
			strs: []string{"alone"},
			want: "alone",
		},
		{
			name: "Empty string inside array",
			strs: []string{"", "b"},
			want: "",
		},
		{
			name: "Common prefix is a full string",
			strs: []string{"ab", "a"},
			want: "a",
		},
		{
			name: "Identical strings",
			strs: []string{"test", "test", "test"},
			want: "test",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := longestCommonPrefix(tt.strs)
			if got != tt.want {
				t.Errorf("longestCommonPrefix(%v) = %q; want %q", tt.strs, got, tt.want)
			}
		})
	}
}
