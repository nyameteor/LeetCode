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

### Binary Search

本题输入的数组经过了旋转，故为局部有序。要在局部有序的数组上使用二分查找，需要添加判断条件查找范围内的数组是否有序。本题中可以通过比较 `nums[left]` 和 `nums[mid]` 的关系，来判断哪一个范围内的数组是有序的：

- if nums[left] <= nums[mid], 前半段数组 [left, mid] 有序。若 target 在 [left, mid] 范围内，在前半段搜索，否则在后半段搜索。
- if nums[left] > nums[mid], 后半段数组 [mid+1, right] 有序。若 target 在 [mid+1, right] 范围内，在后半段搜索，否则在前半段搜索。

Time: O(log n), Space: O(1)

更多：[Wikipedia/Binary search algorithm](https://en.wikipedia.org/wiki/Binary_search_algorithm)
