#include <algorithm>
#include <iostream>
#include <vector>
using namespace std;

/**
 * Using two stacks
 *  1. one stack is the real stack to store the data.
 *  2. another stack store the minimum number when it changes.
 */
class MinStack {
private:
    vector<int> stack;
    vector<int> minStack;

public:
    MinStack() {
        stack = {};
        minStack = {};
    }

    void push(int val) {
        stack.push_back(val);
        if (minStack.size() == 0 || val <= minStack.back()) {
            minStack.push_back(val);
        }
    }

    void pop() {
        if (stack.size() != 0) {
            int val = stack.back();
            stack.pop_back();
            if (minStack.size() != 0 && val == minStack.back()) {
                minStack.pop_back();
            }
        }
    }

    int top() {
        if (stack.size() != 0) {
            return stack.back();
        } else {
            return -1;
        }
    }

    int getMin() {
        if (minStack.size() != 0) {
            return minStack.back();
        } else {
            return -1;
        }
    }
};

/**
 * Your MinStack object will be instantiated and called as such:
 * MinStack* obj = new MinStack();
 * obj->push(val);
 * obj->pop();
 * int param_3 = obj->top();
 * int param_4 = obj->getMin();
 */
int main() {
    MinStack *obj = new MinStack();
    obj->push(3);
    obj->push(1);
    obj->push(2);
    cout << obj->top() << endl;
    cout << obj->getMin() << endl;

    obj = new MinStack();
    obj->push(0);
    obj->push(1);
    obj->push(0);
    cout << obj->getMin() << endl;
    obj->pop();
    cout << obj->getMin() << endl;
}