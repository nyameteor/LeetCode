#include <iostream>
using namespace std;

/**
 * XOR + Count bits
 *
 * Time: O(n), Space: O(1)
 */
class Solution {
public:
    int hammingDistance(int x, int y) {
        int c = x ^ y; // Bitwise XOR

        int res = 0;
        while (c > 0) {
            res += c & 1; // Bitwise AND
            c >>= 1;
        }
        return res;
    }
};

/**
 * XOR + Count bits
 *
 * Time: O(n), Space: O(1)
 */
class Solution2 {
public:
    int hammingDistance(int x, int y) {
        int c = x ^ y;

        // Decimal to Binary
        int res = 0;
        while (c != 0) {
            if (c % 2 == 1) {
                res += 1;
            }
            c /= 2;
        }
        return res;
    }
};

int main() {
    Solution solution;

    cout << solution.hammingDistance(1, 4) << endl;
    cout << solution.hammingDistance(3, 1) << endl;

    return 0;
}