# 655. Print Binary Tree

- Difficulty: Medium
- Topics: Tree, Depth-First Search, Breadth-First Search, Binary Tree
- Link: https://leetcode.com/problems/print-binary-tree/

## Description

Given the `root` of a binary tree, construct a **0-indexed** `m x n` string matrix `res` that represents a **formatted layout** of the tree. The formatted layout matrix should be constructed using the following rules:

- The **height** of the tree is `height` and the number of rows `m` should be equal to `height + 1`.
- The number of columns `n` should be equal to `2height+1 - 1`.
- Place the **root node** in the **middle** of the **top row** (more formally, at location `res[0][(n-1)/2]`).
- For each node that has been placed in the matrix at position `res[r][c]`, place its **left child** at `res[r+1][c-2height-r-1]` and its **right child** at `res[r+1][c+2height-r-1]`.
- Continue this process until all the nodes in the tree have been placed.
- Any empty cells should contain the empty string `""`.

Return _the constructed matrix_ `res`.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/05/03/print1-tree.jpg)

```
Input: root = [1,2]
Output:
[["","1",""],
 ["2","",""]]
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2021/05/03/print2-tree.jpg)

```
Input: root = [1,2,3,null,4]
Output:
[["","","","1","","",""],
 ["","2","","","","3",""],
 ["","","4","","","",""]]
```

**Constraints:**

- The number of nodes in the tree is in the range `[1, 2^10]`.
- `-99 <= Node.val <= 99`
- The depth of the tree will be in the range `[1, 10]`.

## Solution

### DFS/BFS Two Pass

- 由于每个节点放置在矩阵中的的位置依赖于树的高度，所以先用一次 DFS/BFS 获取树的深度。
- 再用 DFS/BFS 保存每一个节点到结果矩阵中。
  - 以 BFS 为例，由于左子树和右子树放置在矩阵中的位置，与当前的节点放置的位置有对应关系，所以 Queue 中除了保存当前节点，还需要保存当前节点的位置信息（row 与 column）。由于 row 即为深度，在 BFS 的过程中可求，所以 Queue 中可以额外保存当前节点的 column，示例：`queue<pair<TreeNode*, int>>`
