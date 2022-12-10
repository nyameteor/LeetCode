# 1339. Maximum Product of Splitted Binary Tree

- Difficulty: Medium
- Topics: Tree, Depth-First Search, Binary Tree
- Link: https://leetcode.com/problems/maximum-product-of-splitted-binary-tree/

## Description

Given the `root`
of a binary tree, split the binary tree into two subtrees by removing
one edge such that the product of the sums of the subtrees is maximized.

Return _the maximum product of the sums of the two subtrees_. Since the answer may be too large, return it **modulo** `10^9 + 7`.

**Note** that you need to maximize the answer before taking the mod and not after taking it.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/01/21/sample_1_1699.png)

```
**Input:** root = [1,2,3,4,5,6]
**Output:** 110
**Explanation:** Remove the red edge and get 2 binary trees with sum 11 and 10. Their product is 110 (11\*10)
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2020/01/21/sample_2_1699.png)

```
**Input:** root = [1,null,2,3,4,null,null,5,6]
**Output:** 90
**Explanation:** Remove the red edge and get 2 binary trees with sum 15 and 6.Their product is 90 (15\*6)
```

**Constraints:**

- The number of nodes in the tree is in the range `[2, 5 * 10^4]`.
- `1 <= Node.val <= 10^4`

## Solution

### Two Pass

1. Get the `totalSum` of the tree.
2. lookup the sums of subtree, and try to find a `targetSum`, which can make the maximum product `targetSum * (totalSum - targetSum)`.

### Two Pass with memo

In above approach, we compute the sums of each subtree twice, which can cause additional time complexity.

We can avoid this by memorizing them during first pass:

- We can store the sum in a HashMap `{node : sumOfThisTree}`.
- Or we can store the sum in the tree node directly, by setting the `node.Val` to `sumOfThisTree`.

**References:**

- https://leetcode.com/problems/maximum-product-of-splitted-binary-tree/solutions/1004644/golang-simple-solution-beat-95/
