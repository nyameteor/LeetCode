# 53. Maximum Subarray

- Difficulty: Easy
- Topics: Array, Divide and Conquer, Dynamic Programming
- Link: https://leetcode.com/problems/maximum-subarray/

## Description

Given an integer array `nums`, find the contiguous subarray (containing at least one number) which has the largest sum and return _its sum_.

A **subarray** is a **contiguous** part of an array.

**Example 1:**

```
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
```

**Example 2:**

```
Input: nums = [1]
Output: 1
```

**Example 3:**

```
Input: nums = [5,4,-1,7,8]
Output: 23
```

**Constraints:**

- `1 <= nums.length <= 10^5`
- `-10^4 <= nums[i] <= 10^4`

**Follow up:** If you have figured out the `O(n)` solution, try coding another solution using the **divide and conquer** approach, which is more subtle.

## Solution

### Dynamic Programming

Todo

### Greedy

本题的贪心方法不是很直观：

- 局部最优：若连续和 sum < 0，则放弃当前连续和（将 sum 赋值为 0），从下一个元素开始计算连续和，因为负数会使之后的和更小。
- 全局最优：取最大连续和 max(sum)

```shell
nums                            sums
-2   1  -3   4  -1   2   1
 i                              -2, sum += nums[i], sum < 0 -> sum = 0
     i                           1, sum += nums[i]
     ^   i                      -2, sum += nums[i], sum < 0 -> sum = 0
             i                   4, sum += nums[i]
             ^   i               3, sum += nums[i]
             ^       i           5, sum += nums[i]
             ^           i       6, sum += nums[i]

res = max sum = 6
```

Time: O(n), Space: O(1)
