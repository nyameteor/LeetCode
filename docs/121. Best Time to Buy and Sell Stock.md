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

给定一个数组，它的第 i 个元素是一支给定股票第 i 天的价格。

如果你最多只允许完成一笔交易（即买入和卖出一支股票一次），设计一个算法来计算你所能获取的最大利润。

注意：你不能在买入股票前卖出股票。

### Dynamic Programming

这是一道典型的动态规划问题。由于有两个变量（当前天数，持有/未持有股票），所以是一个二维动态规划问题。

如果使用 `tabulation` 方法，创建 table[i][j]，i 代表第 i 天，j = 0 代表未持有股票，j = 1 代表持有股票。递推关系如下：

```shell
// 今天未持有股票 = max(昨天未持有股票，今天卖了股票)
table(i, 0) = max(table(i-1, 0), table(i-1, 1) + price[i])

// 今天持有股票 = max(昨天持有股票，今天买了股票)
table(i, 1) = max(table(i-1, 1), -price[i])
```

Java 实现：

- Runtime: 24 ms, faster than 5.40% of Java online submissions for Best Time to Buy and Sell Stock.
- Memory Usage: 55.7 MB, less than 13.22% of Java online submissions for Best Time to Buy and Sell Stock.

```java
public int maxProfit(int[] prices) {
  int n = prices.length;
  if (n <= 1)
    return 0;

  int[][] table = new int[n][2];

  table[0][0] = 0;
  table[0][1] = -prices[0];

  for (int i = 1; i < n; i++) {
    table[i][0] = Math.max(table[i - 1][0], table[i - 1][1] + prices[i]);
    table[i][1] = Math.max(table[i - 1][1], -prices[i]);
  }

  return table[n - 1][0];
}
```
