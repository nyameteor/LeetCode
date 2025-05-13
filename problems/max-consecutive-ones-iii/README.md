# 1004. Max Consecutive Ones III

- Difficulty: Medium
- Topics: Array, Binary Search, Sliding Window, Prefix Sum
- Link: https://leetcode.com/problems/max-consecutive-ones-iii/

## Description

Given a binary array `nums` and an integer `k`, return _the maximum number of consecutive_ `1`_'s in the array if you can flip at most_ `k` `0`'s.

**Example 1:**

```
**Input:** nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
**Output:** 6
**Explanation:** [1,1,1,0,0,**1**,1,1,1,1,**1**]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
```

**Example 2:**

```
**Input:** nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
**Output:** 10
**Explanation:** [0,0,1,1,**1**,**1**,1,1,1,**1**,1,1,0,0,0,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
```

**Constraints:**

- `1 <= nums.length <= 105`
- `nums[i]` is either `0` or `1`.
- `0 <= k <= nums.length`

## Solution

### Approach: Sliding Window

- Maintain a window with at most `k` zeros.
- Expand the window to the right; if zeros exceed `k`, shrink from the left.
- Track the maximum valid window length.

Runs in O(n) time and handles all edge cases (including `k = 0`) cleanly.

### References

- [[Java/C++/Python] Sliding Window](https://leetcode.com/problems/max-consecutive-ones-iii/solutions/247564/java-c-python-sliding-window/)
