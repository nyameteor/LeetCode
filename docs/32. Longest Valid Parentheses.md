# 32. Longest Valid Parentheses

- Difficulty: Hard
- Topics: String, Dynamic Programming, Stack
- Link: https://leetcode.com/problems/longest-valid-parentheses/

## Description

Given a string containing just the characters `'('` and `')'`, find the length of the longest valid (well-formed) parentheses substring.

**Example 1:**

```
Input: s = "(()"
Output: 2
Explanation: The longest valid parentheses substring is "()".
```

**Example 2:**

```
Input: s = ")()())"
Output: 4
Explanation: The longest valid parentheses substring is "()()".
```

**Example 3:**

```
Input: s = ""
Output: 0
```

**Constraints:**

- `0 <= s.length <= 3 * 10^4`
- `s[i]` is `'('`, or `')'`.

## Solution

Refer: https://leetcode.com/problems/longest-valid-parentheses/solution/

### Using Stack

We can make use of stack while scanning the given string to:

1. check if the string scanned so far is valid.
2. find the length of the longest valid string.

In order to do so:

- start by pushing `-1` onto the stack(for handle edge case).
- for every `(`, push `current_element_index` onto the stack.
- for every `)`, pop the topmost element(if stack is not empty).
  - if the stack is empty, push `current_element_index` onto the stack.
  - else, the length of the currently valid string is `current_element_index - stack.top()`.

Example:

```plaintext
index   0   1   2   3   4   5   6   7   8
string  )   (   )   )   (   (   (   )   )


stack   -1                                          init


stack   0                                           ele = ')' => st.pop(); st.empty() => st.push(ele_index);


stack   0   1                                       ele = '(' => st.push(ele_index);


stack   0                                           ele = ')' => st.pop(); !st.empty() => len = ele_index - st.top();
len     2 - 0 = 2

stack   3                                           ele = ')' => st.pop(); st.emtpy() => st.push(ele_index);


stack   3   4                                       ele = '(' => st.push(ele_index);


stack   3   4   5                                   ele = '(' => st.push(ele_index);


stack   3   4   5   6                               ele = '(' => st.push(ele_index);


stack   3   4   5                                   ele = ')' => st.pop(); !st.empty() => len = ele_index - st.top();
len     7 - 5 = 2

stack   3   4                                       ele = ')' => st.pop(); !st.empty() => len = ele_index - st.top();
len     8 - 4 = 4


max_len = 4

---------------------------------------

index   0   1   2
string  (   )   )

stack   -1                                          init


stack   -1  0                                       ele = '(' => st.push(ele_index);


stack   -1                                          ele = ')' => st.pop(); !st.empty() => len = ele_index - st.top();
len     1 - (-1) = 2

stack   2                                           ele = ')' => st.pop(); st.empty() => st.push(ele_index);
```

### Dynamic Programming

Todo
