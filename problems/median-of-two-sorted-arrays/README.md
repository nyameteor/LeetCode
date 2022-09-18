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

寻找 A1, A2 两个已排序数组的中位数 M：

- 对 A1 和 A2 进行切分，设 `N1 = size(A)`，`N2 = size(A2)`：
  - 设切点 C1 将 A1 切分为 [...L1 | R1...]，切点 C2 将 A2 切分为 [...L2 | R2...]；
  - C1、C2 满足 `C1 + C2 = floor((N1 + N2) / 2)`，即 [...L1] + [...L2] 和 [R1...] + [R2...] 具有相同数量的元素。
- 使用二分搜索法，在 A1 上搜索目标的切点 C1：
  - 初始化：
    - 设 `lo = 0`, `hi = N1`；
    - 设 `C1 = (lo + hi) / 2`，且 `C2 = floor((N1 + N2) / 2) - C1`。
  - 搜索：
    - 若 `L1 > R2`，说明 C1 过大，故令 `hi = C1 - 1`；
    - 若 `L2 > R1`，说明 C1 过小，故令 `lo = C1 + 1`；
    - 若 `L1 <= R2 && L2 <= R1`，则 C1 为目标切分点。由于 C1 和 C2 的和为常数，此时 C2 亦确定。
- 计算中位数 M：
  - 若 N1 + N2 为奇数，则 `M = min(R1, R2)`；
  - 若 N1 + N2 为偶数，则 `M = (max(L1, L2) + min(R1, R2)) / 2`。

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

Refer: stellari@leetcode https://leetcode.com/problems/median-of-two-sorted-arrays/discuss/2471
