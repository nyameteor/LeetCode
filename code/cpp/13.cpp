#include <iostream>
#include <string>
using namespace std;

class Solution {
  public:
    int romanToInt(string s) {
        int res = 0;
        int sz = size(s);
        for (int i = 0; i < sz; i++) {
            int cur = symbolToInt(s[i]);
            // avoid next letter out of bounds
            int next = i < sz - 1 ? symbolToInt(s[i + 1]) : 0;
            // if current letter is less than its next one,
            // such as: IV(4), IX(9), XL(40), XC(90)
            if (cur < next) {
                res -= cur;
            } else {
                res += cur;
            }
        }
        return res;
    }

    int romanToInt2(string s) {
        int res = 0;
        res += symbolToInt(s[0]);
        for (int i = 1; i < size(s); i++) {
            int prev = symbolToInt(s[i - 1]);
            int cur = symbolToInt(s[i]);
            if (prev < cur) {
                res += -prev - prev + cur;
            } else {
                res += cur;
            }
        }
        return res;
    }

    int symbolToInt(char ch) {
        int v = 0;
        switch (ch) {
        case 'I':
            v = 1;
            break;
        case 'V':
            v = 5;
            break;
        case 'X':
            v = 10;
            break;
        case 'L':
            v = 50;
            break;
        case 'C':
            v = 100;
            break;
        case 'D':
            v = 500;
            break;
        case 'M':
            v = 1000;
            break;
        default:
            break;
        }
        return v;
    }
};

int main() {
    Solution solution;
    string s;

    s = "III";
    cout << solution.romanToInt(s) << endl;
    s = "X";
    cout << solution.romanToInt(s) << endl;
    s = "LVIII";
    cout << solution.romanToInt(s) << endl;
    s = "MCMXCIV";
    cout << solution.romanToInt(s) << endl;

    return 0;
}