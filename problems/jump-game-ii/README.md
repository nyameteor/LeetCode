# 45. Jump Game II

- Difficulty: Medium
- Topics: Array, Dynamic Programming, Greedy
- Link: https://leetcode.com/problems/jump-game-ii/

## Description

Given an array of non-negative integers `nums`, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Your goal is to reach the last index in the minimum number of jumps.

You can assume that you can always reach the last index.

**Example 1:**

```
Input: nums = [2,3,1,1,4]
Output: 2
Explanation: The minimum number of jumps to reach the last index is 2. Jump 1 step from index 0 to 1, then 3 steps to the last index.
```

**Example 2:**

```
Input: nums = [2,3,0,1,4]
Output: 2
```

**Constraints:**

- `1 <= nums.length <= 104`
- `0 <= nums[i] <= 1000`

## Solution

### Greedy

- 将步长数组的长度记为 `n`；
- 将当前位置记为 `curPos`；
- 将从当前位置能到达的最远位置记为 `coverPos`；
- 将步数记为 `step`。

遍历步长数组：

1. 维护 `coverPos` 的值。若扫描过程中实际能到达的位置大于 `coverPos`，更新 `coverPos` 并令 `step += 1`。
2. 若 `coverPos >= n-1`，说明已到达终点，返回 `step`。
3. 遍历 `curPos` 到 `coverPos` 之间的位置，找到其中能到达最远位置的位置，记为 `nextPos`。
4. 令 `curPos = nextPos`。

**References:**

- https://books.halfrost.com/leetcode/ChapterFour/0001~0099/0045.Jump-Game-II/
- https://github.com/haoel/leetcode/blob/master/algorithms/cpp/jumpGame/jumpGame.II.cpp
