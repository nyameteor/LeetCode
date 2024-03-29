# 42. Trapping Rain Water

- Difficulty: Hard
- Topics: Array, Two Pointers, Dynamic Programming, Stack, Monotonic Stack
- Link: https://leetcode.com/problems/trapping-rain-water/

## Description

Given `n` non-negative integers representing an elevation map where the width of each bar is `1`, compute how much water it can trap after raining.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2018/10/22/rainwatertrap.png)

```
Input: height = [0,1,0,2,1,0,1,3,2,1,2,1]
Output: 6
Explanation: The above elevation map (black section) is represented by array [0,1,0,2,1,0,1,3,2,1,2,1]. In this case, 6 units of rain water (blue section) are being trapped.
```

**Example 2:**

```
Input: height = [4,2,0,3,2,5]
Output: 9
```

**Constraints:**

- `n == height.length`
- `1 <= n <= 2 * 10^4`
- `0 <= height[i] <= 10^5`

## Solution

给定 n 个非负整数表示每个宽度为 1 的柱子的高度图，计算按此排列的柱子，下雨之后能接多少雨水。

### Two Pointer

按列求每一列雨水的高度：

当前列雨水面积：min(左边柱子的最高高度，右边柱子的最高高度) - 当前柱子高度。(宽度等于 1 故在数值计算上省略)

遍历所有的列，使用双指针求出左右两边柱子最高的高度，按照上面的式子计算并累积结果即可。

注意边界情况(第一个柱子和最后一个柱子不能接雨水)。

Time: O(n^2), Space: O(1)

### Dynamic Programming

双指针法中寻找左右柱子的最高高度有重复计算，可以用动态规划来保存子结果：

```shell
# 使用数组 maxLeft 和 maxRight 保存位置 i 上的左边与右边柱子的最高高度

# 递推关系
maxLeft[i] = max(maxLeft[i - 1], height[i]);
maxRight[i] = max(maxRight[i + 1], height[i]);

# 使用 tabulation 法记录 maxLeft 和 maxRight
```

参考：[leetcode-master/problems/0042.接雨水.md](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0042.%E6%8E%A5%E9%9B%A8%E6%B0%B4.md)

### Monotonic Stack

Todo
