# 15. 3Sum

- Difficulty: Medium
- Topics: Array, Two Pointers, Sorting
- Link: https://leetcode.com/problems/3sum/

## Description

Given an integer array nums, return all the triplets `[nums[i], nums[j], nums[k]]` such that `i != j`, `i != k`, and `j != k`, and `nums[i] + nums[j] + nums[k] == 0`.

Notice that the solution set must not contain duplicate triplets.

**Example 1:**

```
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
```

**Example 2:**

```
Input: nums = []
Output: []
```

**Example 3:**

```
Input: nums = [0]
Output: []
```

**Constraints:**

- `0 <= nums.length <= 3000`
- `-10^5 <= nums[i] <= 10^5`

## Solution

### Approach

- Sort the array to simplify duplicate handling and enable two-pointer scanning.
- For each `nums[i]`, fix it as the first number of the triplet:
  - Skip duplicates for `i`.
  - Use two pointers `l` and `r` to find pairs such that `nums[i] + nums[l] + nums[r] == 0`.
  - After finding a valid triplet, move both pointers and skip duplicates.
