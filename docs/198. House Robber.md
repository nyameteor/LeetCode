# 198. House Robber

- Difficulty: Medium
- Topics: Array, Dynamic Programming
- Link: https://leetcode.com/problems/house-robber/

## Description

You are a professional robber planning to rob houses along a street. Each house has a certain amount of money stashed, the only constraint stopping you from robbing each of them is that adjacent houses have security systems connected and **it will automatically contact the police if two adjacent houses were broken into on the same night**.

Given an integer array `nums` representing the amount of money of each house, return \*the maximum amount of money you can rob tonight **without alerting the police\***.

**Example 1:**

```
Input: nums = [1,2,3,1]
Output: 4
Explanation: Rob house 1 (money = 1) and then rob house 3 (money = 3).
Total amount you can rob = 1 + 3 = 4.
```

**Example 2:**

```
Input: nums = [2,7,9,3,1]
Output: 12
Explanation: Rob house 1 (money = 2), rob house 3 (money = 9) and rob house 5 (money = 1).
Total amount you can rob = 2 + 9 + 1 = 12.
```

**Constraints:**

- `1 <= nums.length <= 100`
- `0 <= nums[i] <= 400`

## Solution

### Dynamic Programming

给定一个包含 $n$ 个元素的数组 $A$（下标自 1 开始），设偷盗 $A$ 中前 $n$ 个元素能获益的最大值为 $C(n)$。

$C(n)$ 的方程：

$$
C(n) =
    \begin{cases}
        0 & n = 0 \\
        A_1 & n = 1 \\
        max\{C(n-2) + A_n, C(n-1)\} & n \geq 2
    \end{cases}
$$

计算 $C(n)$ 的过程中存在重叠子问题，可用动态规划优化。
