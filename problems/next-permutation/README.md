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

Take a look the following continous permutation, can you find the pattern?

```shell
1 2 3 4
1 2 4 3
1 3 2 4
1 3 4 2
1 4 2 3
2 1 3 4
...
```

The pattern can be descripted as below:

```shell
1) from n-1 to 0, find the first place [i-1] which num[i-1] < num[i]
2) from n-1 to i, find the first place [j] which num[j] > num[i-1]
3) swap the num[j] with num[i-1]
4) sort the sub-array [i, n)
```

For example:

```shell
1 4 3 2     <--1) find the first place [i-1] which num[i-1] < num[i]
  i

1 4 3 2     <--2) find the first place [j] which num[j] > num[i-1]
i-1   j

2 4 3 1     <--3) swap num[j] with num[i-1]
i-1   j

2 1 3 4     <--4) sort the sub-array [i, n]
  i   n

# Corner case
4 3 2 1, the next permutation is 1 2 3 4
```

Ref: [haoel/leetcode/nextPermutation.cpp](https://github.com/haoel/leetcode/blob/master/algorithms/cpp/nextPermutation/nextPermutation.cpp)
