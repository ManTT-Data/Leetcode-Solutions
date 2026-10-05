package array

// Problem: 121. Best Time to Buy and Sell Stock
// Difficulty: Easy
// Link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/
//
// Time Complexity: O(n) - Duyệt qua mảng giá đúng một lần duy nhất.
// Space Complexity: O(1) - Chỉ sử dụng hai biến phụ minPrice và maxProfit.

// maxProfit tính lợi nhuận lớn nhất có thể đạt được bằng cách mua ở một ngày và bán ở một ngày tương lai.
func maxProfit(prices []int) int {
	if len(prices) == 0 {
		return 0
	}

	minPrice := prices[0]
	maxProfit := 0

	for _, price := range prices[1:] {
		if price < minPrice {
			// Cập nhật giá mua thấp nhất tìm thấy cho đến thời điểm hiện tại
			minPrice = price
		} else if profit := price - minPrice; profit > maxProfit {
			// Bán tại ngày hôm nay nếu lợi nhuận cao hơn mức cao nhất từng đạt
			maxProfit = profit
		}
	}

	return maxProfit
}
