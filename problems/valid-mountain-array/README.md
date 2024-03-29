# 941. Valid Mountain Array

- Difficulty: Easy
- Topics: Array
- Link: https://leetcode.com/problems/valid-mountain-array/

## Description

Given an array of integers `arr`, return _`true` if and only if it is a valid mountain array_.

Recall that arr is a mountain array if and only if:

- `arr.length >= 3`

- There exists some `i` with `0 < i < arr.length - 1` such that:

  - `arr[0] < arr[1] < ... < arr[i - 1] < arr[i] `
  - `arr[i] > arr[i + 1] > ... > arr[arr.length - 1]`

![img](https://assets.leetcode.com/uploads/2019/10/20/hint_valid_mountain_array.png)

**Example 1:**

```
Input: arr = [2,1]
Output: false
```

**Example 2:**

```
Input: arr = [3,5,5]
Output: false
```

**Example 3:**

```
Input: arr = [0,3,2,1]
Output: true
```

**Constraints:**

- `1 <= arr.length <= 10^4`
- `0 <= arr[i] <= 10^4`

## Solution

### One Pass

To valid if an array is a mountain array, we move strictly up, then strictly down.

Use `N` to represent size of array `A`. Traversal the array:

- First, while `arr[i] < arr[i + 1]`, do `i++`, it represent move strictly up. When the loop ends:
  - if `i == 0`, it means peak is the start, so `A` is not a mountain array.
  - if `i == N`, it means peak is the end, so `A` is not a mountain array.
  - else, we check the reamaining elements:
    while `arr[i] > arr[i + 1]`, do `i++`, it represent move strictly down. When the loop ends:
    - if `i != N - 1`, it means there are more than one peak, so `A` is not a mountain array.
    - else, it means `A` is a mountain array.

Refer: https://leetcode.com/problems/valid-mountain-array/solution/
