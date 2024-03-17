# 15. 3Sum

- Difficulty: Medium
- Topics: Array, Two Pointers, Sorting
- Link: https://leetcode.com/problems/3sum/

## Description

Given an integer array nums, return all the triplets `[nums[i], nums[j], nums[k]]` such that `i != j`, `i != k`, and `j != k`, and `nums[i] + nums[j] + nums[k] == 0`.

Notice that the solution set must not contain duplicate triplets.

**Example 1:**

```
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
```

**Example 2:**

```
Input: nums = []
Output: []
```

**Example 3:**

```
Input: nums = [0]
Output: []
```

**Constraints:**

- `0 <= nums.length <= 3000`
- `-10^5 <= nums[i] <= 10^5`

## Solution

We can use the two pointers approach.

First, sort the array in ascending order. Iterate the array using pointer `i`. Let `l` = `i+1`, `r` = `size(nums) - 1`.
Let `sum = nums[i] + nums[l] + nums[r]`:

- If `sum == 0`, add tuple to the result, `l++` and `r--`, until `l >= r`.
- If `sum < 0`, `l++`, until `l >= r`.
- If `sum > 0`, `r--`, until `l >= r`.

And when iterating the `nums` with `i`, `l` and `r`, we need to remove duplicate elements, which will prevent adding duplicate tuples and prevent timeouts.

Example:

```text
sorted_nums         operation
-4 -1 -1  0  1  2
 i  l           r   -4 + -1 + 2 = -3 <  0, l -> nums[l] < nums[l+1], l++
 i        l     r   -4 +  0 + 2 = -2 <  0, l -> nums[l] < nums[l+1], l++
 i           l  r   -4 +  1 + 2 = -1 <  0, l -> nums[l] < nums[l+1], l++
                                           i -> nums[i] < nums[i+1], i++
    i  l        r   -1 + -1 + 2 =  0 == 0, l -> nums[l] < nums[l+1], l++; r -> nums[r] > nums[r-1], r++; add to res
    i     l     r   -1 +  0 + 2 =  1 >  0, r -> nums[r] > nums[r-1], r++
    i     l  r      -1 +  0 + 1 =  0 == 0, l -> nums[l] < nums[l+1], l++; r -> nums[r] > nums[r-1], r++; add to res
                                           i -> nums[i] < nums[i+1], i++
          i  l  r   0 +   1 + 2 =  3 >  0, l -> nums[r] > nums[r-1], l++
```

- time: O(n^2)
- space: O(n^2)

References:

- https://leetcode.com/problems/3sum/solutions/7402/share-my-ac-c-solution-around-50ms-o-n-n-with-explanation-and-comments/
