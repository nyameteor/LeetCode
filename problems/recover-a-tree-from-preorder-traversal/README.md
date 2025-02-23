# 1028. Recover a Tree From Preorder Traversal

- Difficulty: Hard
- Topics: String, Tree, Depth-First Search, Binary Tree
- Link: https://leetcode.com/problems/recover-a-tree-from-preorder-traversal/

## Description

We run a preorder depth-first search (DFS) on the `root` of a binary tree.

At each node in this traversal, we output `D` dashes (where `D` is the depth of this node), then we output the value of this node.  If the depth of a node is `D`, the depth of its immediate child is `D + 1`.  The depth of the `root` node is `0`.

If a node has only one child, that child is guaranteed to be **the left child**.

Given the output `traversal` of this traversal, recover the tree and return *its* `root`.

**Example 1:**

![](https://assets.leetcode.com/uploads/2024/09/10/recover_tree_ex1.png)

```
Input: traversal = "1-2--3--4-5--6--7"
Output: [1,2,5,3,4,6,7]

```

**Example 2:**

![](https://assets.leetcode.com/uploads/2024/09/10/recover_tree_ex2.png)

```
Input: traversal = "1-2--3---4-5--6---7"
Output: [1,2,5,3,null,6,null,4,null,7]

```

**Example 3:**

![](https://assets.leetcode.com/uploads/2024/09/10/recover_tree_ex3.png)

```
Input: traversal = "1-401--349---90--88"
Output: [1,401,null,349,88,90]

```

**Constraints:**

- The number of nodes in the original tree is in the range `[1, 1000]`.
- `1 <= Node.val <= 10^9`

## Solution

1. **Stack-Based Tree Reconstruction** – A stack is used to keep track of nodes at each depth level.  
2. **Parsing Depth and Value** – Iterate through `traversal`, counting `'-'` for depth and extracting numeric values for node creation.  
3. **Maintaining Parent-Child Relationships** – If the current depth is less than the stack size, trim the stack to maintain the correct parent reference.  
4. **Assigning Children** – Attach the node as the left child if available; otherwise, set it as the right child.  
5. **Iterative and Optimized** – Uses a single pass with constant space (excluding output storage), making it efficient for large inputs.
