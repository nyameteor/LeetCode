# 856. Score of Parentheses

- Difficulty: Medium
- Topics: String, Stack
- Link: https://leetcode.com/problems/score-of-parentheses/

## Description

Given a balanced parentheses string `s`, return _the **score** of the string_.

The **score** of a balanced parentheses string is based on the following rule:

- `"()"` has score `1`.
- `AB` has score `A + B`, where `A` and `B` are balanced parentheses strings.
- `(A)` has score `2 * A`, where `A` is a balanced parentheses string.

**Example 1:**

```
Input: s = "()"
Output: 1
```

**Example 2:**

```
Input: s = "(())"
Output: 2
```

**Example 3:**

```
Input: s = "()()"
Output: 2
```

**Constraints:**

- `2 <= s.length <= 50`
- `s` consists of only `'('` and `')'`.
- `s` is a balanced parentheses string.

## Solution

### Split and Eval (Recursive)

- Recursively split the string into balanced parts.
- For `"()"`, return `1`.
- For `"(A)"`, return `2 * score(A)`.
- For `"AB"`, return `score(A) + score(B)` by summing recursively.

### Stack Based (Iterative)

- Use a stack to track score at each depth.
- Push `0` for each `'('`, pop on `')'`.
- On pop:
  - If top is `0`, it's a `"()"` -> score `1`.
  - Else, it's nested -> score `2 * top`.
- Add the computed score to the previous stack level.
