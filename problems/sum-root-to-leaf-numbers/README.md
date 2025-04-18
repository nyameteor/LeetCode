# 129. Sum Root to Leaf Numbers

- Difficulty: Medium
- Topics: Tree, Depth-First Search, Binary Tree
- Link: https://leetcode.com/problems/sum-root-to-leaf-numbers/

## Description

You are given the `root` of a binary tree containing digits from `0` to `9` only.

Each root-to-leaf path in the tree represents a number.

- For example, the root-to-leaf path `1 -> 2 -> 3` represents the number `123`.

Return _the total sum of all root-to-leaf numbers_. Test cases are generated so that the answer will fit in a **32-bit** integer.

A **leaf** node is a node with no children.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/02/19/num1tree.jpg)

```
Input: root = [1,2,3]
Output: 25
Explanation:
The root-to-leaf path 1->2 represents the number 12.
The root-to-leaf path 1->3 represents the number 13.
Therefore, sum = 12 + 13 = 25.
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2021/02/19/num2tree.jpg)

```
Input: root = [4,9,0,5,1]
Output: 1026
Explanation:
The root-to-leaf path 4->9->5 represents the number 495.
The root-to-leaf path 4->9->1 represents the number 491.
The root-to-leaf path 4->0 represents the number 40.
Therefore, sum = 495 + 491 + 40 = 1026.
```

**Constraints:**

- The number of nodes in the tree is in the range `[1, 1000]`.
- `0 <= Node.val <= 9`
- The depth of the tree will not exceed `10`.

## Solution

### Approach: Recursive DFS

Use DFS to accumulate the path value. At each leaf, return the value; sum them all for the final result.

### References

- [[C++/Python] Recursive & Iterative DFS + BFS + Morris Traversal O(1)](https://leetcode.com/problems/sum-root-to-leaf-numbers/solutions/1556417/C++Python-Recursive-and-Iterative-DFS-+-BFS-+-Morris-Traversal-O/)
