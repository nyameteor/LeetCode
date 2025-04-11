# 56. Merge Intervals

- Difficulty: Medium
- Topics: Array, Sorting
- Link: https://leetcode.com/problems/merge-intervals/

## Description

Given an array of `intervals` where `intervals[i] = [starti, endi]`, merge all overlapping intervals, and return _an array of the non-overlapping intervals that cover all the intervals in the input_.

**Example 1:**

```
Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
Output: [[1,6],[8,10],[15,18]]
Explanation: Since intervals [1,3] and [2,6] overlaps, merge them into [1,6].
```

**Example 2:**

```
Input: intervals = [[1,4],[4,5]]
Output: [[1,5]]
Explanation: Intervals [1,4] and [4,5] are considered overlapping.
```

**Constraints:**

- `1 <= intervals.length <= 10^4`
- `intervals[i].length == 2`
- `0 <= start_i <= end_i <= 10^4`

## Solution

### Approach: Sort & Merge

Sort intervals, then iterate and merge if the current interval overlaps with the last added interval.

**Key Idea:** Two intervals `[a, b]` and `[c, d]` overlap if `c <= b`. Merge them into `[a, max(b, d)]`.

**Complexity:**

- **Time:** `O(n log n)`.
- **Space:** `O(log n)`.
