# 31. Next Permutation

- Difficulty: Medium
- Topics: Array, Two Pointers
- Link: https://leetcode.com/problems/next-permutation/

## Description

Implement **next permutation**, which rearranges numbers into the lexicographically next greater permutation of numbers.

If such an arrangement is not possible, it must rearrange it as the lowest possible order (i.e., sorted in ascending order).

The replacement must be **[in place](http://en.wikipedia.org/wiki/In-place_algorithm)** and use only constant extra memory.

**Example 1:**

```
Input: nums = [1,2,3]
Output: [1,3,2]
```

**Example 2:**

```
Input: nums = [3,2,1]
Output: [1,2,3]
```

**Example 3:**

```
Input: nums = [1,1,5]
Output: [1,5,1]
```

**Example 4:**

```
Input: nums = [1]
Output: [1]
```

**Constraints:**

- `1 <= nums.length <= 100`
- `0 <= nums[i] <= 100`

## Solution

Accorading to [Wikipedia](https://en.wikipedia.org/wiki/Permutation#Generation_in_lexicographic_order):

The method goes back to [Narayana Pandita](https://en.wikipedia.org/wiki/Narayana_Pandita_(mathematician)) in 14th century India, and has been rediscovered frequently.

The following algorithm generates the next permutation lexicographically after a given permutation. It changes the given permutation in-place.

1. Find the largest index `k` such that `a[k] < a[k + 1]`. If no such index exists, the permutation is the last permutation.
2. Find the largest index `l` greater than `k` such that `a[k] < a[l]`.
3. Swap the value of `a[k]` with that of `a[l]`.
4. Reverse the sequence from `a[k + 1]` up to and including the final element `a[n]`.
