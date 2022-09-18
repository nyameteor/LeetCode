# 461. Hamming Distance

- Difficulty: Easy
- Topics: Bit Manipulation
- Link: https://leetcode.com/problems/hamming-distance/

## Description

The [Hamming distance](https://en.wikipedia.org/wiki/Hamming_distance) between two integers is the number of positions at which the corresponding bits are different.

Given two integers `x` and `y`, return _the **Hamming distance** between them_.

**Example 1:**

```
Input: x = 1, y = 4
Output: 2
Explanation:
1   (0 0 0 1)
4   (0 1 0 0)
       ↑   ↑
The above arrows point to positions where the corresponding bits are different.
```

**Example 2:**

```
Input: x = 3, y = 1
Output: 1
```

**Constraints:**

- `0 <= x, y <= 2^31 - 1`

## Solution

### XOR + Count bits

对 x 和 y 按位异或得到 z，计算 z 中为 1 的位数（可以逐位和 1 执行按位与），即为结果。

更多解法参考：[@leetcode/archit91](https://leetcode.com/problems/hamming-distance/discuss/1585474/C%2B%2BPython-4-Simple-Solutions-w-Explanations-or-XOR-and-Brian-Kernighan-method)

Refer:

- [Bitwise operations in C](https://en.wikipedia.org/wiki/Bitwise_operations_in_C)
- [Operators in C and C++](https://en.wikipedia.org/wiki/Operators_in_C_and_C%2B%2B)
