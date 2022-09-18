#include <iostream>
#include <stack>
#include <string>
using namespace std;

/**
 * Two Stack
 *
 * Original approach
 */
class Solution {
public:
    string decodeString(string s) {
        stack<int> st_a;  // put the number
        stack<char> st_b; // put the string
        for (int i = 0; i < s.size(); i++) {
            if (isdigit(s[i])) { // parse string to get number
                int num = 0;
                while (isdigit(s[i])) {
                    num = num * 10 + s[i] - '0';
                    i++;
                }
                st_a.push(num);
                st_b.push(s[i]); // to push '['
            } else if (s[i] == ']') {
                string encoded;
                while (st_b.top() != '[') {
                    encoded = st_b.top() + encoded;
                    st_b.pop();
                }
                st_b.pop(); // to pop '['

                // decode the encoded string, and push characters to stack
                int k = st_a.top(); // times to repeat
                st_a.pop();
                for (int i = 0; i < k; i++) {
                    for (int j = 0; j < encoded.size(); j++) {
                        st_b.push(encoded[j]);
                    }
                }
            } else {
                st_b.push(s[i]);
            }
        }

        string res;
        while (!st_b.empty()) {
            res = st_b.top() + res;
            st_b.pop();
        }
        return res;
    }
};

int main() {

    Solution solution;
    string s;

    // aaabcbc
    s = "3[a]2[bc]";
    cout << solution.decodeString(s) << endl;

    // accaccacc
    s = "3[a2[c]]";
    cout << solution.decodeString(s) << endl;

    // abcabccdcdcdef
    s = "2[abc]3[cd]ef";
    cout << solution.decodeString(s) << endl;

    // abccdcdcdxyz
    s = "abc3[cd]xyz";
    cout << solution.decodeString(s) << endl;

    return 0;
}