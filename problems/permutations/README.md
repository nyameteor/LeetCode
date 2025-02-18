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

1. **Backtracking Approach**: We use DFS to explore all possible permutations by adding elements sequentially.  
2. **Tracking Usage**: A `used` boolean slice ensures each element is used only once per permutation.  
3. **Recursive Exploration**: If the path reaches the length of `nums`, we add a clone of it to `res`.  
4. **Backtracking Step**: After recursion, we reset `used[i]` to explore other possibilities.  
