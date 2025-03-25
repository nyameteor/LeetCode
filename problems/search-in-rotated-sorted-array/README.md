# 33. Search in Rotated Sorted Array

- Difficulty: Medium
- Topics: Array, Binary Search
- Link: https://leetcode.com/problems/search-in-rotated-sorted-array/

## Description

There is an integer array `nums` sorted in ascending order (with **distinct** values).

Prior to being passed to your function, `nums` is **rotated** at an unknown pivot index `k` (`0 <= k < nums.length`) such that the resulting array is `[nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]]` (**0-indexed**). For example, `[0,1,2,4,5,6,7]` might be rotated at pivot index `3` and become `[4,5,6,7,0,1,2]`.

Given the array `nums` **after** the rotation and an integer `target`, return _the index of_ `target` _if it is in_ `nums`_, or_ `-1` _if it is not in_ `nums`.

You must write an algorithm with `O(log n)` runtime complexity.

**Example 1:**

```
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4
```

**Example 2:**

```
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1
```

**Example 3:**

```
Input: nums = [1], target = 0
Output: -1
```

**Constraints:**

- `1 <= nums.length <= 5000`
- `-10^4 <= nums[i] <= 10^4`
- All values of `nums` are **unique**.
- `nums` is guaranteed to be rotated at some pivot.
- `-10^4 <= target <= 10^4`

## Solution

### Approach: Binary Search on Rotated Sorted Array

- The array is rotated, so part of it is sorted.
- Perform binary search:
  - If the left part is sorted (`nums[m] >= nums[l]`), check if the target is in this part. If so, discard the right half; otherwise, discard the left half.
  - If the right part is sorted (`nums[m] < nums[l]`), check if the target is in this part. If so, discard the left half; otherwise, discard the right half.
- Repeat until the target is found or the search space is exhausted.

**Complexity**:

- Time: `O(log n)`.

### References

- [Wikipedia/Binary search algorithm](https://en.wikipedia.org/wiki/Binary_search_algorithm)
