# 2523. Closest Prime Numbers in Range

- Difficulty: Medium
- Topics: Math, Number Theory
- Link: https://leetcode.com/problems/closest-prime-numbers-in-range/

## Description

Given two positive integers `left` and `right`, find the two integers `num1` and `num2` such that:

- `left <= num1 < num2 <= right` .
- Both `num1` and `num2` are

- .
- `num2 - num1` is the **minimum** amongst all other pairs satisfying the above conditions.

Return the positive integer array `ans = [num1, num2]`. If there are multiple pairs satisfying these conditions, return the one with the **smallest** `num1` value. If no such numbers exist, return `[-1, -1]`*.*

**Example 1:**

```
Input: left = 10, right = 19
Output: [11,13]
Explanation: The prime numbers between 10 and 19 are 11, 13, 17, and 19.
The closest gap between any pair is 2, which can be achieved by [11,13] or [17,19].
Since 11 is smaller than 17, we return the first pair.

```

**Example 2:**

```
Input: left = 4, right = 6
Output: [-1,-1]
Explanation: There exists only one prime number in the given range, so the conditions cannot be satisfied.

```

**Constraints:**

- `1 <= left <= right <= 10^6`

## Solution

We use the **Sieve of Eratosthenes** to efficiently find all prime numbers up to `right`. Then, we iterate through the range `[left, right]` to find the closest pair of prime numbers.

### **Key Steps**

1. **Precompute primes** using the **Sieve of Eratosthenes** in `O(n log log n)`.
2. **Find the closest prime pair** by iterating through `[left, right]`:
   - Keep track of the previous prime.
   - Update the result when a smaller gap is found.

### **Edge Cases**

- If there are **fewer than two primes** in the range, return `[-1, -1]`.
- The sieve ensures **efficient prime checking**, avoiding unnecessary computations.
