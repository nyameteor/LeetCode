# 221. Maximal Square

- Difficulty: Medium
- Topics: Array, Dynamic Programming, Matrix
- Link: https://leetcode.com/problems/maximal-square/

## Description

Given an `m x n` binary `matrix` filled with `0`'s and `1`'s, _find the largest square containing only_ `1`'s _and return its area_.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2020/11/26/max1grid.jpg)

```
Input: matrix = [["1","0","1","0","0"],["1","0","1","1","1"],["1","1","1","1","1"],["1","0","0","1","0"]]
Output: 4
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2020/11/26/max2grid.jpg)

```
Input: matrix = [["0","1"],["1","0"]]
Output: 1
```

**Example 3:**

```
Input: matrix = [["0"]]
Output: 0
```

**Constraints:**

- `m == matrix.length`
- `n == matrix[i].length`
- `1 <= m, n <= 300`
- `matrix[i][j]` is `'0'` or `'1'`.

## Solution

### Brute Force (Accepted)

使用 maximum 记录当前符合题意的最大值。遍历整个 matrix，若当前元素为 1，斜对角（向右与向下）移动，检查斜对角元素以及行列元素是否都为 1，然后检查新添加的行列是否为 1，若到达边界或找到一个为 0 即停止，然后尝试更新 maximum。

**Complexity Analysis**

- Time: O((mn)^2)
- Space: O(1)

### Dynamic Programming

> Todo: 理解得不透，题解待完善

新建一个和 matrix 维度和大小相同的矩阵 memo，初始时所有元素为 0。

memo(i, j) 表示，其右下角是原始矩阵中索引为（i,j) 的单元格的**最大正方形**的边长。（注意，不要设置为从索引 (0, 0) 开始，右下角为 (i, j) 的长方形所包含的最大正方形的边长）

递推关系需要进行观察：
![img](https://leetcode.com/media/original_images/221_Maximal_Square.PNG?raw=true)

递推关系：

```shell
memo(i, j) = min(memo(i-1, j), memo(i-1, j-1), memo(i, j-1)) + matrix[i][j] == '1' ? 1 : 0
```

**Complexity Analysis**

- Time: O(mn), Single pass
- Space: O(mn), 使用了另一个矩阵作为 memo

本题有官方解：https://leetcode.com/problems/maximal-square/solution/
