#include <iostream>
#include <map>
#include <vector>
using namespace std;

/**
 * Dynamic Programming
 *
 * Time: O(m*n), Space: O(m+n)
 */
class Solution {
  public:
    int uniquePaths(int m, int n) {
        map<pair<int, int>, int> memo;
        return dfs(m, n, memo);
    }

    int dfs(int m, int n, map<pair<int, int>, int> &memo) {
        if (memo.find({m, n}) != end(memo)) {
            return memo[{m, n}];
        }
        if (m == 0 || n == 0) {
            return 0;
        }
        if (m == 1 && n == 1) {
            return 1;
        }
        memo[{m, n}] = dfs(m - 1, n, memo) + dfs(m, n - 1, memo);
        return memo[{m, n}];
    }
};

/**
 * Brute Force
 *
 * Time: O(2^(m+n)), Space: O(m+n)
 */
class Solution2 {
  public:
    int uniquePaths(int m, int n) {
        if (m == 0 || n == 0) {
            return 0;
        }
        if (m == 1 && n == 1) {
            return 1;
        }
        return uniquePaths(m - 1, n) + uniquePaths(m, n - 1);
    }
};

int main() {
    Solution solution;

    cout << solution.uniquePaths(7, 3) << endl;
    cout << solution.uniquePaths(3, 3) << endl;

    return 0;
}