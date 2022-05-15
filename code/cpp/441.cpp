#include <iostream>
#include <vector>
using namespace std;

/**
 * Binary Search
 *
 * - If target is found in the sequence, return the corresponding index.
 * - If target is not found in the sequence, return the index x of the lower
 * bound, satisfying A[x] <n <A[x+1].
 *
 * Time: O(logN), Space: O(1)
 */
class Solution {
public:
    int arrangeCoins(int n) {
        long l = 1,
             r = n; // n is the simplest upper bound, we might shrink it by
                    // using some math

        while (l <= r) {
            long m = (l + r) / 2;
            long e =
                m * (m + 1) / 2; // calculate element e corresponding to index m
            if (e < n) {
                l = m + 1;
            } else if (e > n) {
                r = m - 1;
            } else {
                return m;
            }
        }

        return int(r); // if not found, then return lower bound
    }
};

/**
 * Simulation with brute force, trivial approach.
 *
 * Time: O(sqrt(N)), Space: O(1)
 */
class Solution2 {
public:
    int arrangeCoins(int n) {
        int i = 1;
        while (i <= n) {
            n -= i;
            i++;
        }
        return i - 1;
    }
};

int main() {
    Solution solution;

    cout << solution.arrangeCoins(1) << endl;
    cout << solution.arrangeCoins(5) << endl;
    cout << solution.arrangeCoins(8) << endl;
    cout << solution.arrangeCoins(65536) << endl;

    return 0;
}