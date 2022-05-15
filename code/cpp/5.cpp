#include <cmath>
#include <iostream>
#include <string>
#include <vector>
using namespace std;

/**
 * Dynamic Programming
 *
 * P(i, j): is substring s[i...j] a palindrome
 *
 * P(i, j) == P(i+1, j-1) && s[i] == s[j]
 *
 * Base cases:
 * P(i, i) = true
 * P(i, i+1) = (s[i] == s[i+1])
 *
 * Time: O(n^2), Space: O(n^2)
 */
class Solution {
public:
    string longestPalindrome(string s) {
        int N = size(s);
        if (N < 2) {
            return s;
        }

        vector<vector<bool>> table(N, vector<bool>(N, false));
        // init table value with base case
        for (int i = 0; i < N; i++) {
            table[i][i] = true;
        }

        string res = "";
        res += s[0];
        for (int i = N - 1; i >= 0; i--) {
            for (int j = i + 1; j < N; j++) {
                // if two left and right end letters are the same.
                if (s[i] == s[j]) {
                    // if it is 2 characters or its substring is palindrome
                    if (j - i == 1 || table[i + 1][j - 1]) {
                        table[i][j] = true;
                        // check for longest palindrome substring
                        if (j + 1 - i > size(res)) {
                            res = s.substr(i, j + 1 - i);
                        }
                    }
                }
            }
        }
        return res;
    }
};

/**
 * Brute Force
 *
 * Time: O(n^3), Space: O(1)
 */
class Solution2 {
public:
    string longestPalindrome(string s) {
        string res = "";
        int N = size(s);
        int i, j;
        for (i = 0; i < N; i++) {
            for (j = i; j < N; j++) {
                if (isPalindrome(s.substr(i, j + 1 - i))) {
                    if (size(res) < j + 1 - i) {
                        res = s.substr(i, j + 1 - i);
                    }
                }
            }
        }
        return res;
    }

    bool isPalindrome(string s) {
        int N = size(s);
        for (int i = 0; i < N / 2; i++) {
            if (s[i] != s[N - i - 1]) {
                return false;
            }
        }
        return true;
    }
};

int main() {
    Solution solution;
    string s;

    s = "babad";
    cout << solution.longestPalindrome(s) << endl;

    s = "cbbd";
    cout << solution.longestPalindrome(s) << endl;

    s = "aacabdkacaa";
    cout << solution.longestPalindrome(s) << endl;

    return 0;
}