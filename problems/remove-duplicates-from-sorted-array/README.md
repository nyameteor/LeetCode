# 26. Remove Duplicates from Sorted Array

- Difficulty: Easy
- Topics: Array, Two Pointers
- Link: https://leetcode.com/problems/remove-duplicates-from-sorted-array/

## Description

- 输入一个非降序的数组 nums，要求就地删除重复的元素，使每个元素只出现一次。
- 将最终结果放入 nums 的前 k 个槽后返回 k。
- 要求就地改变，以 O(1) 内存完成。不要为别的数组申请额外空间。

## Solution

使用双指针，j 指针顺序读取，i 指针顺序写入。当 j 读取到新非重复元素时，i 指针前移一位并将 j 读取到的数值写入。
