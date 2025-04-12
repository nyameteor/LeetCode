# 18. 4Sum

- Difficulty: Medium
- Topics: Array, Two Pointers, Sorting
- Link: https://leetcode.com/problems/4sum/

## Description

Given an array `nums` of `n` integers, return _an array of all the **unique** quadruplets_ `[nums[a], nums[b], nums[c], nums[d]]` such that:

- `0 <= a, b, c, d < n`
- `a`, `b`, `c`, and `d` are **distinct**.
- `nums[a] + nums[b] + nums[c] + nums[d] == target`

You may return the answer in **any order**.

**Example 1:**

```
Input: nums = [1,0,-1,0,-2,2], target = 0
Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
```

**Example 2:**

```
Input: nums = [2,2,2,2,2], target = 8
Output: [[2,2,2,2]]
```

**Constraints:**

- `1 <= nums.length <= 200`
- `-10^9 <= nums[i] <= 10^9`
- `-10^9 <= target <= 10^9`

## Solution

### Approach

- Sort the array to simplify duplicate handling and enable two-pointer scanning.
- Use a recursive `kSum` function:
  - For `k > 2`, fix one number and recursively solve `(k-1)`-Sum.
  - For `k == 2`, apply the two-pointer technique.
- Skip duplicates at each level to avoid repeating subsets.

This approach works efficiently for small values of `k` (like 3Sum and 4Sum).
