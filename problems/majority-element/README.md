# 169. Majority Element

- Difficulty: Easy
- Topics: Array, Hash Table, Divide and Conquer, Sorting, Counting
- Link: https://leetcode.com/problems/majority-element/

## Description

Given an array `nums` of size `n`, return _the majority element_.

The majority element is the element that appears more than `⌊n / 2⌋` times. You may assume that the majority element always exists in the array.

**Example 1:**

```
Input: nums = [3,2,3]
Output: 3
```

**Example 2:**

```
Input: nums = [2,2,1,1,1,2,2]
Output: 2
```

**Constraints:**

- `n == nums.length`
- `1 <= n <= 5 * 10^4`
- `-2^31 <= nums[i] <= 2^31 - 1`

**Follow-up:** Could you solve the problem in linear time and in `O(1)` space?

## Solution

### HashMap

定义序列长度为 N，使用 HashMap 计数每个元素的出现次数，遍历 HashMap 找到出现次数大于 N/2 的元素 m ，即为结果。

C++ 中 HashMap 是不可遍历的，可以先转换为 Vector。

- Time complexity: O(n)
- Space complexity: O(n)

### Boyer-Moore Voting Algorithm

Refer: https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_majority_vote_algorithm

多数投票算法，是一种用来寻找一组元素中占多数元素的算法，也是处理数据流的一种典型算法。

伪代码表示如下：

```shell
- 定义候选元素 m，计数器 i = 0
- 对输入序列的每一个元素 x：
  - 若 i = 0，则 m = x
  - 若 m = x，则 i = i + 1；否则，i = i - 1
- 返回 m
```

- Time complexity: O(n)
- Space complexity: O(1)
