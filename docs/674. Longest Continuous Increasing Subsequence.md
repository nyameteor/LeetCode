# 674. Longest Continuous Increasing Subsequence

- Difficulty: Easy
- Topics: Array
- Link: https://leetcode.com/problems/longest-continuous-increasing-subsequence/

## Description

Given an unsorted array of integers `nums`, return _the length of the longest **continuous increasing subsequence** (i.e. subarray)_. The subsequence must be **strictly** increasing.

A **continuous increasing subsequence** is defined by two indices `l` and `r` (`l < r`) such that it is `[nums[l], nums[l + 1], ..., nums[r - 1], nums[r]]` and for each `l <= i < r`, `nums[i] < nums[i + 1]`.

**Example 1:**

```
Input: nums = [1,3,5,4,7]
Output: 3
Explanation: The longest continuous increasing subsequence is [1,3,5] with length 3.
Even though [1,3,5,7] is an increasing subsequence, it is not continuous as elements 5 and 7 are separated by element
4.
```

**Example 2:**

```
Input: nums = [2,2,2,2,2]
Output: 1
Explanation: The longest continuous increasing subsequence is [2] with length 1. Note that it must be strictly
increasing.
```

**Constraints:**

- `1 <= nums.length <= 10^4`
- `-10^9 <= nums[i] <= 10^9`

## Solution

### Sliding Window

使用两个指针 l，r；r 作为遍历指针，l 作为锚点（anchor）记录当前递增子序列开始位置的索引，当前的 sliding window 即为 A[l, r]，子序列长度为 r - l + 1。

若 nums[r-1] < nums[r]，说明从该位置 r 开始是一个新的子序列，故更新锚点 l = r。

遍历中每一次中得到的子序列长度作为候选值去更新结果，最终即得到结果。

```shell
 1 3 5 2 4 2
l,r

 1 3 5 1 4 2
 l r

 1 3 5 1 4 2
 l   r

 1 3 5 1 4 2    nums[r-1] < nums[r], l = r
      l,r

 1 3 5 1 4 2
       l r

 1 3 5 1 4 2    nums[r-1] < nums[r], l = r
          l,r
```

滑动窗口是本题的官方解：https://leetcode.com/problems/longest-continuous-increasing-subsequence/solution/
