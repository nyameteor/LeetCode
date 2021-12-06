# 1217. Minimum Cost to Move Chips to The Same Position

- Difficulty: Easy
- Topics: Array, Math, Greedy
- Link: https://leetcode.com/problems/minimum-cost-to-move-chips-to-the-same-position/

## Description

We have `n` chips, where the position of the `ith` chip is `position[i]`.

We need to move all the chips to **the same position**. In one step, we can change the position of the `ith` chip from `position[i]` to:

- `position[i] + 2` or `position[i] - 2` with `cost = 0`.
- `position[i] + 1` or `position[i] - 1` with `cost = 1`.

Return _the minimum cost_ needed to move all the chips to the same position.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2020/08/15/chips_e1.jpg)

```
Input: position = [1,2,3]
Output: 1
Explanation: First step: Move the chip at position 3 to position 1 with cost = 0.
Second step: Move the chip at position 2 to position 1 with cost = 1.
Total cost is 1.
```

**Example 2:**

![img](https://assets.leetcode.com/uploads/2020/08/15/chip_e2.jpg)

```
Input: position = [2,2,2,3,3]
Output: 2
Explanation: We can move the two chips at position  3 to position 2. Each move has cost = 1. The total cost = 2.
```

**Example 3:**

```
Input: position = [1,1000000000]
Output: 1
```

**Constraints:**

- `1 <= position.length <= 100`
- `1 <= position[i] <= 10^9`

## Solution

### Greedy

观察、发现规律后代码很简洁。

- 将 chip 移动 2 个位置是无代价的，故所有在偶数位上的 chip 可以无代价地移动到位置 0 上，所有在奇数位上的 chip 可以无代价地移动到位置 1 上。
- 比较位置 0 和 位置 1 上 chip 的数量（即比较偶数位与奇数位的数量），偶数位和奇数位之间移动的代价是 1，故 $minCost = min\{evenCnt, oddCnt\}$。

[Official Solution](https://leetcode.com/problems/minimum-cost-to-move-chips-to-the-same-position/solution/)
