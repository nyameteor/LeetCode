# 53. Maximum Subarray

- Difficulty: Medium
- Topics: Array, Divide and Conquer, Dynamic Programming
- Link: https://leetcode.com/problems/maximum-subarray/

## Description

Given an integer array `nums`, find the contiguous subarray (containing at least one number) which has the largest sum and return _its sum_.

A **subarray** is a **contiguous** part of an array.

**Example 1:**

```
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: [4,-1,2,1] has the largest sum = 6.
```

**Example 2:**

```
Input: nums = [1]
Output: 1
```

**Example 3:**

```
Input: nums = [5,4,-1,7,8]
Output: 23
```

**Constraints:**

- `1 <= nums.length <= 10^5`
- `-10^4 <= nums[i] <= 10^4`

**Follow up:** If you have figured out the `O(n)` solution, try coding another solution using the **divide and conquer** approach, which is more subtle.

## Solution

### Dynamic Programming (DP)

- Idea: Track the maximum subarray sum ending at each index
- Definition:
  - `dp[i]` = max subarray sum ending at index `i`
- Transition:
  - `dp[i] = max(nums[i], nums[i] + dp[i-1])`
- Initialization:
  - `dp[0] = nums[0]`
  - `maxSum = dp[0]`
- Complexity:
  - Time: O(n)
  - Space: O(n)

---

### Kadane’s Algorithm (Optimized DP)

- Idea: Track the max subarray sum ending at the current index using a single variable
- Definition:
  - `curSum` = max subarray sum ending at index `i`
- Transition:
  - `curSum = max(nums[i], curSum + nums[i])`
  - `maxSum = max(maxSum, curSum)`
- Initialization:
  - `curSum = nums[0]`
  - `maxSum = nums[0]`
- Complexity:
  - Time: O(n)
  - Space: O(1)

### Divide and Conquer

- Recursively split the array into halves.
- The answer is the max of:
  1. Left half's max subarray
  2. Right half's max subarray
  3. Max subarray crossing the middle
- Combine results like in merge sort.
- Complexity:
  - Time: O(n log n)
  - Space: O(log n)

### References

- [[C++/Python] 7 Simple Solutions w/ Explanation | Brute-Force + DP + Kadane + Divide & Conquer](https://leetcode.com/problems/maximum-subarray/solutions/1595195/c-python-7-simple-solutions-w-explanation-brute-force-dp-kadane-divide-conquer/)
- [Maximum subarray problem](https://en.wikipedia.org/wiki/Maximum_subarray_problem)
