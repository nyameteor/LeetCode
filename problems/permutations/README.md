# 46. Permutations

- Difficulty: Medium
- Topics: Array, Backtracking
- Link: https://leetcode.com/problems/permutations/

## Description

Given an array `nums` of distinct integers, return _all the possible permutations_. You can return the answer in **any order**.

**Example 1:**

```
Input: nums = [1,2,3]
Output: [[1,2,3],[1,3,2],[2,1,3],[2,3,1],[3,1,2],[3,2,1]]
```

**Example 2:**

```
Input: nums = [0,1]
Output: [[0,1],[1,0]]
```

**Example 3:**

```
Input: nums = [1]
Output: [[1]]
```

**Constraints:**

- `1 <= nums.length <= 6`
- `-10 <= nums[i] <= 10`
- All the integers of `nums` are **unique**.

## Solution

给定一个由不同整数组成的数组 `nums`，返回*所有可能的排列*。 您可以按**任意顺序**返回答案。

经典的求全排列，本题适合使用回溯法，因为排列数的数量是 N!，状态空间会非常大，回溯法强调仅使用一份变量去搜索所有可能的状态，可防止 Memory Limit Exceeded Error。

### Backtracking

回溯搜索 = 深度优先遍历 + 状态重置 + 剪枝。

全排列问题中，nums 中的元素在一个 path 中只能使用一次，故每访问到一个 nums[i] 后要将其从可选的元素中移除。

可以 copy 一个 tmpNums 并移除 nums[i] 后再交给下一层搜索，但性能更好的方式是使用一个 visited 数组记录当前已访问过的元素，由较浅节点至较深节点时置 visited[i] 为 true，在由较深节点返回较浅节点时重置 visited[i] 为 false，这样也达成了“可选分支”的回溯。
