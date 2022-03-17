#include <cassert>
#include <stack>
#include <string>
using namespace std;

class Solution {
  public:
    int scoreOfParentheses(string s) {
        stack<int> st;
        st.push(0); // avoid edge case

        for (auto x : s) {
            if (x == '(') {
                st.push(0);
            } else {
                int t = st.top();
                st.pop();
                st.top() += t > 0 ? 2 * t : 1;
            }
        }

        return st.top();
    }
};

int main() {
    Solution solution;
    string s;
    int res;

    s = "(())";
    res = 2;
    assert(solution.scoreOfParentheses(s) == res);

    s = "(()(()))";
    res = 6;
    assert(solution.scoreOfParentheses(s) == res);

    return 0;
}