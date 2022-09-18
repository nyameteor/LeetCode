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

Refer: https://leetcode.com/problems/score-of-parentheses/solution/

### Using Stack

Every position in the string has a `depth`. We need to maintain the score at the current depth we are on:

- for every `(`, we increase the depth, and our score at the new depth is 0.
- for every `)`, we add twice the score of the previous deeper part - except `()`, which has a score of 1.

We can use stack to implement this.

Example:

```plaintext
string  ()(())

stack   0                                   init
        0   0                               '(', stack.push(0)
        1                                   ')', t = stack.pop(), stack.top() += t > 0 ? 2 * t : 1;
        1   0                               '(', stack.push(0);
        1   0   0                           '(', stack.push(0);
        1   1                               ')', t = stack.pop(), stack.top() += t > 0 ? 2 * t : 1;
        3                                   ')', t = stack.pop(), stack.top() += t > 0 ? 2 * t : 1;
```
