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

Similar to [15. 3Sum](../3sum/README.md), we can also use two pointers approach. Sort the array, and get target `sum` by iterating with left and right pointers.
The difference is that the solution can contain the same number, such as `-1, -1, 1`.
