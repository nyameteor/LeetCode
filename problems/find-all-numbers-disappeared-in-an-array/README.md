# 448. Find All Numbers Disappeared in an Array

- Difficulty: Easy
- Topics: Array, Hash Table
- Link: https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

## Description

Given an array `nums` of `n` integers where `nums[i]` is in the range `[1, n]`, return *an array of all the integers in the range* `[1, n]` *that do not appear in* `nums`.

**Example 1:**

```
Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]
```

**Example 2:**

```
Input: nums = [1,1]
Output: [2]
```

**Constraints:**

- `n == nums.length`
- `1 <= n <= 10^5`
- `1 <= nums[i] <= n`

**Follow up:** Could you do it without extra space and in `O(n)` runtime? You may assume the returned list does not count as extra space.

## Solution

### Approach: In-Place Marking

- For each `x` in `nums`, mark the value at index `abs(x) - 1` as negative.
- After marking, any index with a positive value means `index + 1` is missing.
