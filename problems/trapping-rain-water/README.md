# 42. Trapping Rain Water

- Difficulty: Hard
- Topics: Array, Two Pointers, Dynamic Programming, Stack, Monotonic Stack
- Link: https://leetcode.com/problems/trapping-rain-water/

## Description

Given `n` non-negative integers representing an elevation map where the width of each bar is `1`, compute how much water it can trap after raining.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2018/10/22/rainwatertrap.png)

```
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
```

**Example 2:**

```
Input: height = [4,2,0,3,2,5]
Output: 9
```

**Constraints:**

- `n == height.length`
- `1 <= n <= 2 * 10^4`
- `0 <= height[i] <= 10^5`

## Solution

### Key Idea

To calculate water trapped at each index `i`, we need to find the tallest bar on the left and right sides:

```
water[i] = min(maxLeft[i], maxRight[i]) - height[i]
```

### Approach: Bottom-Up DP

Avoid repeated work by precomputing:

- `maxLeft[i] = max(maxLeft[i-1], height[i])`
- `maxRight[i] = max(maxRight[i+1], height[i])`

Then for each index `i`, accumulate the trapped water.

Complexity:

- Time: O(n)
- Space: O(n)

### Approach: Two Pointers

- Use two pointers `l` and `r` at both ends, tracking `maxLeft` and `maxRight`.
- Always move the side with the smaller height:
  - If `height[l] <= height[r]`, water is limited by `maxLeft`, move `l++`.
  - Else, water is limited by `maxRight`, move `r--`.
- Accumulate trapped water at each step.

Complexity:

- Time: O(n)
- Space: O(1)
