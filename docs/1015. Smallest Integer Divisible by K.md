# 1015. Smallest Integer Divisible by K

- Difficulty: Medium
- Topics: Hash Table, Math
- Link: https://leetcode.com/problems/smallest-integer-divisible-by-k/

## Description

Given a positive integer `k`, you need to find the **length** of the **smallest** positive integer `n` such that `n` is divisible by `k`, and `n` only contains the digit `1`.

Return _the **length** of_ `n`. If there is no such `n`, return -1.

**Note:** `n` may not fit in a 64-bit signed integer.

**Example 1:**

```
Input: k = 1
Output: 1
Explanation: The smallest answer is n = 1, which has length 1.
```

**Example 2:**

```
Input: k = 2
Output: -1
Explanation: There is no such positive integer n divisible by 2.
```

**Example 3:**

```
Input: k = 3
Output: 3
Explanation: The smallest answer is n = 111, which has length 3.
```

**Constraints:**

- `1 <= k <= 10^5`

## Solution

### Brute Force

As the hint said, `111 = 11 * 10 + 1`, we only need to store remainders modulo `k` (the number `n` is easy to overflow).

Then the problem can be divided into two cases:

`n` exists:

We find the `n` that makes `n % k == 0`.

```shell
1 % 7
1
11 % 7
4
41 % 7
6
61 % 7
5
51 % 7
2
21 % 7
0
```

no such `n` exists:

The same remainder is sometimes repeated again, so the search will go in an endless loop.

```shell
1 % 6
1
11 % 6
5
51 % 6
3
31 % 6
1
11 % 6  # endless loop...
5
```

Refer: https://leetcode.com/problems/smallest-integer-divisible-by-k/discuss/1656121
