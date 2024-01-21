# 1530. Number of Good Leaf Nodes Pairs

- Difficulty: Medium
- Topics: Tree, Depth-First Search, Binary Tree
- Link: https://leetcode.com/problems/number-of-good-leaf-nodes-pairs/

## Description

You are given the `root` of a binary tree and an integer `distance`. A pair of two different **leaf** nodes of a binary tree is said to be good if the length of **the shortest path** between them is less than or equal to `distance`.

Return _the number of good leaf node pairs_ in the tree.

**Example 1:**

![](https://assets.leetcode.com/uploads/2020/07/09/e1.jpg)

```
Input: root = [1,2,3,null,4], distance = 3
Output: 1
Explanation: The leaf nodes of the tree are 3 and 4 and the length of the shortest path between them is 3. This is the only good pair.
```

**Example 2:**

![](https://assets.leetcode.com/uploads/2020/07/09/e2.jpg)

```
Input: root = [1,2,3,4,5,6,7], distance = 3
Output: 2
Explanation: The good pairs are [4,5] and [6,7] with shortest path = 2. The pair [4,6] is not good because the length of ther shortest path between them is 4.
```

**Example 3:**

```
Input: root = [7,1,4,6,null,5,3,null,null,null,null,null,2], distance = 3
Output: 1
Explanation: The only good pair is [2,5].
```

**Constraints:**

- The number of nodes in the `tree` is in the range `[1, 210].`
- `1 <= Node.val <= 100`
- `1 <= distance <= 10`

## Solution

We can create a helper function to:

- Update the number of pairs. Let's name it as `pairsNum`.
- Return the number of leaf nodes for each distance from its parent node. Let's name it as `parentDists`. `parentDists[i]` is the number of leaf nodes, `i` is the distance to parent node.

So for each `helper(root)`, if `root` is null:

- Return `parentDists` filled with `0`.

Else if `root` is a leaf node:

- Return `parentDists` with `parentDists[1] = 1`.

Else:

- `leftDists = helper(root->left)`.
- `rightDists = helper(root->right)`.
- Add number of pairs from current node to `pairsNum`.
- Return `parentDists`, where `parentDists[i] = leftDists[i-1] + rightDists[i-1]`.

References:

- https://leetcode.com/problems/number-of-good-leaf-nodes-pairs/solutions/756198/java-dfs-solution-with-a-twist-100-faster-explained/
- https://leetcode.com/problems/number-of-good-leaf-nodes-pairs/solutions/755784/java-detailed-explanation-post-order-cache-in-array/
