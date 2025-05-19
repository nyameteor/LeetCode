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

### In-place Evaluation with Two Stacks

Use two stacks:

- One for operands (numbers)
- One for operators (`+`, `-`, `*`, `/`)

#### Operator Precedence

| Operator | Precedence |
| -------- | ---------- |
| `+`, `-` | 1          |
| `*`, `/` | 2          |

#### Key Idea

When processing an operator:

- While the top of the operator stack has **greater or equal precedence**, evaluate it.
- Then, push the current operator.

This ensures correct order of operations and left-to-right evaluation for same-precedence operators.

#### Examples

- `1 + 3` -> Evaluate immediately -> Result: `4`

- `1 + 3 / 5` -> `/` has higher precedence than `+`, so just push it

- `1 + 3 - 4` -> `-` has equal precedence to `+`, so evaluate `1 + 3`, then push `-`

#### Final Step

After parsing the expression, evaluate all remaining operators on the stack.
