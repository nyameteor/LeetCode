# 169. Majority Element

- Difficulty: Easy
- Topics: Array, Hash Table, Divide and Conquer, Sorting, Counting
- Link: https://leetcode.com/problems/majority-element/

## Description

Given an array `nums` of size `n`, return _the majority element_.

The majority element is the element that appears more than `⌊n / 2⌋` times. You may assume that the majority element always exists in the array.

**Example 1:**

```
Input: nums = [3,2,3]
Output: 3
```

**Example 2:**

```
Input: nums = [2,2,1,1,1,2,2]
Output: 2
```

**Constraints:**

- `n == nums.length`
- `1 <= n <= 5 * 10^4`
- `-2^31 <= nums[i] <= 2^31 - 1`

**Follow-up:** Could you solve the problem in linear time and in `O(1)` space?

## Solution

Use the [Boyer–Moore majority vote algorithm](https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_majority_vote_algorithm#Description) with these steps:

- Initialize an element `m` and a counter `c` with `c = 0`
- For each element `x` of the input sequence:
  - If `c = 0`, then assign `m = x` and `c = 1`
  - else if `m = x`, then assign `c = c + 1`
  - else assign `c = c − 1`
- Return `m`
