# 11. Container With Most Water

- Difficulty: Medium
- Topics: Array, Two Pointers, Greedy
- Link: https://leetcode.com/problems/container-with-most-water/

## Description

You are given an integer array `height` of length `n`. There are `n` vertical lines drawn such that the two endpoints of the `ith` line are `(i, 0)` and `(i, height[i])`.

Find two lines that together with the x-axis form a container, such that the container contains the most water.

Return _the maximum amount of water a container can store_.

**Notice** that you may not slant the container.

**Example 1:**

![img](https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/17/question_11.jpg)

```
Input: height = [1,8,6,2,5,4,8,3,7]
Output: 49
Explanation: The above vertical lines are represented by array [1,8,6,2,5,4,8,3,7]. In this case, the max area of water (blue section) the container can contain is 49.
```

**Example 2:**

```
Input: height = [1,1]
Output: 1
```

**Constraints:**

- `n == height.length`
- `2 <= n <= 10^5`
- `0 <= height[i] <= 10^4`

## Solution

### Two Pointers

We have an array `A` of length `N`, `A` stores the height of vertical lines.

1. Let the initial container have the maximum width, we start from outermost lines.
   - initial `i = 0`
   - initial `j = N - 1`
2. Go to container with shorter width, if there is a vertical line longer than current containers shorter line.
   - if A[i] < A[j], then i++
   - else, then j--
   - when i == j, end the loop
