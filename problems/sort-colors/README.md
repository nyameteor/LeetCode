# 75. Sort Colors

- Difficulty: Medium
- Topics: Array, Two Pointers, Sorting
- Link: https://leetcode.com/problems/sort-colors/

## Description

Given an array `nums` with `n` objects colored red, white, or blue, sort them **[in-place](https://en.wikipedia.org/wiki/In-place_algorithm)** so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We will use the integers `0`, `1`, and `2` to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.

**Example 1:**

```
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
```

**Example 2:**

```
Input: nums = [2,0,1]
Output: [0,1,2]
```

**Constraints:**

- `n == nums.length`
- `1 <= n <= 300`
- `nums[i]` is either `0`, `1`, or `2`.

**Follow up:** Could you come up with a one-pass algorithm using only constant extra space?

## Solution

We can use the [Dutch National Flag Algorithm](https://en.wikipedia.org/wiki/Dutch_national_flag_problem):

- Use three pointers:
  - `i`: next position to place 0 (red)
  - `j`: current element under examination
  - `k`: next position to place 2 (blue)
- Traverse the array once (`j <= k`):
  - If `nums[j] == 0`: swap with `i`, increment both `i` and `j`
  - If `nums[j] == 2`: swap with `k`, decrement `k` only (recheck `j`)
  - If `nums[j] == 1`: just increment `j`
- Ensures all 0s are on the left, 2s on the right, and 1s in the middle.
