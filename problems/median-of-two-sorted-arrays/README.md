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

### Binary Search (Recursive)

- Find the **k-th smallest element** in the union of two arrays recursively:
  - Discard `k/2` elements from the array with the smaller `k/2`-th element
  - Base cases:
    - If one array is empty: return `k`-th from the other
    - If `k == 1`: return min of both heads

- Median is:
  - Odd total: `findKth(k)`
  - Even total: average of `findKth(k)` and `findKth(k+1)`

### Binary Search (Iterative)

- Perform binary search on the **shorter array** to find the correct partition:
  - Partition `nums1` at `i`, and `nums2` at `j = (m+n+1)/2 - i`
  - Invariant: `max(left1, left2) <= min(right1, right2)`

- If valid partition:
  - Odd total: return `max(left1, left2)`
  - Even total: return average of `max(left1, left2)` and `min(right1, right2)`

- Use `math.MinInt32` and `math.MaxInt32` to handle empty partitions.

### Reference

- https://leetcode.com/problems/median-of-two-sorted-arrays/editorial/
- https://leetcode.com/problems/median-of-two-sorted-arrays/solutions/2471/very-concise-o-log-min-m-n-iterative-solution-with-detailed-explanation/
