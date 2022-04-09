# 347. Top K Frequent Elements

- Difficulty: Medium
- Topics: Array, Hash Table, Divide and Conquer, Sorting, Heap(Priority Queue), Bucket Sort, Counting, Quickselect
- Link: https://leetcode.com/problems/top-k-frequent-elements/

## Description

Given an integer array `nums` and an integer `k`, return _the_ `k` _most frequent elements_. You may return the answer in **any order**.

**Example 1:**

```
Input: nums = [1,1,1,2,2,3], k = 2
Output: [1,2]
```

**Example 2:**

```
Input: nums = [1], k = 1
Output: [1]
```

**Constraints:**

- `1 <= nums.length <= 10^5`
- `k` is in the range `[1, the number of unique elements in the array]`.
- It is **guaranteed** that the answer is **unique**.

**Follow up:** Your algorithm's time complexity must be better than `O(n log n)`, where n is the array's size.

## Solution

### Heap(Priority Queue)

1. build a hash map `elements -> count`.
2. build a max-heap with the maximum `count` at the top. push all elements of hash map into heap, then pop the heap top until the size of heap is `k`.
3. convert heap into an result array.

### Map + Bucket Sort

使用 Map 记录元素出现次数，并通过桶排序获得结果。

桶排序的好处在于用空间换时间，时间复杂度可以接近线性时间。参考：[Wikipedia/Bucket_sort](https://en.wikipedia.org/wiki/Bucket_sort)

遍历一次数组 nums，使用 map 记录每个元素的出现次数, key = element, value = element 的出现次数；同时记录最大的元素出现次数 M

使用桶排序获取前 k 个出现次数最多的元素，即根据元素的出现次数进行排序。由于本题返回的结果可以无序，所以桶内元素无需排序。

- 使用二维数组作为桶 buckets，桶的数量 = M + 1 (index 0 作为 dummy)，桶的大小 = ciel(nums.size() / M) (如果桶使用容器类型可以不考虑大小)。
- 遍历 map，将元素按照出现次数插入到桶中：insert(map.key) into buckets[map.value]
- 逆序（即从大到小）遍历 buckets，将桶中元素记录到结果数组 res 中，当 res.size() == k 时，返回 res 即可

```shell
nums    1 1 1 2 3 2 4 5 4

map     key    value
        1       3
        2       2
        3       1
        4       2
        5       1

buckets 0       1       2       3
        []     [3,5]    [2,4]   [1]
     (dummy)

bucket3 = [1]       res = [1]
bucket2 = [2, 4]    res = [1, 2, 4]

k.size = 3, return res
```
