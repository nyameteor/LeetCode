# 136. Single Number

- Difficulty: Easy
- Topics: Array, Bit Manipulation
- Link: https://leetcode.com/problems/single-number/

## Description

Given a **non-empty** array of integers `nums`, every element appears _twice_ except for one. Find that single one.

You must implement a solution with a linear runtime complexity and use only constant extra space.

**Example 1:**

```
Input: nums = [2,2,1]
Output: 1
```

**Example 2:**

```
Input: nums = [4,1,2,1,2]
Output: 4
```

**Example 3:**

```
Input: nums = [1]
Output: 1
```

**Constraints:**

- `1 <= nums.length <= 3 * 10^4`
- `-3 * 10^4 <= nums[i] <= 3 * 10^4`
- Each element in the array appears twice except for one element which appears only once.

## Solution

### Count Frequency

遍历数组元素，若元素不存在于集合则插入，若元素再次出现则删除。遍历结束后，集合中留下的元素为结果。

### Bit Manipulation

由于 `A XOR A = 0`, `0 XOR A = A`，且 XOR 满足交换律，故对数组中的所有元素以任意顺序进行 XOR 运算，即为所求结果。

> XOR 满足交换律的证明：https://math.stackexchange.com/questions/293793/prove-xor-is-commutative-and-associative
