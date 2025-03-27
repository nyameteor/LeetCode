# 740. Delete and Earn

- Difficulty: Medium
- Topics: Array, Hash Table, Dynamic Programming
- Link: https://leetcode.com/problems/delete-and-earn/

## Description

You are given an integer array `nums`. You want to maximize the number of points you get by performing the following operation any number of times:

- Pick any `nums[i]` and delete it to earn `nums[i]` points. Afterwards, you must delete **every** element equal to `nums[i] - 1` and **every** element equal to `nums[i] + 1`.

Return *the **maximum number of points** you can earn by applying the above operation some number of times*.

**Example 1:**

```
Input: nums = [3,4,2]
Output: 6
Explanation: You can perform the following operations:
- Delete 4 to earn 4 points. Consequently, 3 is also deleted. nums = [2].
- Delete 2 to earn 2 points. nums = [].
You earn a total of 6 points.
```

**Example 2:**

```
Input: nums = [2,2,3,3,3,4]
Output: 9
Explanation: You can perform the following operations:
- Delete a 3 to earn 3 points. All 2's and 4's are also deleted. nums = [3,3].
- Delete a 3 again to earn 3 points. nums = [3].
- Delete a 3 once more to earn 3 points. nums = [].
You earn a total of 9 points.
```

**Constraints:**

- `1 <= nums.length <= 2 * 10^4`
- `1 <= nums[i] <= 10^4`

## Solution

### Approach: Dynamic Programming

- **Frequency Count:**  
   Calculate how many times each number appears in `nums` using a frequency map (`freq`).

- **Dynamic Programming:**  
   Use dynamic programming to determine the maximum points. Define `dp[i]` as the maximum points we can earn from the first `i` unique numbers:
  - Skip the current number (`dp[i-1]`)
  - Take the current number and its frequency, but skip its adjacent number (`dp[i-2] + earn[i]`)

- **Memoization:**  
   Memoize results in the `memo` array to avoid redundant calculations.

- **Base Case:**  
   If no elements are left, return 0.

#### Complexity

- **Time:** `O(n log n)` (due to sorting unique numbers)
- **Space:** `O(n)` (for memoization and storing unique numbers)
