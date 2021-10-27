# 75. Sort Colors

- Difficulty: Medium
- Topics: Array, Two Pointers, Sorting
- Link: https://leetcode.com/problems/sort-colors/

## Description

Given an array `nums` with `n` objects colored red, white, or blue, sort them **[in-place](https://en.wikipedia.org/wiki/In-place_algorithm)** so that objects of the same color are adjacent, with the colors in the order red, white, and blue.

We will use the integers `0`, `1`, and `2` to represent the color red, white, and blue, respectively.

You must solve this problem without using the library's sort function.

**Example 1:**

```
Input: nums = [2,0,2,1,1,0]
Output: [0,0,1,1,2,2]
```

**Example 2:**

```
Input: nums = [2,0,1]
Output: [0,1,2]
```

**Example 3:**

```
Input: nums = [0]
Output: [0]
```

**Example 4:**

```
Input: nums = [1]
Output: [1]
```

**Constraints:**

- `n == nums.length`
- `1 <= n <= 300`
- `nums[i]` is `0`, `1`, or `2`.

**Follow up:** Could you come up with a one-pass algorithm using only constant extra space?

## Solution

本题是经典的 **Dutch national flag problem**，Wikipedia 上有[详细的介绍](https://en.wikipedia.org/wiki/Dutch_national_flag_problem)。

参考 Wikipedia 的介绍，本题的解决方案对设计排序算法很有意义，之后可以尝试补充它们的关联。

init i, j := 0, k := size of A - 1;

while j<=k, check A[j]:

- A[j] < mid -> swap(A[i], A[j]), i++, j++
- A[j] > mid -> swap(A[j], A[k]), k--
- A[j] = mid -> j++

```shell
 2 0 1 2 0 1 1      i,j := 0, k := size of A - 1
 ^           ^
i,j          k


 2 0 1 2 0 1 1      A[j] > mid -> swap(A[j], A[k]), k--
 ^           ^
i,j          k

 1 0 1 2 0 1 2      A[j] = mid -> j++
 ^         ^
i,j        k

 1 0 1 2 0 1 2      A[j] < mid -> swap(A[i], A[j]), i++, j++
 ^ ^       ^
 i j       k

 0 1 1 2 0 1 2      A[j] = mid -> j++
   ^ ^     ^
   i j     k

 0 1 1 2 0 1 2      A[j] > mid -> swap(A[j], A[k]), k--
   ^   ^   ^
   i   j   k

 0 1 1 1 0 2 2      A[j] = mid -> j++
   ^   ^ ^
   i   j k

 0 1 1 1 0 2 2      A[j] < mid -> swap(A[i], A[j]), i++, j++
   ^     ^
   i    j,k

 0 0 1 1 1 2 2      j > k, return
     ^   ^ ^
     i   k j
```
