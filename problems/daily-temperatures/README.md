# 739. Daily Temperatures

- Difficulty: Medium
- Topics: Array, Stack, Monotonic Stack
- Link: https://leetcode.com/problems/daily-temperatures/

## Description

Given an array of integers `temperatures` represents the daily temperatures, return _an array_ `answer` _such that_ `answer[i]` _is the number of days you have to wait after the_ `ith` _day to get a warmer temperature_. If there is no future day for which this is possible, keep `answer[i] == 0` instead.

**Example 1:**

```
Input: temperatures = [73,74,75,71,69,72,76,73]
Output: [1,1,4,2,1,1,0,0]
```

**Example 2:**

```
Input: temperatures = [30,40,50,60]
Output: [1,1,1,0]
```

**Example 3:**

```
Input: temperatures = [30,60,90]
Output: [1,1,0]
```

**Constraints:**

- `1 <= temperatures.length <= 10^5`
- `30 <= temperatures[i] <= 100`

## Solution

### Key Idea

Use a **monotonic decreasing stack** (store indices) to efficiently find the **next warmer day** for each temperature.

- Traverse the array **from right to left**.
- For each day `i`, pop from the stack until you find a **warmer future day**.
- If the stack is not empty, the difference in indices gives the wait time.
- Push `i` onto the stack to be a future candidate.

This avoids nested loops and ensures **O(n)** time complexity.
