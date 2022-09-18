# 849. Maximize Distance to Closest Person

- Difficulty: Medium
- Topics: Array
- Link: https://leetcode.com/problems/maximize-distance-to-closest-person/

## Description

You are given an array representing a row of `seats` where `seats[i] = 1` represents a person sitting in the `ith` seat, and `seats[i] = 0` represents that the `ith` seat is empty **(0-indexed)**.

There is at least one empty seat, and at least one person sitting.

Alex wants to sit in the seat such that the distance between him and the closest person to him is maximized.

Return _that maximum distance to the closest person_.

**Example 1:**

![img](https://assets.leetcode.com/uploads/2020/09/10/distance.jpg)

```
Input: seats = [1,0,0,0,1,0,1]
Output: 2
Explanation:
If Alex sits in the second open seat (i.e. seats[2]), then the closest person has distance 2.
If Alex sits in any other open seat, the closest person has distance 1.
Thus, the maximum distance to the closest person is 2.
```

**Example 2:**

```
Input: seats = [1,0,0,0]
Output: 3
Explanation:
If Alex sits in the last seat (i.e. seats[3]), the closest person is 3 seats away.
This is the maximum distance possible, so the answer is 3.
```

**Example 3:**

```
Input: seats = [0,1]
Output: 1
```

**Constraints:**

- `2 <= seats.length <= 2 * 10^4`
- `seats[i]` is `0` or `1`.
- At least one seat is **empty**.
- At least one seat is **occupied**.

## Solution

### Two Pointers

Scan the array `seats`, use two pointers to represent occupied seats, `i` is previous, `j` is current.

the maximize distance to closest person is `distance = floor((j - i) / 2)`.

Edge cases:

- `j` is the first occupied seat(`i` can be represented as `i == -1`), then `distance = j - 0`
- `i` is the last occupied seat(`j` can be represented as `j >= size(seats)`), then `distance = (size(seats) - 1) - i`, which `size(seats) - 1` is the index of last seat.
