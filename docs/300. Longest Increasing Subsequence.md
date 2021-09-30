# 300. Longest Increasing Subsequence

- Difficulty: Medium
- Topics: Array, Binary Search, Dynamic Programming
- Link: https://leetcode.com/problems/longest-increasing-subsequence/

## Description

Given an integer array `nums`, return the length of the longest strictly increasing subsequence.

A **subsequence** is a sequence that can be derived from an array by deleting some or no elements without changing the order of the remaining elements. For example, `[3,6,2,7]` is a subsequence of the array `[0,3,1,6,2,2,7]`.

**Example 1:**

```
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
```

**Example 2:**

```
Input: nums = [0,1,0,3,2,3]
Output: 4
```

**Example 3:**

```
Input: nums = [7,7,7,7,7,7,7]
Output: 1
```

**Constraints:**

- `1 <= nums.length <= 2500`
- `-10^4 <= nums[i] <= 10^4`

**Follow up:** Can you come up with an algorithm that runs in `O(n log(n))` time complexity?

## Solution

给定一个无序的整数数组，找到其中最长上升子序列的长度。

### Dynamic Programming

动态规划很适合解决子序列问题，但子结构需要设置合理，否则无法找出递推关系。

将 table[i] 表示为**包含** i 位置的元素的最长上升子序列长度（不能表示为总的最长上升子序列长度，即不论是否包含 i 位置元素）。

设 j 属于 [0, i)，若 j 位置的元素包含在以 i 位置结尾的最长上升子序列中：

由于该子序列是**上升**的，故满足以下条件：

- nums[j] < nums[i]

同时，由于该子序列是**最长**的，故满足以下条件：

- maximum = max_of table[j]
- table[i] = maximum + 1

故可以得到递推关系如下：

```shell
maximum = max_of table[j]
table[i] = maximum + 1

where   0 <= j < i
        nums[j] < nums[i]
```

在程序中 `maximum` 作为中间变量可以省略，递推关系可写作 `table[i] = max(table[i], talbe[j]), j in [0, i)`

最后，对所有 table[i] 取最大值，得到该题结果。

Todo: 寻求清晰的数学表达式，如：取一个序列（如一段数组）中的最大值

Refer: [Triple-Z/LeetCode/docs/300](https://github.com/Triple-Z/LeetCode/blob/master/docs/300.%20Longest%20Increasing%20Subsequence%20%E6%9C%80%E9%95%BF%E4%B8%8A%E5%8D%87%E5%AD%90%E5%BA%8F%E5%88%97.md)

### Binary Search

Todo
