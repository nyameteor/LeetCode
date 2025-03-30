# 300. Longest Increasing Subsequence

- Difficulty: Medium
- Topics: Array, Binary Search, Dynamic Programming
- Link: https://leetcode.com/problems/longest-increasing-subsequence/

## Description

Given an integer array `nums`, return the length of the longest strictly increasing subsequence.

A **subsequence** is a sequence that can be derived from an array by deleting some or no elements without changing the order of the remaining elements. For example, `[3,6,2,7]` is a subsequence of the array `[0,3,1,6,2,2,7]`.

**Example 1:**

```
Input: nums = [10,9,2,5,3,7,101,18]
Output: 4
Explanation: The longest increasing subsequence is [2,3,7,101], therefore the length is 4.
```

**Example 2:**

```
Input: nums = [0,1,0,3,2,3]
Output: 4
```

**Example 3:**

```
Input: nums = [7,7,7,7,7,7,7]
Output: 1
```

**Constraints:**

- `1 <= nums.length <= 2500`
- `-10^4 <= nums[i] <= 10^4`

**Follow up:** Can you come up with an algorithm that runs in `O(n log(n))` time complexity?

## Solution

The problem requires finding the **longest strictly increasing subsequence** (LIS) in an array. A **subsequence** allows skipping elements but must maintain the original order.

### Approach: Dynamic Programming

#### Key Observations

1. **Each element contributes to an increasing subsequence**
   - Every number in `nums` can either **start a new subsequence** or **extend an existing one**.

2. **The LIS ending at each index depends on previous elements**
   - If `nums[i]` is greater than `nums[j]` (`j < i`), then `nums[i]` can extend the subsequence ending at `j`.

3. **Recursive Structure**
   - Define `dp[i]` as the **length of the LIS ending at index `i`**.
   - To compute `dp[i]`, check **all previous indices `j`** where `nums[j] < nums[i]`, then take the **longest** valid subsequence and extend it:

     ```
     dp[i] = max(dp[i], dp[j] + 1)
     ```

   - The overall LIS is the maximum value in `dp`.

#### Complexity

- Time: `O(n^2)`
- Space: `O(n)`

### Approach: Binary Search

Todo

### References

[[C++/Python] DP, Binary Search, BIT, Segment Tree Solutions - Picture explain - O(NlogN)](https://leetcode.com/problems/longest-increasing-subsequence/solutions/1326308/c-python-dp-binary-search-bit-segment-tree-solutions-picture-explain-o-nlogn)
