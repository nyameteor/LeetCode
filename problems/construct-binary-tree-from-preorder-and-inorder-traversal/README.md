# 105. Construct Binary Tree from Preorder and Inorder Traversal

- Difficulty: Medium
- Topics: Array, Hash Table, Divide and Conquer, Tree, Binary Tree
- Link: https://leetcode.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

## Description

Given two integer arrays `preorder` and `inorder` where `preorder` is the preorder traversal of a binary tree and `inorder` is the inorder traversal of the same tree, construct and return _the binary tree_.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/02/19/tree.jpg)

```
Input: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
Output: [3,9,20,null,null,15,7]
```

**Example 2:**

```
Input: preorder = [-1], inorder = [-1]
Output: [-1]
```

**Constraints:**

- `1 <= preorder.length <= 3000`
- `inorder.length == preorder.length`
- `-3000 <= preorder[i], inorder[i] <= 3000`
- `preorder` and `inorder` consist of **unique** values.
- Each value of `inorder` also appears in `preorder`.
- `preorder` is **guaranteed** to be the preorder traversal of the tree.
- `inorder` is **guaranteed** to be the inorder traversal of the tree.

## Solution

### Recursion

Preorder:

```shell
{root, {preorder of left subtree}, {preorder of right subtree}}
```

Inorder:

```shell
{{inorder of left subtree}, root, {inorder of right subtree}}
```

- From `preorder`, we can find the root node in current tree is `preorder[0]`.
- Find this root node in `inorder`, we can find the range of `left subtree` and `right subtree`.
- The base case is `preorder` and `inorder` both empty.

Denote `i` as the lenght of `left subtree`:

|          | range of left subtree | range of right subtree |
| -------- | --------------------- | ---------------------- |
| preorder | `[1, i + 1)`          | `[i + 1, end)`         |
| inorder  | `[0, i)`              | `[i + 1, end)`         |

Denote `b1, e1` and `b2, e2` as the begin and end index of current range in preorder and inorder respectively.

We can solve this problem recursively:

```plaintext
F(preorder, b1, e1, inorder, b2, e2) =
        null, if preorder and inorder is empty
        node = {
            val = preorder[b1]
            left = F(preorder, b1 + 1, b1 + i + 1, inorder, b2, b2 + i)
            right = F(preorder, b1 + i + 1, e1, inorder, b2 + i + 1, e2)
        }
```

Example:

```plaintext
index       0   1   2   3   4
preorder    3   9   20  15  7
            ^
inorder     9   3   15  20  7
                ^
current root: 3

--- left subtree
preorder    9
            ^
inorder     9
            ^
current root: 9

--- right subtree
preorder    20  15  7
            ^
inorder     15  20  7
                ^
current root: 20

--- left subtree of right subtree...
```
