# 450. Delete Node in a BST

- Difficulty: Medium
- Topics: Tree, Binary Search Tree, Binary Tree
- Link: https://leetcode.com/problems/delete-node-in-a-bst/

## Description

Given a root node reference of a BST and a key, delete the node with the given key in the BST. Return the root node reference (possibly updated) of the BST.

Basically, the deletion can be divided into two stages:

1. Search for a node to remove.
2. If the node is found, delete the node.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2020/09/04/del_node_1.jpg)

```
Input: root = [5,3,6,2,4,null,7], key = 3
Output: [5,4,6,2,null,null,7]
Explanation: Given key to delete is 3. So we find the node with value 3 and delete it.
One valid answer is [5,4,6,2,null,null,7], shown in the above BST.
Please notice that another valid answer is [5,2,6,null,4,null,7] and it's also accepted.
```

**Example 2:**

```
Input: root = [5,3,6,2,4,null,7], key = 0
Output: [5,3,6,2,4,null,7]
Explanation: The tree does not contain a node with value = 0.
```

**Example 3:**

```
Input: root = [], key = 0
Output: []
```

**Constraints:**

- The number of nodes in the tree is in the range `[0, 10^4]`.
- `-10^5 <= Node.val <= 10^5`
- Each node has a **unique** value.
- `root` is a valid binary search tree.
- `-10^5 <= key <= 10^5`

**Follow up:** Could you solve it with time complexity `O(height of tree)`?

## Solution

solution description here.

### Recursive Preorder Traverse

TODO

| Case                                 | Description                          |
| ------------------------------------ | ------------------------------------ |
| `T` not found in BST                 | directly return root of BST as it is |
| `T` has no children                  |                                      |
| `T` has no left children             |                                      |
| `T` has no right children            |                                      |
| `T` has both left and right children |                                      |

Refer：
https://leetcode.com/problems/delete-node-in-a-bst/discuss/1591176/C%2B%2B-Simple-Solution-w-Images-and-Detailed-Explanation-or-Iterative-and-Recursive-Approach
