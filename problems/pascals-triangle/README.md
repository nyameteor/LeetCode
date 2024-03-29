# 118. Pascal's Triangle

- Difficulty: Easy
- Topics: Array, Dynamic Programming
- Link: https://leetcode.com/problems/pascals-triangle/

## Description

给定一个非负整数 `numRows`，返回杨辉三角形的前 `numRows` 行。

![](https://upload.wikimedia.org/wikipedia/commons/0/0d/PascalTriangleAnimated2.gif)

在杨辉三角（帕斯卡三角）中，每个数是它左上方和右上方的数的和。

```shell
Example 1:

Input: numRows = 5
Output: [[1],[1,1],[1,2,1],[1,3,3,1],[1,4,6,4,1]]

Example 2:

Input: numRows = 1
Output: [[1]]
```

Constraints:

- 1 <= numRows <= 30

## Solution

如题意，比较直观的方式是按照递推关系，将每一行都按照上一行计算出来，然后在首尾补 1。

个人认为一个更好的方式是在上一行的前后补 0，这样计算该行时，不需要考虑 edge case:

```shell
              0
            0   0
          0   1   0
        0   1   1   0
      0   1   2   1   0
    0   1   3   3   1   0
  0   1   4   6   4   1   0
0   0   0   0   0   0   0   0
```

- n = numRows
- time: O(n^2)
- space: O(n^2)
