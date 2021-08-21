import java.util.Stack;

class MyQueue {

  Stack<Integer> input = new Stack<>();
  Stack<Integer> output = new Stack<>();

  /** Initialize your data structure here. */
  public MyQueue() {

  }

  /** Push element x to the back of queue. */
  public void push(int x) {
    input.add(x);
  }

  /** Removes the element from in front of queue and returns that element. */
  public int pop() {
    if (!output.isEmpty()) {
      return output.pop();
    } else {
      while (!input.isEmpty()) {
        output.add(input.pop());
      }
      return output.pop();
    }
  }

  /** Get the front element. */
  public int peek() {
    if (!output.isEmpty()) {
      return output.peek();
    } else {
      while (!input.isEmpty()) {
        output.add(input.pop());
      }
      return output.peek();
    }
  }

  /** Returns whether the queue is empty. */
  public boolean empty() {
    return input.isEmpty() && output.isEmpty();
  }
}

class Solution {
  public static void main(String[] args) {
    MyQueue myQueue = new MyQueue();
    myQueue.push(1); // queue is: [1]
    myQueue.push(2); // queue is: [1, 2] (leftmost is front of the queue)
    assert myQueue.peek() == 1; // return 1
    assert myQueue.pop() == 1; // return 1, queue is [2]
    assert myQueue.empty() == false; // return false
  }
}