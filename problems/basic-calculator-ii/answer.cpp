#include <iostream>
#include <stack>
#include <string>
using namespace std;

/**
 * Inplace Evaluation
 *
 * Refer:
 * https://github.com/haoel/leetcode/blob/master/algorithms/cpp/basicCalculator/BasicCalculator.II.cpp
 */
class Solution {
public:
    int calculate(string s) {
        stack<int> num_stk; // put the numbers
        stack<char> op_stk; // put the operators
        for (int i = 0; i < size(s); i++) {
            if (isspace(s[i])) {
                continue;
            } else if (isdigit(s[i])) { // read integer
                string num;
                num += s[i];
                while (isdigit(s[i + 1])) {
                    num += s[i + 1];
                    i++;
                }
                num_stk.push(stoi(num));
            } else if (isOperator(s[i])) {
                while (!empty(op_stk) &&
                       priority(op_stk.top()) >= priority(s[i])) {
                    calculate_once(num_stk, op_stk);
                }
                op_stk.push(s[i]);
            }
        }
        while (!empty(op_stk)) {
            calculate_once(num_stk, op_stk);
        }
        return num_stk.top();
    }

    void calculate_once(stack<int> &num_stk, stack<char> &op_stk) {
        int b = num_stk.top();
        num_stk.pop();
        int a = num_stk.top();
        num_stk.pop();
        char op = op_stk.top();
        op_stk.pop();

        num_stk.push(calculate_exp(a, b, op));
    }

    bool isOperator(char c) {
        return c == '+' || c == '-' || c == '*' || c == '/';
    }

    int priority(char c) {
        if (c == '*' || c == '/') {
            return 2;
        } else if (c == '+' || c == '-') {
            return 1;
        } else {
            return 0;
        }
    }

    int calculate_exp(int a, int b, char op) {
        switch (op) {
        case '+':
            return a + b;
        case '-':
            return a - b;
        case '*':
            return a * b;
        case '/':
            return a / b;
        }
        return -1;
    }
};

/**
 * Convert to Postfix & Evaluate
 *
 * Refer: https://leetcode.com/problems/basic-calculator-ii/discuss/1646099
 */

int main() {
    Solution solution;
    string s;

    s = "3+2*2";
    cout << solution.calculate(s) << endl;

    s = " 3/2 ";
    cout << solution.calculate(s) << endl;

    s = " 3+5 / 2 ";
    cout << solution.calculate(s) << endl;

    s = " 42";
    cout << solution.calculate(s) << endl;

    s = "1-1+1";
    cout << solution.calculate(s) << endl;

    return 0;
}