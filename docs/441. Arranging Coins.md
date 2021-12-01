# 441. Arranging Coins

- Difficulty: Easy
- Topics: Math, Binary Search
- Link: https://leetcode.com/problems/arranging-coins/

## Description

You have `n` coins and you want to build a staircase with these coins. The staircase consists of `k` rows where the `ith` row has exactly `i` coins. The last row of the staircase **may be** incomplete.

Given the integer `n`, return _the number of **complete rows** of the staircase you will build_.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2021/04/09/arrangecoins1-grid.jpg)

```
Input: n = 5
Output: 2
Explanation: Because the 3rd row is incomplete, we return 2.
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2021/04/09/arrangecoins2-grid.jpg)

```
Input: n = 8
Output: 3
Explanation: Because the 4th row is incomplete, we return 3.
```

**Constraints:**

- `1 <= n <= 2^31 - 1`

## Solution

### Brute Force

琐碎的实现，从 1 开始累加判断，直至满足条件。

### Binary Search

本题可使用 Binary Search，因为解题过程可以抽象为从以下 sorted array 中，找到 target index：

```shell
1   2   3   4   5   6   ...
1   3   6   10  15  21  ...
```

本题没有直接显式给出 array， 但 array element 和 index 存在关系：`A[i] = i * (i + 1) / 2`。

搜索返回的结果为：

- 若 `n` 在 array 中找到，则返回对应索引；
- 若 `n` 在 array 中未找到，则返回下界的索引 `x`，满足 `A[x]` < n < `A[x+1]`，对应退出搜索时的右指针 `r` 。

Pseudocode:

```plaintext
function binary_search(T) is
    L := 0
    R := T # 将上界设为 T 是最简单的方法，应该可以用数学关系缩减上界
    while L ≤ R do
        m := floor((L + R) / 2)
        E := m * (m + 1) / 2    # 计算索引 m 对应的元素 E
        if E < T then
            L := m + 1
        else if E > T then
            R := m - 1
        else:
            return m
    return R    # 若未找到相等的 T，返回最接近的下界
```

另外参考 @zayne-siew 的[题解](<https://leetcode.com/problems/arranging-coins/discuss/1559984/C%2B%2BJavaPython-O(sqrt(n))-O(logn)-O(1)-Approaches-with-Explanation>)，写得很详细。

### Math

Todo
