package string_problems

// Problem: 14. Longest Common Prefix
// Difficulty: Easy
// Link: https://leetcode.com/problems/longest-common-prefix/
//
// Time Complexity: O(S) where S is the sum of all characters in all strings.
// Space Complexity: O(1) - No extra memory allocation besides the result slice.

// longestCommonPrefix finds the longest common prefix string amongst an array of strings.
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}

	// Quét theo cột dọc (Vertical Scanning) dựa trên chuỗi đầu tiên
	first := strs[0]
	for i := 0; i < len(first); i++ {
		char := first[i]
		for j := 1; j < len(strs); j++ {
			// Dừng khi chỉ số vượt độ dài chuỗi strs[j] hoặc ký tự không khớp
			if i == len(strs[j]) || strs[j][i] != char {
				return first[:i]
			}
		}
	}

	return first
}
