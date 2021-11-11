#include <iostream>
#include <vector>
using namespace std;

/**
 * Refer:
 * https://github.com/haoel/leetcode/blob/master/algorithms/cpp/multiplyStrings/multiplyStrings.cpp
 */
class Solution {
  public:
    string multiply(string num1, string num2) {
        string res;
        int shift = 0;
        for (int i = num1.size() - 1; i >= 0; i--) {
            string s = multiplyOneDigit(num2, num1[i]);
            for (int k = 0; k < shift; k++) {
                s.insert(s.end(), '0');
            }
            res = numericAdd(res, s);
            shift++;
        }
        while (res[0] == '0' && res.size() > 1) { // remove leading '0'
            res.erase(0, 1);
        }
        return res;
    }

    string multiplyOneDigit(string num, char ch) {
        string s;
        int carry = 0;
        for (int i = num.size() - 1; i >= 0; i--) {
            int x = (num[i] - '0') * (ch - '0') + carry;
            carry = x / 10;
            s.insert(s.begin(), x % 10 + '0');
        }
        if (carry > 0) {
            s.insert(s.begin(), carry + '0');
        }
        return s;
    }

    string numericAdd(string num1, string num2) {
        string s;
        int carry = 0;
        int n1 = num1.size(), n2 = num2.size();

        int i, j;
        for (i = n1 - 1, j = n2 - 1; i >= 0 || j >= 0; i--, j--) {
            int x1 = i >= 0 ? num1[i] - '0' : 0;
            int x2 = j >= 0 ? num2[j] - '0' : 0;
            int x = x1 + x2 + carry;
            carry = x / 10;
            s.insert(s.begin(), x % 10 + '0');
        }
        if (carry > 0) {
            s.insert(s.begin(), carry + '0');
        }
        return s;
    }
};

/**
 * Original, not optimized
 */
class Solution2 {
  public:
    string multiply(string num1, string num2) {
        string res = "";
        for (int i = num1.size() - 1; i >= 0; i--) {
            int carry = 0, bit = 0;
            string part = "";
            for (int j = num2.size() - 1; j >= 0; j--) {
                int p = (num1[i] - '0') * (num2[j] - '0') + carry;
                carry = p / 10;
                part.insert(part.begin(), p % 10 + '0');
            }
            if (carry > 0) {
                part.insert(part.begin(), carry + '0');
            }
            for (int k = num1.size() - i - 2; k >= 0; k--) {
                part += '0';
            }
            res = numericAdd(res, part);
        }
        // remove leading '0'
        while (res[0] == '0' && res.size() > 1) {
            res.erase(0, 1);
        }

        return res;
    }

    string numericAdd(string num1, string num2) {
        string res = "";
        if (num1.size() < num2.size()) {
            string temp = num1;
            num1 = num2;
            num2 = temp;
        }
        for (int i = num1.size() - num2.size() - 1; i >= 0; i--) {
            num2 = '0' + num2;
        }
        int carry = 0, bit = 0;
        for (int i = num2.size() - 1; i >= 0; i--) {
            int p = (num2[i] - '0') + (num1[i] - '0') + carry;
            carry = p / 10;
            res.insert(res.begin(), p % 10 + '0');
        }
        if (carry > 0) {
            res.insert(res.begin(), '0' + carry);
        }
        return res;
    }
};

int main() {
    Solution solution;
    cout << solution.numericAdd("123", "456") << endl;
    cout << solution.multiplyOneDigit("123", '3') << endl;
    cout << solution.multiply("123", "456") << endl;
    cout << solution.multiply("12345", "45678") << endl;
    cout << solution.multiply("12345678987654321", "98765432188") << endl;
    cout << solution.multiply("9311", "0001") << endl;

    return 0;
}