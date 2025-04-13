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

### Key Idea

Similar to [15. 3Sum](../3sum/README.md), we can also use two pointers approach:

- Sort the array and fix one number at a time.
- Use two pointers to find the other two numbers whose sum is closest to the target.
- Track the minimum distance between current sum and target.

This reduces the time complexity from **O(n^3)** to **O(n^2)**.
