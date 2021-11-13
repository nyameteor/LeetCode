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

### Monotonic Stack

单调栈适合处理典型的问题：

- Next Greater Element
- Range queries in an array

参考： [Medium/Monotonic Stack](https://medium.com/techtofreedom/algorithms-for-interview-2-monotonic-stack-462251689da8)

```shell
index
0   1   2   3   4   5   6   7
nums
73  74  75  71  69  72  76  73

monotonic stack
0
1   1       3   4
2   2   2   5   5   5
6   6   6   6   6   6   6   7

result
1   1   4   2   1   1   0   0
```
