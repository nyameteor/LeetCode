# 416. Partition Equal Subset Sum

- Difficulty: Medium
- Topics: Array, Dynamic Programming
- Link: https://leetcode.com/problems/partition-equal-subset-sum/

## Description

Given a **non-empty** array `nums` containing **only positive integers**, find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.

**Example 1:**

```
Input: nums = [1,5,11,5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11].
```

**Example 2:**

```
Input: nums = [1,2,3,5]
Output: false
Explanation: The array cannot be partitioned into equal sum subsets.
```

**Constraints:**

- `1 <= nums.length <= 200`
- `1 <= nums[i] <= 100`

## Solution

### Recursion

我们需要将数组划分为两个子集。这说明对于数组的每一个元素，我们可以将其放在子集 1，或者放在子集 2。

- 由于我们只关心子集内元素的总和是否相当，所以用 `sum1` 保存子集 1 的总和，`sum2` 保存子集 2 的总和。
- 遍历数组，对于每一个元素，可以选择将其放在子集 1 并增加 `sum1`，或者放在子集 2 并增加 `sum2`。
- 每次到达数组末尾时，检查并返回 `sum1` 和 `sum2` 是否相等。
- 如果最终没有任何一条路径能给出相等的总和，返回 false；反之，返回 true。

Refer: https://leetcode.com/problems/partition-equal-subset-sum/discuss/1624939

### Recursion + Memoization

Todo
