# 448. Find All Numbers Disappeared in an Array

- Difficulty: Easy
- Topics: Array, Hash Table
- Link: https://leetcode.com/problems/find-all-numbers-disappeared-in-an-array/

## Description

给定一个包含 `n` 个整数的数组 `nums`，其中 `nums[i]` 在范围 `[1, n]` 内，要求返回一个属于于 `[1, n]` 范围内且未在 `nums` 中出现的整数的数组。

要求不使用额外空间且时间复杂度为 `O(n)`。

```shell
Example 1:

Input: nums = [4,3,2,7,8,2,3,1]
Output: [5,6]

Example 2:

Input: nums = [1,1]
Output: [2]
```

Constraints:

- n == nums.length
- 1 <= n <= 10^5
- 1 <= nums[i] <= n

## Solution

因为不能使用额外的空间，所以需要在原数组上标注出出现过的数字。一个可行的方法是两次遍历：

- 遍历数组，将 `nums[ |nums[i| - 1 ]` 对应的元素置为负数。
- 遍历数组，若 `nums[i] > 0`，说明其 `i + 1` 未出现过，将其添加到返回结果。
