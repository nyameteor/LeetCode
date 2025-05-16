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

### Key Idea

Use a **monotonic stack** to find the largest rectangle for each bar quickly.

- The stack keeps bar indices with heights in increasing order.
- When a shorter bar appears, pop taller bars from the stack.
- For each popped bar, calculate the rectangle area using the current index and the previous smaller bar's index.

This method helps find the widest rectangle possible for every bar by knowing where smaller bars are on both sides.

### References

- [5ms O(n) Java solution explained (beats 96%)](https://leetcode.com/problems/largest-rectangle-in-histogram/solutions/28902/5ms-o-n-java-solution-explained-beats-96/)
- [AC Python clean solution using stack 76ms](https://leetcode.com/problems/largest-rectangle-in-histogram/solutions/28917/ac-python-clean-solution-using-stack-76ms/)
