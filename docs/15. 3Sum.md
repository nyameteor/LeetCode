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

找出 nums 中的三个元素 nums[i], nums[j], nums[k] 使得三者相加之和等于 0，需要找出所有满足条件的三元组。注意：答案中不能包含重复的三元组。

比较好的一个解法是使用双指针(Refer: [三数之和](https://github.com/youngyangyang04/leetcode-master/blob/master/problems/0015.%E4%B8%89%E6%95%B0%E4%B9%8B%E5%92%8C.md))：

![](https://camo.githubusercontent.com/3ee1e9d4e153718a7c15146e6b619968f18ba06a39ba732d6d1fa875238483f5/68747470733a2f2f636f64652d7468696e6b696e672e63646e2e626365626f732e636f6d2f676966732f31352e2545342542382538392545362539352542302545342542392538422545352539322538432e676966)

首先将数组按照升序排序。

使用指针 i 遍历 nums 数组，定义左指针 l 的位置在 i+1 , 定义右指针 r 的位置在数组的结尾。

设 sum = nums[i] + nums[l] + nums[r]

- sum == 0，将三元组添加到结果中，同时 l 向右移动，r 向左移动，直到 l 和 r 相遇；
- sum < 0，说明结果偏小了，l 需要向右移动，使 sum 大一些，直到 l 和 r 相遇；
- sum > 0，说明结果偏大了，r 需要向左移动，使 sum 小一些，直到 l 和 r 相遇；

重要的优化：i, l, r 在遍历时都需要去重相同的元素，防止在答案中添加重复的元组，并可以防止超时：

```cpp
while (l < r && nums[l] == nums[l + 1])
  l++;
```

一个遍历过程实例：

```shell
sorted_nums:
-4 -1 -1  0  1  2   sum:
 i  l           r   -4 + -1 + 2 = -3 <  0, l -> nums[l] < nums[l+1], l++
 i        l     r   -4 +  0 + 2 = -2 <  0, l -> nums[l] < nums[l+1], l++
 i           l  r   -4 +  1 + 2 = -1 <  0, l -> nums[l] < nums[l+1], l++
                                           i -> nums[i] < nums[i+1], i++
    i  l        r   -1 + -1 + 2 =  0 == 0, l -> nums[l] < nums[l+1], l++; r -> nums[r] > nums[r-1], r++; add to ans
    i     l     r   -1 +  0 + 2 =  1 >  0, r -> nums[r] > nums[r-1], r++
    i     l  r      -1 +  0 + 1 =  0 == 0, l -> nums[l] < nums[l+1], l++; r -> nums[r] > nums[r-1], r++; add to ans
                                           i -> nums[i] < nums[i+1], i++
          i  l  r   0 +   1 + 2 =  3 >  0, l -> nums[r] > nums[r-1], l++
```

- time: O(n^2)
- space: O(n^2)
