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

### Sorting

先对 intervals 排序（可以使用快速排序），按照 start_i 的大小，从小到大的顺序。

```shell
[(1, 9), (2, 5), (0, 3), (0, 1)]

[(0, 1), (0, 3), (1, 9), (2, 5)] # sorted
```

然后两两比较 ints[i] 和 ints[j]，若 ints[i].end < ints[j].start，则直接将 ints[j] 加入结果；否则，将 merge 后的 interval 添加到结果中。

merge ints[i] 和 ints[j] 时会有如下可能，按照情况 merge 即可：

```shell
ints[i].end < ints[j].start ≤ ints[j].end < ints[k].start
                              ints[i].end ≥ ints[k].start​
```

**Complexity Analysis**

- Time complexity: O(nlog(n))

  快速排序的平均时间复杂度为 O(nlog(n))，另外我们对数组进行了简单的线性扫描，于是得到这个时间复杂度。

- Space complexity: O(log(n)) or O(n)

  如果对 intervals 进行就地 merge，只需要常数级的额外空间，但快速排序需要 O(log(n)) 的空间；如果不使用就地 merge，则需分配额外 O(n) 的线性空间来存储结果。

### Connected Components

Todo

本题有官方解：https://leetcode.com/problems/merge-intervals/solution/
