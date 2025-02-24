# 889. Construct Binary Tree from Preorder and Postorder Traversal

- Difficulty: Medium
- Topics: Array, Hash Table, Divide and Conquer, Tree, Binary Tree
- Link: https://leetcode.com/problems/construct-binary-tree-from-preorder-and-postorder-traversal/

## Description

Given two integer arrays, `preorder` and `postorder` where `preorder` is the preorder traversal of a binary tree of **distinct** values and `postorder` is the postorder traversal of the same tree, reconstruct and return *the binary tree*.

If there exist multiple answers, you can **return any** of them.

**Example 1:**

![](https://assets.leetcode.com/uploads/2021/07/24/lc-prepost.jpg)

```
Input: preorder = [1,2,4,5,3,6,7], postorder = [4,5,2,6,7,3,1]
Output: [1,2,3,4,5,6,7]

```

**Example 2:**

```
Input: preorder = [1], postorder = [1]
Output: [1]

```

**Constraints:**

- `1 <= preorder.length <= 30`
- `1 <= preorder[i] <= preorder.length`
- All the values of `preorder` are **unique**.
- `postorder.length == preorder.length`
- `1 <= postorder[i] <= postorder.length`
- All the values of `postorder` are **unique**.
- It is guaranteed that `preorder` and `postorder` are the preorder traversal and postorder traversal of the same binary tree.

## Solution

The problem requires reconstructing a binary tree using its **preorder** and **postorder** traversals.  

- **Preorder** traversal: `[root, left subtree, right subtree]`  
- **Postorder** traversal: `[left subtree, right subtree, root]`  

Given the properties of these traversals:  

1. The first element of `preorder` is always the root.  
2. The last element of `postorder` is also always the root.  
3. The second element of `preorder` is the root of the left subtree (if it exists).  
4. Using this, we can determine the size of the left subtree by finding this element in `postorder`.  

This forms a **divide and conquer** approach where we recursively reconstruct left and right subtrees.  

### **Steps**  

1. **Base Cases**: If the tree is empty (`n == 0`), return `nil`. If it has only one node, return that node.  
2. **Identify Left Subtree**: The second element in `preorder` is the root of the left subtree (`leftChildRoot`).  
3. **Find Left Subtree Size**: Locate `leftChildRoot` in `postorder` to determine the size of the left subtree.  
4. **Recursive Construction**:  
   - The left subtree is built from `preorder[1:1+leftChildSize]` and `postorder[:leftChildSize]`.  
   - The right subtree is built from `preorder[1+leftChildSize:]` and `postorder[leftChildSize:n-1]`.  
5. **Construct and Return Root**: Recursively build the tree and return the root node.  
