# 167. Two Sum II - Input Array Is Sorted

- Difficulty: Easy
- Topics: Array, Two Pointers, Binary Search
- Link: https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/

## Description

Given a **1-indexed** array of integers `numbers` that is already **\*sorted in non-decreasing order\***, find two numbers such that they add up to a specific `target` number. Let these two numbers be `numbers[index1]` and `numbers[index2]` where `1 <= index1 < index2 <= numbers.length`.

Return _the indices of the two numbers,_ `index1` _and_ `index2`_, **added by one** as an integer array_ `[index1, index2]` _of length 2._

The tests are generated such that there is **exactly one solution**. You **may not** use the same element twice.

**Example 1:**

```
Input: numbers = [2,7,11,15], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore, index1 = 1, index2 = 2. We return [1, 2].
```

**Example 2:**

```
Input: numbers = [2,3,4], target = 6
Output: [1,3]
Explanation: The sum of 2 and 4 is 6. Therefore index1 = 1, index2 = 3. We return [1, 3].
```

**Example 3:**

```
Input: numbers = [-1,0], target = -1
Output: [1,2]
Explanation: The sum of -1 and 0 is -1. Therefore index1 = 1, index2 = 2. We return [1, 2].
```

**Constraints:**

- `2 <= numbers.length <= 3 * 10^4`
- `-1000 <= numbers[i] <= 1000`
- `numbers` is sorted in **non-decreasing order**.
- `-1000 <= target <= 1000`
- The tests are generated such that there is **exactly one solution**.

## Solution

### Two Pointers

由于数组已经排序，可以设置双指针 l 和 r，l 在数组的起始位置，r 在数组的终止位置，且遍历过程中满足 l < r：

- A[l] + A[r] < target --> r--
- A[l] + A[r] > target --> l--
- A[l] + A[r] == target --> done
