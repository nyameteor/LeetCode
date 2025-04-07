# 78. Subsets

- Difficulty: Medium
- Topics: Array, Backtracking, Bit Manipulation
- Link: https://leetcode.com/problems/subsets/

## Description

Given an integer array `nums` of **unique** elements, return _all possible subsets (the power set)_.

The solution set **must not** contain duplicate subsets. Return the solution in **any order**.

**Example 1:**

```
Input: nums = [1,2,3]
Output: [[],[1],[2],[1,2],[3],[1,3],[2,3],[1,2,3]]
```

**Example 2:**

```
Input: nums = [0]
Output: [[],[0]]
```

**Constraints:**

- `1 <= nums.length <= 10`
- `-10 <= nums[i] <= 10`
- All the numbers of `nums` are **unique**.

## Solution

This problem generates all subsets (the power set) of a given array. Two common approaches:

- **Recursive (Backtracking)**: Explore each index by including or excluding it. Uses DFS with a path that gets backtracked.
- **Iterative**: Start from the empty set and build new subsets by adding each number to existing subsets.

Both run in O(2^n) time and space.
