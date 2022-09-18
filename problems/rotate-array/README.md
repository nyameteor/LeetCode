# 189. Rotate Array

- Difficulty: Medium
- Topics: Array, Math, Two Pointers
- Link: https://leetcode.com/problems/rotate-array/

## Description

Given an array, rotate the array to the right by `k` steps, where `k` is non-negative.

**Example 1:**

```
Input: nums = [1,2,3,4,5,6,7], k = 3
Output: [5,6,7,1,2,3,4]
Explanation:
rotate 1 steps to the right: [7,1,2,3,4,5,6]
rotate 2 steps to the right: [6,7,1,2,3,4,5]
rotate 3 steps to the right: [5,6,7,1,2,3,4]
```

**Example 2:**

```
Input: nums = [-1,-100,3,99], k = 2
Output: [3,99,-1,-100]
Explanation:
rotate 1 steps to the right: [99,-1,-100,3]
rotate 2 steps to the right: [3,99,-1,-100]
```

**Constraints:**

- `1 <= nums.length <= 10^5`
- `-2^31 <= nums[i] <= 2^31 - 1`
- `0 <= k <= 10^5`

**Follow up:**

- Try to come up with as many solutions as you can. There are at least **three** different ways to solve this problem.
- Could you do it in-place with `O(1)` extra space?

## Solution

### Extra Space

使用额外空间。简单直观，但空间性能较差。

- time: O(n)
- space: O(n)

### Array Flip-over

可通过多次反转数组达到旋转数组的效果（不清楚背后的原理...）。

该题是向右旋转数组，步骤如下：

- 反转整个数组 (反转数组不需额外空间，时间复杂度为 O(n))
- 反转前 `k` 个元素 && 反转 `k` 到末尾的元素

如果是向左旋转数组，那么将上述步骤反过来即可。

- time: O(n)
- space: O(1)
