# 416. Partition Equal Subset Sum

- Difficulty: Medium
- Topics: Array, Dynamic Programming
- Link: https://leetcode.com/problems/partition-equal-subset-sum/

## Description

Given a **non-empty** array `nums` containing **only positive integers**, find if the array can be partitioned into two subsets such that the sum of elements in both subsets is equal.

**Example 1:**

```
Input: nums = [1,5,11,5]
Output: true
Explanation: The array can be partitioned as [1, 5, 5] and [11].
```

**Example 2:**

```
Input: nums = [1,2,3,5]
Output: false
Explanation: The array cannot be partitioned into equal sum subsets.
```

**Constraints:**

- `1 <= nums.length <= 200`
- `1 <= nums[i] <= 100`

## Solution

### Key Idea

This is a subset sum problem: determine if a subset of `nums` sums to `total / 2`. If the total sum is odd, it's impossible.

We can use dynamic programming to track possible subset sums:

- State: `dp[i][sum] = true` if a sum `sum` can be formed using the first `i` numbers, i.e., `nums[0]` to `nums[i-1]`. This indexing avoids handling the base case (`dp[0][0]` means using zero elements).
- Transition: From `dp[i-1][sum]`:
  - Don't take `nums[i-1]` -> `dp[i][sum] = true`
  - Take `nums[i-1]` -> `dp[i][sum + nums[i-1]] = true`

### References

- [[C++/Python] 5 Simple Solutions w/ Explanation | Optimization from Brute-Force to DP to Bitmask](https://leetcode.com/problems/partition-equal-subset-sum/solutions/1624939/c-python-5-simple-solutions-w-explanation-optimization-from-brute-force-to-dp-to-bitmask/)
