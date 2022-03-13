#include <cassert>
#include <iostream>
#include <stack>
#include <string>
using namespace std;

class Solution {
  public:
    int longestValidParentheses(string s) {
        int max_len = 0;
        stack<int> st;
        st.push(-1);
        for (int i = 0; i < size(s); i++) {
            if (s[i] == '(') {
                st.push(i);
            } else if (s[i] == ')') {
                if (!empty(st)) {
                    st.pop();
                }
                if (empty(st)) {
                    st.push(i);
                } else {
                    max_len = max(max_len, i - st.top());
                }
            }
        }

        return max_len;
    }
};

int main() {
    Solution solution;
    string s;

    s = "(()";
    assert(solution.longestValidParentheses(s) == 2);

    s = ")()())";
    assert(solution.longestValidParentheses(s) == 4);

    s = "()(()";
    assert(solution.longestValidParentheses(s) == 2);

    s = "";
    assert(solution.longestValidParentheses(s) == 0);
}