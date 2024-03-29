# 16. 3Sum Closest

- Difficulty: Medium
- Topics: Array, Two Pointers, Sorting
- Link: https://leetcode.com/problems/3sum-closest/

## Description

Given an integer array `nums` of length `n` and an integer `target`, find three integers in `nums` such that the sum is closest to `target`.

Return _the sum of the three integers_.

You may assume that each input would have exactly one solution.

**Example 1:**

```
Input: nums = [-1,2,1,-4], target = 1
Output: 2
Explanation: The sum that is closest to the target is 2. (-1 + 2 + 1 = 2).
```

**Example 2:**

```
Input: nums = [0,0,0], target = 1
Output: 0
```

**Constraints:**

- `3 <= nums.length <= 1000`
- `-1000 <= nums[i] <= 1000`
- `-10^4 <= target <= 10^4`

## Solution

也是双指针算法，对数组排序，通过左右指针遍历更新得到符合条件的 `sum`。

和 [15. 3Sum](15.%203Sum.md) 不同的地方在于其解可以包含相同的数字，如 `-1, -1, 1` 这样的组合，因此不能通过忽略相同的数字来加速。

另外，该题给变量设置初始的大数时避免使用 `INT_MAX`，因为该题中 `target` 可以为负数，`INT_MAX - target` 容易越界产生 Runtime Error。更好的做法还是根据输入数据的范围设置初始值。
