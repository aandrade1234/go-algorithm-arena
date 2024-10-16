# 121. Best Time to Buy and Sell Stock

You are given an array `prices` where `prices[i]` is the price of a given stock on the `ith` day.<br><br>
You want to maximize your profit by choosing a **single day** to buy one stock and choosing a **different day in the
future** to sell that stock.<br><br
Return the _maximum profit you can achieve from this transaction_. If you cannot achieve any profit, return `0`.

**Example 1:**

> **Input:** prices = [7,1,5,3,6,4]<br>
> **Output:** 5<br>
> **Explanation:** Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.<br>
> Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.

**Example 2:**
> **Input:** prices = [7,6,4,3,1]<br>
> **Output:** 0<br>
> **Explanation:** In this case, no transactions are done and the max profit = 0.

**Constraints:**
- `1 <= prices.length <= 10^5`
- `0 <= prices[i] <= 10^4`

Source: [Leetcode](https://leetcode.com/problems/best-time-to-buy-and-sell-stock/description/)
