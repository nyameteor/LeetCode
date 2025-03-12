# 29. Divide Two Integers

- Difficulty: Medium
- Topics: Math, Bit Manipulation
- Link: https://leetcode.com/problems/divide-two-integers/

## Description

Given two integers `dividend` and `divisor`, divide two integers without using multiplication, division, and mod operator.

Return the quotient after dividing `dividend` by `divisor`.

The integer division should truncate toward zero, which means losing its fractional part. For example, `truncate(8.345) = 8` and `truncate(-2.7335) = -2`.

**Note:** Assume we are dealing with an environment that could only store integers within the **32-bit** signed integer range: `[−2^31, 2^31 − 1]`. For this problem, assume that your function **returns** `2^31 − 1` **when the division result overflows**.

**Example 1:**

```
Input: dividend = 10, divisor = 3
Output: 3
Explanation: 10/3 = truncate(3.33333..) = 3.
```

**Example 2:**

```
Input: dividend = 7, divisor = -3
Output: -2
Explanation: 7/-3 = truncate(-2.33333..) = -2.
```

**Example 3:**

```
Input: dividend = 0, divisor = 1
Output: 0
```

**Example 4:**

```
Input: dividend = 1, divisor = 1
Output: 1
```

**Constraints:**

- `-2^31 <= dividend, divisor <= 2^31 - 1`
- `divisor != 0`

## Solution

To solve the problem of dividing two integers without multiplication, division, or modulus operators:

1. **Overflow Handling**: The only overflow case is dividing `math.MinInt32` by `-1`, which exceeds the 32-bit range. This is handled upfront by returning `math.MaxInt32`.
2. **Sign Determination**: Check if the result should be negative by comparing the signs of the dividend and divisor.
3. **Absolute Values**: Work with the absolute values of the dividend and divisor to simplify the division logic.
4. **Bit Manipulation**: Use bit shifting to perform division efficiently. Shift the divisor left (multiply by powers of 2) until it exceeds the dividend, then subtract and adjust the quotient.
5. **Return Result**: Adjust the sign of the quotient and return the result.
