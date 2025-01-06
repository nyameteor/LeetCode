# 1769. Minimum Number of Operations to Move All Balls to Each Box

- Difficulty: Medium
- Topics: Array, String, Prefix Sum
- Link: https://leetcode.com/problems/minimum-number-of-operations-to-move-all-balls-to-each-box/

## Description

You have `n` boxes. You are given a binary string `boxes` of length `n`, where `boxes[i]` is `'0'` if the `ith` box is **empty**, and `'1'` if it contains **one** ball.

In one operation, you can move **one** ball from a box to an adjacent box. Box `i` is adjacent to box `j` if `abs(i - j) == 1`. Note that after doing so, there may be more than one ball in some boxes.

Return an array `answer` of size `n`, where `answer[i]` is the **minimum** number of operations needed to move all the balls to the `ith` box.

Each `answer[i]` is calculated considering the **initial** state of the boxes.

**Example 1:**

```
**Input:** boxes = "110"
**Output:** [1,1,3]
**Explanation:** The answer for each box is as follows:
1) First box: you will have to move one ball from the second box to the first box in one operation.
2) Second box: you will have to move one ball from the first box to the second box in one operation.
3) Third box: you will have to move one ball from the first box to the third box in two operations, and move one ball from the second box to the third box in one operation.

```

**Example 2:**

```
**Input:** boxes = "001011"
**Output:** [11,8,5,4,3,4]
```

**Constraints:**

- `n == boxes.length`
- `1 <= n <= 2000`
- `boxes[i]` is either `'0'` or `'1'`.

## Solution

### Intuition

1. **Two Passes for Efficiency**:  
   Instead of calculating the cost for each box individually $(O(n^2))$, we use two passes (left-to-right and right-to-left) to compute the operations in $O(n)$.

2. **Left-to-Right Pass**:  
   - Keep track of the number of balls encountered (`leftBalls`).
   - Accumulate the total operations needed to move these balls to the current box (`leftMoves`).
   - Update the result for each box as we traverse.

3. **Right-to-Left Pass**:  
   - Similarly, track the number of balls on the right (`rightBalls`) and the operations needed to move them to the current box (`rightMoves`).
   - Add this to the result from the left pass.

4. **Final Result**:  
   - For each box, the total operations are the sum of the costs from the left and right passes.

This approach leverages cumulative sums to efficiently calculate the minimum operations required for each box, avoiding redundant calculations.
