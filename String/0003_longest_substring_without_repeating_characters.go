package string_problems

// Problem: 3. Longest Substring Without Repeating Characters
// Difficulty: Medium
// Link: https://leetcode.com/problems/longest-substring-without-repeating-characters/
//
// Time Complexity: O(n) - Con trỏ phải duyệt qua chuỗi đúng 1 lần, con trỏ trái chỉ tiến lên.
// Space Complexity: O(1) - Dùng mảng cố định 128 phần tử cho bảng mã ASCII (hoặc O(min(n, m)) với m là kích thước bảng chữ cái).

// lengthOfLongestSubstring tìm độ dài chuỗi con liên tiếp dài nhất không chứa ký tự lặp lại.
//
// Ý tưởng giải thuật (Sliding Window):
// 1. Duy trì cửa sổ trượt `[left, right]` chứa các ký tự phân biệt.
// 2. Dùng mảng `lastSeen[128]` (khởi tạo -1) lưu vị trí xuất hiện gần nhất của mỗi ký tự ASCII.
// 3. Khi duyệt `right` từ 0 đến `len(s)-1`:
//    - Nếu ký tự `ch = s[right]` đã xuất hiện bên trong cửa sổ hiện tại (`lastSeen[ch] >= left`),
//      ta nhảy `left` thẳng đến `lastSeen[ch] + 1` để loại bỏ ký tự trùng.
//    - Cập nhật `lastSeen[ch] = right` và cập nhật `maxLen`.
func lengthOfLongestSubstring(s string) int {
	var lastSeen [128]int
	for i := range lastSeen {
		lastSeen[i] = -1
	}

	maxLen := 0
	left := 0

	for right := 0; right < len(s); right++ {
		ch := s[right]
		if lastSeen[ch] >= left {
			left = lastSeen[ch] + 1
		}
		lastSeen[ch] = right

		if currLen := right - left + 1; currLen > maxLen {
			maxLen = currLen
		}
	}

	return maxLen
}
