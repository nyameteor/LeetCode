# 155. Min Stack

- Difficulty: Easy
- Topics: Stack, Design
- Link: https://leetcode.com/problems/min-stack/

## Description

Design a stack that supports push, pop, top, and retrieving the minimum element in constant time.

Implement the `MinStack` class:

- `MinStack()` initializes the stack object.
- `void push(int val)` pushes the element `val` onto the stack.
- `void pop()` removes the element on the top of the stack.
- `int top()` gets the top element of the stack.
- `int getMin()` retrieves the minimum element in the stack.

**Example 1:**

```
Input
["MinStack","push","push","push","getMin","pop","top","getMin"]
[[],[-2],[0],[-3],[],[],[],[]]

Output
[null,null,null,null,-3,null,0,-2]

Explanation
MinStack minStack = new MinStack();
minStack.push(-2);
minStack.push(0);
minStack.push(-3);
minStack.getMin(); // return -3
minStack.pop();
minStack.top();    // return 0
minStack.getMin(); // return -2
```

**Constraints:**

- `-2^31 <= val <= 2^31 - 1`
- Methods `pop`, `top` and `getMin` operations will always be called on **non-empty** stacks.
- At most `3 * 10^4` calls will be made to `push`, `pop`, `top`, and `getMin`.

## Solution

### Two Stacks

使用两个栈：

- 一个栈 stack 存储真正的数据
- 一个栈 minStack 存储**当前** stack 中的最小值

栈操作：

- 当 stack 插入数据 val 时，检查 val 是否为当前 stack 中的最小值，若是则 val 同时插入到 minStack
- 当 stack 弹出数据 val 时，检查 val 是否为当前 stack 中的最小值，若是则同时弹出 minStack 的栈顶

```shell
# Example
5 2 1 4 3 1     <- keep push to stack
5 2 1 1         <- min stack will be like this
```

Refer: [haoel/minStack.cpp](https://github.com/haoel/leetcode/blob/master/algorithms/cpp/minStack/minStack.cpp)
