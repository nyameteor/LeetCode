# 1. Two Sum

- Difficulty: Easy
- Topics: Array, Hash Table
- Link: https://leetcode.com/problems/two-sum/

## Description

Given an array of integers `nums` and an integer `target`, return _indices of the two numbers such that they add up to `target`_.

You may assume that each input would have **\*exactly\* one solution**, and you may not use the _same_ element twice.

You can return the answer in any order.

**Example 1:**

```
Input: nums = [2,7,11,15], target = 9
Output: [0,1]
Output: Because nums[0] + nums[1] == 9, we return [0, 1].
```

**Example 2:**

```
Input: nums = [3,2,4], target = 6
Output: [1,2]
```

**Example 3:**

```
Input: nums = [3,3], target = 6
Output: [0,1]
```

**Constraints:**

- `2 <= nums.length <= 10^4`
- `-10^9 <= nums[i] <= 10^9`
- `-10^9 <= target <= 10^9`
- **Only one valid answer exists.**

**Follow-up:** Can you come up with an algorithm that is less than `O(n^2) `time complexity?

## Solution

给定一个整数数组 nums 和一个目标值 target，请在该数组中找出和为目标值的*那两个整数*，并返回他们的数组下标。

### Brute Force

暴力解法，两层 for 循环查找，时间复杂度为 O(n^2)。

### Hash Table

该题需要找到整数后返回其数组下标，故 set 不能选用，可以选用 map：

- key 保存数值，value 保存数值所在的下标
- 遍历数组时，通过查询另一个关联的 key 在 map 中是否存在，来判断是否找到答案

在 Java 中，当 map 使用 HashMap 时，查询平均时间复杂度为 O(1)，最坏为 O(n)，再加上外层的数组遍历，得出总时间复杂度：

- Time Average: O(n)
- Time Worst: O(n^2)
