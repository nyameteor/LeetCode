# 121. Best Time to Buy and Sell Stock

- Difficulty: Easy
- Topics: Array, Dynamic Programming
- Link: https://leetcode.com/problems/best-time-to-buy-and-sell-stock/

## Description

You are given an array `prices` where `prices[i]` is the price of a given stock on the `ith` day.

You want to maximize your profit by choosing a **single day** to buy one stock and choosing a **different day in the future** to sell that stock.

Return _the maximum profit you can achieve from this transaction_. If you cannot achieve any profit, return `0`.

**Example 1:**

```
Input: prices = [7,1,5,3,6,4]
Output: 5
Explanation: Buy on day 2 (price = 1) and sell on day 5 (price = 6), profit = 6-1 = 5.
Note that buying on day 2 and selling on day 1 is not allowed because you must buy before you sell.
```

**Example 2:**

```
Input: prices = [7,6,4,3,1]
Output: 0
Explanation: In this case, no transactions are done and the max profit = 0.
```

**Constraints:**

- `1 <= prices.length <= 10^5`
- `0 <= prices[i] <= 10^4`

## Solution

The problem requires finding the maximum profit from a single buy-sell transaction. The key observation is that **we need to buy at a minimum price seen so far and sell at a later, higher price** to maximize profit.

### **Approach: One-Pass with Min Tracking**

- Track the minimum price so far (`minPrice`).
- Update the maximum profit (`bestProfit`) by computing the difference between the current price and `minPrice`.
- Iterate through the array once, updating both values dynamically.
