# 4. Median of Two Sorted Arrays

- Difficulty: Hard
- Topics: Array, Binary Search, Divide and Conquer
- Link: https://leetcode.com/problems/median-of-two-sorted-arrays/

## Description

Given two sorted arrays `nums1` and `nums2` of size `m` and `n` respectively, return **the median** of the two sorted arrays.

The overall run time complexity should be `O(log (m+n))`.

**Example 1:**

```
Input: nums1 = [1,3], nums2 = [2]
Output: 2.00000
Explanation: merged array = [1,2,3] and median is 2.
```

**Example 2:**

```
Input: nums1 = [1,2], nums2 = [3,4]
Output: 2.50000
Explanation: merged array = [1,2,3,4] and median is (2 + 3) / 2 = 2.5.
```

**Constraints:**

- `nums1.length == m`
- `nums2.length == n`
- `0 <= m <= 1000`
- `0 <= n <= 1000`
- `1 <= m + n <= 2000`
- `-10^6 <= nums1[i], nums2[i] <= 10^6`

## Solution

### Binary Search

We want to find the median `M` of two sorted arrays `A1`, `A2`. Let `N1 = size(A)`, `N2 = size(A2)`. Denote the cut point `C1` cut `A1` into `[...L1, R1...]`, `C2` cut `A2` into `[...L2, R2...]`; `C1`, `C2` will satisfy `C1 + C2 = (N1 + N2) / 2`, which means `[...L1] + [...L2]` and `[R1...] + [R2...]` have the same sizes.

We can use the binary search method to find the target `C1` in `A1`:

- Initialization:
  - Let `low = 0`, `high = N1`.
  - Let `C1 = (low + high) / 2`, `C2 = (N1 + N2) / 2 - C1`.
- Search:
  - If `L1 > R2`, then `C1` is too large, so let `high = C1 - 1`;
  - If `L2 > R1`, then `C1` is too small, so let `low = C1 + 1`;
  - If `L1 <= R2 && L2 <= R1`, then `C1` is the target.

Since we know `C1 + C2`, we also know `C2`.

Then we can calculate the median `M`:

- If `N1 + N2` is odd, then `M = min(R1, R2)`;
- If `N1 + N2` is even, then `M = (max(L1, L2) + min(R1, R2)) / 2`.

Example:

```shell
A1      1 2 3 4             N1: 4
          ^ ^
         L1 R1
index   0 1 2 3             C1: C1 + C2 = C
            ^
            C1

A2      1 2 3 4 5           N2: 5
          ^ ^
         L2 R2
index   0 1 2 3 4           C2: C2 + C1 = C
            ^
            C2

A       1 1 2 2 3 3 4 4 5   N: N1 + N2 = 9
              ^ ^
              L R
index   0 1 2 3 4 5 6 7 8   C: floor((N1 + N2) / 2) = 4
                ^
                C

M = R = min(R1, R2) = min(3, 3) = 3.

```

Reference:

- https://leetcode.com/problems/median-of-two-sorted-arrays/solutions/2471/very-concise-o-log-min-m-n-iterative-solution-with-detailed-explanation/
- https://leetcode.com/problems/median-of-two-sorted-arrays/editorial/
