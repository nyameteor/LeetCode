# 166. Fraction to Recurring Decimal

- Difficulty: Medium
- Topics: Hash Table, Math, String
- Link: https://leetcode.com/problems/fraction-to-recurring-decimal/

## Description

Given two integers representing the `numerator` and `denominator` of a fraction, return *the fraction in string format*.

If the fractional part is repeating, enclose the repeating part in parentheses.

If multiple answers are possible, return **any of them**.

It is **guaranteed** that the length of the answer string is less than `10^4` for all the given inputs.

**Example 1:**

```
Input: numerator = 1, denominator = 2
Output: "0.5"
```

**Example 2:**

```
Input: numerator = 2, denominator = 1
Output: "2"
```

**Example 3:**

```
Input: numerator = 4, denominator = 333
Output: "0.(012)"
```

**Constraints:**

- `-2^31 <= numerator, denominator <= 2^31 - 1`
- `denominator != 0`

## Solution

### Approach: Long Division

- Integer Part: Perform standard division for the integer part.
- Fractional Part: Use long division to generate digits after the decimal.
- Cycle Detection:
  - Track remainders and their positions with a hash map.
  - If a remainder repeats, insert parentheses around the repeating part.
- Edge Cases:
  - Handle sign for negative results.
  - If remainder becomes 0, there's no cycle.
