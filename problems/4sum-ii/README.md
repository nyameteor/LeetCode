# 454. 4Sum II

- Difficulty: Medium
- Topics: Array, Hash Table
- Link: https://leetcode.com/problems/4sum-ii/

## Description

Given four integer arrays `nums1`, `nums2`, `nums3`, and `nums4` all of length `n`, return the number of tuples `(i, j, k, l)` such that:

- `0 <= i, j, k, l < n`
- `nums1[i] + nums2[j] + nums3[k] + nums4[l] == 0`

**Example 1:**

```
Input: nums1 = [1,2], nums2 = [-2,-1], nums3 = [-1,2], nums4 = [0,2]
Output: 2
Explanation:
The two tuples are:
1. (0, 0, 0, 1) -> nums1[0] + nums2[0] + nums3[0] + nums4[1] = 1 + (-2) + (-1) + 2 = 0
2. (1, 1, 0, 0) -> nums1[1] + nums2[1] + nums3[0] + nums4[0] = 2 + (-1) + (-1) + 0 = 0
```

**Example 2:**

```
Input: nums1 = [0], nums2 = [0], nums3 = [0], nums4 = [0]
Output: 1
```

**Constraints:**

- `n == nums1.length`
- `n == nums2.length`
- `n == nums3.length`
- `n == nums4.length`
- `1 <= n <= 200`
- `-2^28 <= nums1[i], nums2[i], nums3[i], nums4[i] <= 2^28`

## Solution

### Key Idea

- The equation `a + b + c + d == 0` can be rearranged as `(a + b) == -(c + d)`.
- Instead of checking all 4-tuples directly, we split the arrays into two pairs:
  - Compute all possible sums of elements from `nums1` and `nums2`, and store their frequencies in a hash map.
  - Then, for each sum of elements from `nums3` and `nums4`, check how many times the negated sum appears in the map.
- This observation reduces the problem from **O(n^4)** to **O(n^2)** time complexity by using hashing.
