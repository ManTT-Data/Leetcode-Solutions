package twopointer

// Problem: 121. Best Time to Buy and Sell Stock
// Difficulty: Easy
// Link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
//
// Time Complexity: O(n) - Duyệt qua mảng giá đúng một lần với hai con trỏ (Two Pointers / Sliding Window).
// Space Complexity: O(1) - Chỉ sử dụng bộ nhớ cố định cho các con trỏ và biến lưu kết quả.

// maxProfit tính lợi nhuận lớn nhất có thể đạt được theo hướng tiếp cận Two Pointers (Kỹ thuật hai con trỏ / Cửa sổ trượt).
// - Con trỏ `left` đại diện cho ngày mua (điểm giá thấp tiềm năng).
// - Con trỏ `right` đại diện cho ngày bán (duyệt tới tương lai).
func maxProfit(prices []int) int {
	left, right := 0, 1
	maxProfit := 0

	for right < len(prices) {
		if prices[left] < prices[right] {
			// Có lãi khi mua ở ngày left và bán ở ngày right
			profit := prices[right] - prices[left]
			if profit > maxProfit {
				maxProfit = profit
			}
		} else {
			// prices[right] <= prices[left]: tìm thấy ngày mua có giá rẻ hơn!
			// Dịch con trỏ mua sang ngày right
			left = right
		}
		right++
	}

	return maxProfit
}
