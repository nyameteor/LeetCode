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

### Hash Table

Idea: Count number of solution `a + b + c + d = 0 -> a + b = -(c + d)`.

- Iterate through the first 2 arrays and count the frequency of all possible sums of pairs: `map(a + b, frequency)`.
- Iterate through the other 2 arrays, when `-(c + d)` is a valid key in map, means `a + b = -(c + d)`, and there are `map[-(c + d)]` matches in this case, so `matches += map[-(c + d)]`.
- The count of `matches` is the required result.

Time: O(N^2)
