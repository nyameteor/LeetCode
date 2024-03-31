# 992. Subarrays with K Different Integers

- Difficulty: Hard
- Topics: Array, Hash Table, Sliding Window, Counting
- Link: https://leetcode.com/problems/subarrays-with-k-different-integers/

## Description

Given an integer array `nums` and an integer `k`, return _the number of **good subarrays** of_ `nums`.

A **good array** is an array where the number of different integers in that array is exactly `k`.

- For example, `[1,2,3,1,2]` has `3` different integers: `1`, `2`, and `3`.

A **subarray** is a **contiguous** part of an array.

**Example 1:**

```
Input: nums = [1,2,1,2,3], k = 2
Output: 7
Explanation: Subarrays formed with exactly 2 different integers: [1,2], [2,1], [1,2], [2,3], [1,2,1], [2,1,2], [1,2,1,2]
```

**Example 2:**

```
Input: nums = [1,2,1,3,4], k = 3
Output: 3
Explanation: Subarrays formed with exactly 3 different integers: [1,2,1,3], [2,1,3], [1,3,4].
```

**Constraints:**

- `1 <= nums.length <= 2 * 10^4`
- `1 <= nums[i], k <= nums.length`

## Solution

To count the subarrays whose distinct integers <= `k`, we can use the sliding window pattern. For each iteration:

- Expand the right side by adding one element, and shrink the left side until the distinct integers <= `k`.
- Denote `i` as left index, `j` as right index, `res` as result. Then `res += j + 1 - i`.

If we call the above procedure as `slidingWindowAtMost(nums, k)`, then to count the subarrays whose distinct integers == `k`, we can simply return `slidingWindowAtMost(nums, k) - slidingWindowAtMost(nums, k-1)`.

### References

- https://leetcode.com/problems/subarrays-with-k-different-integers/editorial/
