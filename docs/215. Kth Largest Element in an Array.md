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

在未排序的数组中找到第 k 个最大的元素。请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

### Heap(Priority Queue)

采用一个小顶堆（min heap），维护其容量为 k，来保存前 k 个最大的元素，堆顶（root node）即为第 k 个最大的元素（堆中的最小值）。

- 当堆的规模等于 k，将新元素与堆顶比较
  - 若新元素大于堆顶，则将堆顶推出，将新元素入堆
  - 若新元素小于等于堆顶，则略过
- 当堆的规模小于 k，将新元素入堆
- 数组遍历完后，返回堆顶

关于堆的更多内容参考：[Heap (data structure)](<https://en.wikipedia.org/wiki/Heap_(data_structure)>)

### Quick Select

Todo
