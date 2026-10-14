package string_problems

// Problem: 316. Remove Duplicate Letters
// Difficulty: Medium
// Link: https://leetcode.com/problems/remove-duplicate-letters/
//
// Time Complexity: O(n) - Mỗi ký tự được duyệt qua và đẩy vào/lấy ra khỏi stack tối đa một lần.
// Space Complexity: O(1) - Stack và các mảng phụ chỉ lưu trữ tối đa 26 ký tự chữ cái tiếng Anh (bảng chữ cái cố định).

// removeDuplicateLetters loại bỏ các ký tự trùng lặp sao cho mỗi ký tự chỉ xuất hiện một lần
// và chuỗi kết quả có thứ tự từ điển nhỏ nhất có thể.
//
// Sử dụng kỹ thuật Monotonic Stack (Ngăn xếp đơn điệu) kết hợp Tham lam (Greedy):
// 1. Ghi nhận vị trí xuất hiện cuối cùng (lastIndex) của mỗi ký tự trong chuỗi s.
// 2. Sử dụng mảng boolean `inStack` để theo dõi xem ký tự đã có trong kết quả hay chưa.
// 3. Duyệt qua từng ký tự c tại vị trí i:
//    - Nếu c đã có trong stack (`inStack[c] == true`), bỏ qua vì giữ nguyên vị trí trước đó
//      sẽ mang lại thứ tự từ điển tốt hơn.
//    - Nếu c chưa có trong stack:
//      + Trong khi stack không rỗng, đỉnh stack `top > c` (vi phạm thứ tự từ điển)
//        VÀ ký tự `top` vẫn còn xuất hiện ở phía sau chuỗi (`lastIndex[top] > i`):
//        Ta tham lam pop `top` ra ngoài (vì sau này vẫn có thể thêm lại `top`, nhường chỗ cho `c`
//        đứng trước sẽ tạo ra chuỗi có thứ tự từ điển nhỏ hơn).
//      + Đẩy c vào stack và đánh dấu `inStack[c] = true`.
func removeDuplicateLetters(s string) string {
	// Lưu chỉ số xuất hiện cuối cùng của mỗi ký tự 'a'..'z'
	var lastIndex [26]int
	for i := 0; i < len(s); i++ {
		lastIndex[s[i]-'a'] = i
	}

	var inStack [26]bool
	stack := make([]byte, 0, 26)

	for i := 0; i < len(s); i++ {
		c := s[i]
		idx := c - 'a'

		// Nếu ký tự này đã có trong stack, không thêm lại
		if inStack[idx] {
			continue
		}

		// Pop các ký tự lớn hơn c ra khỏi stack nếu ký tự đó còn xuất hiện phía sau
		for len(stack) > 0 {
			top := stack[len(stack)-1]
			topIdx := top - 'a'

			if top > c && lastIndex[topIdx] > i {
				stack = stack[:len(stack)-1]
				inStack[topIdx] = false
			} else {
				break
			}
		}

		stack = append(stack, c)
		inStack[idx] = true
	}

	return string(stack)
}
