# 560. Subarray Sum Equals K

- Difficulty: Medium
- Topics: Array, Hash Table, Prefix Sum
- Link: https://leetcode.com/problems/subarray-sum-equals-k/

## Description

Given an array of integers `nums` and an integer `k`, return *the total number of subarrays whose sum equals to* `k`.

A subarray is a contiguous **non-empty** sequence of elements within an array.

**Example 1:**

```
Input: nums = [1,1,1], k = 2
Output: 2
```

**Example 2:**

```
Input: nums = [1,2,3], k = 3
Output: 2
```

**Constraints:**

- `1 <= nums.length <= 2 * 104`
- `-1000 <= nums[i] <= 1000`
- `-10^7 <= k <= 10^7`

## Solution

### Approach: Prefix Sum + Hash Map

#### Intuition

- Maintain a running sum `sum` of elements up to index `i`.
- Use a hashmap `sumFreq` to track occurrences of each `sum`.
- If `sum - k` exists in `sumFreq`, it means there are subarrays ending at `i` that sum to `k`, so add their count to the result.

#### Key Observations

- Initialize `sumFreq[0] = 1` to handle subarrays starting at index `0`.
- This avoids the **O(n^2)** brute-force approach, reducing time complexity to **O(n)**.

#### Complexity

- **Time:** `O(n)`, since we iterate through `nums` once.
- **Space:** `O(n)`, for storing prefix sums in the hashmap.
