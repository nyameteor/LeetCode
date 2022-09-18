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

### Recursion with Backtracking

对于子集而言，数组集合中的元素只有两种状态：选中和忽略，因此可以按状态递归遍历数组元素，求得所有子集：

```shell
# `*` represent element not slected
                                *
            +-------------------+-------------------+
            *                                       1
     +------+------+                         +------+------+
     *             2                         *             2
 +--+--+        +--+--+                  +--+--+        +--+--+
 *     3        *     3                  *     3        *     3

{}    {3}      {2}   {2,3}              {1}   {1,3}   {1,2} {1,2,3}
```

然后，可以通过回溯法优化内存使用（全局使用同一份状态变量）。

### Cascading

Todo

### Bit Manipulation

Todo
