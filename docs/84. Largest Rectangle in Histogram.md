# 84. Largest Rectangle in Histogram

- Difficulty: Hard
- Topics: Array, Stack, Monotonic Stack
- Link: https://leetcode.com/problems/largest-rectangle-in-histogram/

## Description

Given an array of integers `heights` representing the histogram's bar height where the width of each bar is `1`, return _the area of the largest rectangle in the histogram_.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/01/04/histogram.jpg)

```
Input: heights = [2,1,5,6,2,3]
Output: 10
Explanation: The above is a histogram where width of each bar is 1.
The largest rectangle is shown in the red area, which has an area = 10 units.
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2021/01/04/histogram-1.jpg)

```
Input: heights = [2,4]
Output: 4
```

**Constraints:**

- `1 <= heights.length <= 10^5`
- `0 <= heights[i] <= 10^4`

## Solution

### Monotonic Stack

使用栈 `stH` 保存高度 `h`，栈 `stX` 保存其左侧最近的大于或等于 `h` 的元素的索引，并维持其单调性，使 `stH` 自顶向下的元素值递减。

遍历 `heights` 数组，当插入 `heights[i] <= stH.top()` 时，弹出 `stH` 的栈顶，此时对应的区域大小为 `stH.pop() * (i - stX.pop())`。

```shell
index     0   1   2   3   4   5   6 (dummy index to pop all left in stack)
heights   2   1   5   6   2   3   0

                      6       3
                  5   5   2   2
stH       2   1   1   1   1   1

popH          2           6       3
                          5       2
                                  1

                      3       5
                  2   2   4   4
stX       0   0   0   0   0   0

popX          0           3       5
                          2       4
                                  0

area          2           6       3
                          10      4
                                  6
area = = popH * (index - popX)
maxArea = 10
```

Refer: https://leetcode.com/problems/largest-rectangle-in-histogram/discuss/951439/Monotonic-Stack-(C%2B%2B)

单调栈即为满足单调性的栈结构：

- push:
  将元素插入单调栈时，需要满足该元素插入后，整个栈后仍满足单调性的前提下，弹出最少的元素。
  e.g. 一个单调栈自顶向下的元素为：`{10, 11, 15, 16}`，插入 `14` 时为了保证单调性，需要先弹出 `10`, `11`，操作后栈的元素为：`{14, 15, 16}`。
- pop:
  从栈顶弹出元素，该元素满足单调性的某一端，如最小值。

单调栈适合解决特定的问题，如下一个更大或更小的元素（Next greater or smaller element）。

Refer: https://oi-wiki.org/ds/monotonous-stack/
