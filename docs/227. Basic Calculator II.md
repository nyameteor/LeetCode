# 227. Basic Calculator II

- Difficulty: Medium
- Topics: Math, String, Stack
- Link: https://leetcode.com/problems/basic-calculator-ii/

## Description

Given a string `s` which represents an expression, _evaluate this expression and return its value_.

The integer division should truncate toward zero.

You may assume that the given expression is always valid. All intermediate results will be in the range of `[-2^31, 2^31 - 1]`.

**Note:** You are not allowed to use any built-in function which evaluates strings as mathematical expressions, such as `eval()`.

**Example 1:**

```
Input: s = "3+2*2"
Output: 7
```

**Example 2:**

```
Input: s = " 3/2 "
Output: 1
```

**Example 3:**

```
Input: s = " 3+5 / 2 "
Output: 5
```

**Constraints:**

- `1 <= s.length <= 3 * 10^5`
- `s` consists of integers and operators `('+', '-', '*', '/')` separated by some number of spaces.
- `s` represents **a valid expression**.
- All the integers in the expression are non-negative integers in the range `[0, 2^31 - 1]`.
- The answer is **guaranteed** to fit in a **32-bit integer**.

## Solution

### Inplace Evaluation - Two Stacks

Use two stacks, one to put numbers, another to put operators.

The priority of operators:

| operator | priority |
| -------- | -------- |
| + -      | 1        |
| \* /     | 2        |

Before put current operator, while the priority of previous operator is **greater than or equal to** the priority of current, calculate expression for once. For example:

```plaintext
1 + 3

    3
+   1
%   %

------ case 1 ------

1 + 3 / 5

priority of '+' is less than '/', so simply add current operator

    5
/   3
+   1
%   %

------ case 2 ------

1 + 3 - 4

priority of '+' is equal to '-', so calculate expression for once

    4
%   %

then add current operator

    4
-   4
%   %

------ cases ends ------

```

After we have read the full expression, calculate it while stack of operator is not empty.
