# 347. Top K Frequent Elements

- Difficulty: Medium
- Topics: Array, Hash Table, Divide and Conquer, Sorting, Heap(Priority Queue), Bucket Sort, Counting, Quickselect
- Link: https://leetcode.com/problems/top-k-frequent-elements/

## Description

Given an integer array `nums` and an integer `k`, return _the_ `k` _most frequent elements_. You may return the answer in **any order**.

**Example 1:**

```
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
```

**Example 2:**

```
Input: nums = [1], k = 1
Output: [1]
```

**Constraints:**

- `1 <= nums.length <= 10^5`
- `k` is in the range `[1, the number of unique elements in the array]`.
- It is **guaranteed** that the answer is **unique**.

**Follow up:** Your algorithm's time complexity must be better than `O(n log n)`, where n is the array's size.

## Solution

### Approach: Min-Heap

This problem can be solved efficiently using a **min-heap**.

1. **Frequency Map**: Count the frequency of each element in the array.
2. **Min-Heap**: Use a min-heap to store the top `k` most frequent elements. The heap will store the element and its frequency.
   - When the heap size exceeds `k`, remove the least frequent element.
3. **Result Extraction**: Extract elements from the heap to form the result.

**Complexity**:

- Time: `O(n log k)`, where `n` is the number of elements in the array and `k` is the number of top frequent elements.
- Space: `O(n)`, for the frequency map and heap storage.
