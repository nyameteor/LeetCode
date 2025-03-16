# 215. Kth Largest Element in an Array

- Difficulty: Medium
- Topics: Array, Divide and Conquer, Sorting, Heap(Priority Queue), Quickselect
- Link: https://leetcode.com/problems/kth-largest-element-in-an-array/

## Description

Given an integer array `nums` and an integer `k`, return _the_ `kth` _largest element in the array_.

Note that it is the `kth` largest element in the sorted order, not the `kth` distinct element.

**Example 1:**

```
Input: nums = [3,2,1,5,6,4], k = 2
Output: 5
```

**Example 2:**

```
Input: nums = [3,2,3,1,2,4,5,5,6], k = 4
Output: 4
```

**Constraints:**

- `1 <= k <= nums.length <= 10^4`
- `-10^4 <= nums[i] <= 10^4`

## Solution

### Approach: Heap (Priority Queue)

- **Use a Min-Heap** to maintain the top `k` largest elements.
- Push elements into the heap until it contains `k` elements. Once the heap exceeds `k` elements, pop the smallest element (the root of the min-heap).
- After processing all elements, the root of the heap will be the `k`th largest element.

**Time Complexity**: **O(n log k)**, where `n` is the length of the array and `k` is the size of the heap. Each heap operation (push/pop) takes `O(log k)` time.

**Space Complexity**: **O(k)**, as we store at most `k` elements in the heap.
