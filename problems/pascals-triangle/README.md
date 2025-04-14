# 118. Pascal's Triangle

- Difficulty: Easy
- Topics: Array, Dynamic Programming
- Link: https://leetcode.com/problems/pascals-triangle/

## Description

Given an integer `numRows`, return the first numRows of **Pascal's triangle**.

In **Pascal's triangle**, each number is the sum of the two numbers directly above it as shown:

![](https://upload.wikimedia.org/wikipedia/commons/0/0d/PascalTriangleAnimated2.gif)

**Example 1:**

```
Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]
```

**Example 2:**

```
Input: numRows = 1
Output: [[1]]
```

**Constraints:**

- `1 <= numRows <= 30`

## Solution

### Key Idea

Build Pascal's Triangle row by row:

- Each row starts and ends with 1.
- Each inner element is the sum of the two elements above it: `res[i][j] = res[i-1][j-1] + res[i-1][j]`
