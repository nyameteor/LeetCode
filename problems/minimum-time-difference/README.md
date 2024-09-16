# 539. Minimum Time Difference

- Difficulty: Medium
- Topics: Array, Math, String, Sorting
- Link: https://leetcode.com/problems/minimum-time-difference/

## Description

Given a list of 24-hour clock time points in **"HH:MM"** format, return *the minimum **minutes** difference between any two time-points in the list*.

**Example 1:**

```
**Input:** timePoints = ["23:59","00:00"]
**Output:** 1

```

**Example 2:**

```
**Input:** timePoints = ["00:00","23:59","00:00"]
**Output:** 0

```

**Constraints:**

- `2 <= timePoints.length <= 2 * 10^4`
- `timePoints[i]` is in the format **"HH:MM"**.

## Solution

We can convert the time points into `minutes` and sort them in ascending order, then we can compare the difference between each adjacent pair in `minutes`, as well as the difference between the first and last element, to find the minimum time difference.
