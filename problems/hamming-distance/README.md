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

- Compute `z = x ^ y` to get the differing bits.
- Count the number of set bits (`1`s) in `z` using Brian Kernighan’s Algorithm:
  - Repeatedly do `x &= x - 1` to clear the least significant set bit.
  - Increment the count each time.

### References

- [Counting bits set, Brian Kernighan's way](https://graphics.stanford.edu/~seander/bithacks.html#CountBitsSetNaive)
