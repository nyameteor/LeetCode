# 232. Implement Queue using Stacks

- Difficulty: Easy
- Topics: Stack, Design, Queue
- Link: https://leetcode.com/problems/implement-queue-using-stacks/

## Description

Implement a first in first out (FIFO) queue using only two stacks. The implemented queue should support all the functions of a normal queue (`push`, `peek`, `pop`, and `empty`).

Implement the `MyQueue` class:

- `void push(int x)` Pushes element x to the back of the queue.
- `int pop()` Removes the element from the front of the queue and returns it.
- `int peek()` Returns the element at the front of the queue.
- `boolean empty()` Returns `true` if the queue is empty, `false` otherwise.

**Notes:**

- You must use **only** standard operations of a stack, which means only `push to top`, `peek/pop from top`, `size`, and `is empty` operations are valid.
- Depending on your language, the stack may not be supported natively. You may simulate a stack using a list or deque (double-ended queue) as long as you use only a stack's standard operations.

**Example 1:**

```
Input
["MyQueue", "push", "push", "peek", "pop", "empty"]
[[], [1], [2], [], [], []]
Output
[null, null, null, 1, 1, false]

Explanation
MyQueue myQueue = new MyQueue();
myQueue.push(1); // queue is: [1]
myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
myQueue.peek(); // return 1
myQueue.pop(); // return 1, queue is [2]
myQueue.empty(); // return false
```

**Constraints:**

- `1 <= x <= 9`
- At most `100` calls will be made to `push`, `pop`, `peek`, and `empty`.
- All the calls to `pop` and `peek` are valid.

**Follow-up:** Can you implement the queue such that each operation is **[amortized](https://en.wikipedia.org/wiki/Amortized_analysis)** `O(1)` time complexity? In other words, performing `n` operations will take overall `O(n)` time even if one of those operations may take longer.

## Solution

使用栈实现队列。这是一道设计题，考察对栈和队列的理解。

由于栈输出的是输入序列的逆序，因此可以维护两个栈，一个负责输入，一个负责输出，通过两次逆序，达到和队列相同的正序输出的效果：

- queue.push(x) -> inputStack.push(x)
- queue.pop() -> 若 ouputStack 为空，将 inputStack 中的**全部**元素导入到 ouputStack（只有全部导入才能实现逆序）；若 outputStack 不为空，outpuStack.pop()
- queue.empty() -> 若 inputStack, outputStack 均为空，则 queue 为空
