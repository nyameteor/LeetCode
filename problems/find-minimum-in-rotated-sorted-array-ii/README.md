# 154. Find Minimum in Rotated Sorted Array II

- Difficulty: Hard
- Topics: Array, Binary Search
- Link: https://leetcode.com/problems/find-minimum-in-rotated-sorted-array-ii/

## Description

Suppose an array of length `n` sorted in ascending order is **rotated** between `1` and `n` times. For example, the array `nums = [0,1,4,4,5,6,7]` might become:

- `[4,5,6,7,0,1,4]` if it was rotated `4` times.
- `[0,1,4,4,5,6,7]` if it was rotated `7` times.

Notice that **rotating** an array `[a[0], a[1], a[2], ..., a[n-1]]` 1 time results in the array `[a[n-1], a[0], a[1], a[2], ..., a[n-2]]`.

Given the sorted rotated array `nums` that may contain **duplicates**, return _the minimum element of this array_.

You must decrease the overall operation steps as much as possible.

**Example 1:**

```
Input: nums = [1,3,5]
Output: 1
```

**Example 2:**

```
Input: nums = [2,2,2,0,1]
Output: 0
```

**Constraints:**

- `n == nums.length`
- `1 <= n <= 5000`
- `-5000 <= nums[i] <= 5000`
- `nums` is sorted and rotated between `1` and `n` times.

**Follow up:** This problem is similar to [Find Minimum in Rotated Sorted Array](https://leetcode.com/problems/find-minimum-in-rotated-sorted-array/description/), but `nums` may contain **duplicates**. Would this affect the runtime complexity? How and why?

## Solution

### Binary Search

Binary Search 的写法根据不同需求有细微的不同，下面是 Wikipedia 中当存在重复元素时，找到**最左边的解**的写法。这是一个常用的写法，需要牢记，后面的题解也用这个作为板子。

Refer: https://en.wikipedia.org/wiki/Binary_search_algorithm#Duplicate_elements

```shell
# Procedure for finding the leftmost element
function binary_search_leftmost(A, n, T):
    L := 0
    R := n
    while L < R:
        m := floor((L + R) / 2)
        if A[m] < T:
            L := m + 1
        else:
            R := m
    return L
```

使用 Binary Search，在搜索时会出现以下情况：

- A[m] > A[r], 说明 A[m, r] 无序，搜索区间 A[m+1, r]；
- A[m] < A[r], 说明 A[l, m] 无序，搜索区间 A[l, m]；
- A[m] = A[r], 此时不清楚该搜索哪个区间，为了缩减搜索区间，可以执行一个巧妙的操作：`r := r - 1`
  - 这个操作不会使数组越界，因为迭代条件保证了 r > l > 0。
  - 这个操作不会使最小值丢失：若 A[r] 是唯一最小值，则不满足条件 A[m] = A[r]；若 A[r] 不是唯一最小值，由于 m < r 且 A[m] = A[r]，所以还有最小值存在于 A[l, r-1] 区间。

Time: O(logn) ~ O(n)；在特例如 [1,1,1,1] 时会退化到 O(n)

参考：[leetcode/@Krahets](https://leetcode-cn.com/problems/find-minimum-in-rotated-sorted-array-ii/solution/154-find-minimum-in-rotated-sorted-array-ii-by-jyd/)

Examples:

```shell
0 1 2 3 4 5 6 7 8 9

2 3 3 3 4 5 6 0 0 1
              ^
l       m         r         A[m] > A[r], l = m + 1
          l   m   r         A[m] < A[r], r = m
          l m r             A[m] > A[r], l = m + 1
             l,r            l == r, break, res = 0

3 1 3 3 3 3 3
  ^
l     m     r               A[m] == A[r], r = r - 1
l     m   r                 A[m] == A[r], r = r - 1
...                         ...
l m r                       A[m] < A[r], r = m
l r                         A[m] > A[r], l = m + 1
 l,r                        l == r, break, res = 1

```
