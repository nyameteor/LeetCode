# 22. Generate Parentheses

- Difficulty: Medium
- Topics: String, Dynamic Programming, Backtracking
- Link: https://leetcode.com/problems/generate-parentheses/

## Description

Given `n` pairs of parentheses, write a function to _generate all combinations of well-formed parentheses_.

**Example 1:**

```
Input: n = 3
Output: ["((()))","(()())","(())()","()(())","()()()"]
```

**Example 2:**

```
Input: n = 1
Output: ["()"]
```

**Constraints:**

- `1 <= n <= 8`

## Solution

给定 n 对括号，编写一个函数来生成格式正确的括号的所有组合。

思路：本题是一个 N 叉树的路径搜索问题。

### Backtracking

回溯算法强调在状态空间很大的情况下，只用一份状态变量去搜索所有可能的状态：

- 在搜索到符合条件的题解时，通常会将其拷贝到结果中。
- 由于全局使用一份状态变量，所以有 _恢复现场_ 和 _撤销选择_ 的需要。

回溯搜索 = 深度优先遍历 + 状态重置 + 剪枝。

递归基本情况：

- 可用左括号数量 == 0 && 可用右括号数量 == 0

剪枝条件：

- 可用左括号数量 > 可用右括号数量

参考：

- official solution：https://leetcode.com/problems/generate-parentheses/solution/
- [@liweiwei1419，严格按照回溯法定义](https://leetcode-cn.com/problems/generate-parentheses/solution/hui-su-suan-fa-by-liweiwei1419/336825)
