# 43. Multiply Strings

- Difficulty: Medium
- Topics: Math, String, Simulation
- Link: https://leetcode.com/problems/multiply-strings/

## Description

Given two non-negative integers `num1` and `num2` represented as strings, return the product of `num1` and `num2`, also represented as a string.

**Note:** You must not use any built-in BigInteger library or convert the inputs to integer directly.

**Example 1:**

```
Input: num1 = "2", num2 = "3"
Output: "6"
```

**Example 2:**

```
Input: num1 = "123", num2 = "456"
Output: "56088"
```

**Constraints:**

- `1 <= num1.length, num2.length <= 200`
- `num1` and `num2` consist of digits only.
- Both `num1` and `num2` do not contain any leading zero, except the number `0` itself.

## Solution

This problem simulates manual string-based multiplication. There are two main approaches:

- **Digit-by-digit multiplication**: Multiply each digit of `num2` with `num1`, shift by position, and sum all partial results.
- **Positional accumulation**: Multiply every digit pair and accumulate results by position, handling carry inline. This method is more efficient and concise.

### References

- [Editorial](https://leetcode.com/problems/multiply-strings/editorial/)
- [C++ Solution](https://github.com/haoel/leetcode/blob/master/algorithms/cpp/multiplyStrings/multiplyStrings.cpp)
- [Golang Solution](https://leetcode.com/problems/multiply-strings/solutions/529103/golang-solution/)
