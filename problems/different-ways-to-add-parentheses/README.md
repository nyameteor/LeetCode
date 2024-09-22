# 241. Different Ways to Add Parentheses

- Difficulty: Medium
- Topics: Math, String, Dynamic Programming, Recursion, Memoization
- Link: https://leetcode.com/problems/different-ways-to-add-parentheses/

## Description

Given a string `expression` of numbers and operators, return *all possible results from computing all the different possible ways to group numbers and operators*. You may return the answer in **any order**.

The test cases are generated such that the output values fit in a
32-bit integer and the number of different results does not exceed `104`.

**Example 1:**

```
**Input:** expression = "2-1-1"
**Output:** [0,2]
**Explanation:**
((2-1)-1) = 0 
(2-(1-1)) = 2
```

**Example 2:**

```
**Input:** expression = "2*3-4*5"
**Output:** [-34,-14,-10,-10,10]
**Explanation:**
(2*(3-(4*5))) = -34 
((2*3)-(4*5)) = -14 
((2*(3-4))*5) = -10 
(2*((3-4)*5)) = -10 
(((2*3)-4)*5) = 10
```

**Constraints:**

- `1 <= expression.length <= 20`
- `expression` consists of digits and the operator `'+'`, `'-'`, and `'*'`.
- All the integer values in the input expression are in the range `[0, 99]`.
- The integer values in the input expression do not have a leading `'-'` or `'+'` denoting the sign.

## Solution

We can start with the base cases:

1. If the expression is empty, return an empty list.
2. If the expression is one or two digits, return a list with that number.

For longer expressions, we can iterate through each character, and when we find an operator, we can do the following:

1. Evaluate the expression before the operator (left part results).
2. Evaluate the expression after the operator (right part results).
3. Combine the results of the left and right parts using the operator and return the final results.

References: https://leetcode.com/problems/different-ways-to-add-parentheses/editorial/
