# 923. 3Sum With Multiplicity

- Difficulty: Medium
- Topics: Array, Hash Table, Two Pointers, Sorting, Counting
- Link: https://leetcode.com/problems/3sum-with-multiplicity/

## Description

Given an integer array `arr`, and an integer `target`, return the number of tuples `i, j, k` such that `i < j < k` and `arr[i] + arr[j] + arr[k] == target`.

As the answer can be very large, return it **modulo** `109 + 7`.

**Example 1:**

```
Input: arr = [1,1,2,2,3,3,4,4,5,5], target = 8
Output: 20
Explanation:
Enumerating by the values (arr[i], arr[j], arr[k]):
(1, 2, 5) occurs 8 times;
(1, 3, 4) occurs 8 times;
(2, 2, 4) occurs 2 times;
(2, 3, 3) occurs 2 times.
```

**Example 2:**

```
Input: arr = [1,1,2,2,2,2], target = 5
Output: 12
Explanation:
arr[i] = 1, arr[j] = arr[k] = 2 occurs 12 times:
We choose one 1 from [1,1] in 2 ways,
and two 2s from [2,2,2,2] in 6 ways.
```

**Constraints:**

- `3 <= arr.length <= 3000`
- `0 <= arr[i] <= 100`
- `0 <= target <= 300`

## Solution

Refer: https://leetcode.com/problems/3sum-with-multiplicity/solution/

### Three Pointer

1. Sort array `A`.
2. Set `T = target - A[i]`, then we can use two-pointer to find `A[j] + A[k] == T`.
   - initial `j = i + 1`
   - initial `k = A.size() - 1`
3. Denote multiplicity of `x` as `M(x)`. Whenever `A[j] + A[k] == T`:
   - if `A[j] != A[k]`, the number of pairs is `M(A[j]) * M(A[k])`.
   - if `A[j] == A[k]`, the number of pairs is `M * (M - 1) / 2`, where `M = M(A[j]) = k - j + 1`.
