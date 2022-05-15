#include <iostream>
#include <string>
using namespace std;

class Solution {
public:
    int myAtoi(string s) {
        if (size(s) == 0 || s[0] == '\0') {
            return 0;
        }

        int ret = 0;

        int i = 0;
        for (; isspace(s[i]); i++)
            ;

        bool neg = false;
        if (s[i] == '-' || s[i] == '+') {
            neg = (s[i] == '-');
            i++;
        }

        for (; isdigit(s[i]); i++) {
            int digit = (s[i] - '0');
            if (neg) {
                if (-ret < (INT32_MIN + digit) / 10) {
                    return INT32_MIN;
                }
            } else {
                if (ret > (INT32_MAX - digit) / 10) {
                    return INT32_MAX;
                }
            }

            if (INT32_MAX - 10 * ret < digit) {
                return neg ? INT32_MIN : INT32_MAX;
            } else {
                ret = 10 * ret + digit;
            }
        }

        return neg ? -ret : ret;
    }
};

int main() {
    Solution solution;
    string s;

    s = "42";
    cout << solution.myAtoi(s) << endl;

    s = "   -42";
    cout << solution.myAtoi(s) << endl;

    s = "4193 with words";
    cout << solution.myAtoi(s) << endl;

    s = "2147483647";
    cout << solution.myAtoi(s) << endl;

    s = "-2147483648";
    cout << solution.myAtoi(s) << endl;

    s = "2147483648";
    cout << solution.myAtoi(s) << endl;

    s = "-2147483649";
    cout << solution.myAtoi(s) << endl;

    return 0;
}