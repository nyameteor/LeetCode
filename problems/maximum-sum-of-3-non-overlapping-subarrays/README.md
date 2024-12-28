# 689. Maximum Sum of 3 Non-Overlapping Subarrays

- Difficulty: Hard
- Topics: Array, Dynamic Programming
- Link: https://leetcode.com/problems/maximum-sum-of-3-non-overlapping-subarrays/

## Description

Given an integer array `nums` and an integer `k`, find three non-overlapping subarrays of length `k` with maximum sum and return them.

Return the result as a list of indices representing the starting position of each interval (**0-indexed**). If there are multiple answers, return the lexicographically smallest one.

**Example 1:**

```
**Input:** nums = [1,2,1,2,6,7,5,1], k = 2
**Output:** [0,3,5]
**Explanation:** Subarrays [1, 2], [2, 6], [7, 5] correspond to the starting indices [0, 3, 5].
We could have also taken [2, 1], but an answer of [1, 3, 5] would be lexicographically larger.

```

**Example 2:**

```
**Input:** nums = [1,2,1,2,1,2,1,2,1], k = 2
**Output:** [0,2,4]

```

**Constraints:**

- `1 <= nums.length <= 2 * 10^4`
- `1 <= nums[i] < 2^16`
- `1 <= k <= floor(nums.length / 3)`

## Solution

This problem is a more complex variant of the [Maximum Subarray Problem](https://en.wikipedia.org/wiki/Maximum_subarray_problem).
Unlike the classical version, which focuses on finding a single subarray with the maximum sum, this problem introduces additional constraints:

- Selecting multiple non-overlapping subarrays.
- Ensuring the result is lexicographically smallest.

These complexities require a combination of sliding window techniques and dynamic programming for an efficient solution.

### Intuition

1. Precompute Subarray Sums:

   Use a sliding window to calculate the sum of every subarray of size kk and store them.

2. Left-to-Right Pass:

   Keep track of the best subarray sum ending at or before each index (using a Kadane-like approach).

3. Right-to-Left Pass:

   Similarly, track the best subarray sum starting at or after each index.

4. Middle Subarray:

   For each possible middle subarray, combine it with the best subarray sums on the left and right (from the precomputed values) to calculate the total sum.

5. Track Indices:

   Maintain the indices of the best subarrays to construct the result.
