#include <iostream>
using namespace std;

/**
 * Math & Bit Hack
 *
 * Using `n & (n - 1)` to figure out if `n` is
 * either 0 or an exact power of two.
 *
 * Refer: https://leetcode.com/problems/power-of-two/discuss/63974
 */
class Solution {
  public:
    bool isPowerOfTwo(int n) {
        if (n > 0) {
            return !(n & (n - 1));
        } else {
            return false;
        }
    }
};

/**
 * Count bits
 *
 * (Time Limit Exceeded)
 */
class Solution2 {
  public:
    bool isPowerOfTwo(int n) {
        while ((n & 1) == 0) {
            n = n >> 1;
        }
        return n == 1;
    }
};

int main() {
    Solution solution;

    cout << solution.isPowerOfTwo(1) << endl;
    cout << solution.isPowerOfTwo(16) << endl;
    cout << solution.isPowerOfTwo(3) << endl;

    return 0;
}