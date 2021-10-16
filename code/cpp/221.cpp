#include <iostream>
#include <vector>
using namespace std;

/**
 * Brute Force (Optimized)
 */
class Solution {
  public:
    int maximalSquare(vector<vector<char>> &matrix) {
        int maxLen = 0;
        int n = matrix.size(), m = n > 0 ? matrix[0].size() : 0;

        for (int i = 0; i < n; i++) {
            for (int j = 0; j < m; j++) {
                if (matrix[i][j] == '1') {
                    // find the wanted square from current cell
                    int len = 1;      // length of wanted square
                    bool flag = true; // is wanted square
                    while (i + len < n && j + len < m && flag) {
                        for (int k = i; k <= i + len; k++) {
                            if (matrix[k][j + len] == '0') {
                                flag = false;
                                break;
                            }
                        }
                        for (int k = j; k <= j + len; k++) {
                            if (matrix[i + len][k] == '0') {
                                flag = false;
                                break;
                            }
                        }
                        if (flag) {
                            len++;
                        }
                    }
                    maxLen = max(len, maxLen);
                }
            }
        }

        return maxLen * maxLen;
    }
};

/**
 * Dynamic Programming (Tabulation)
 *
 * Recommended
 */
class Solution2 {
  public:
    int maximalSquare(vector<vector<char>> &matrix) {
        int r = matrix.size(), c = r > 0 ? matrix[0].size() : 0;
        // add dummy row and cloum to avoid corner case
        vector<vector<int>> table(r + 1, vector<int>(c + 1, 0));
        int maxLen = 0;

        for (int i = 1; i <= r; i++) {
            for (int j = 1; j <= c; j++) {
                if (matrix[i - 1][j - 1] == '1') {
                    table[i][j] = min(table[i - 1][j - 1],
                                      min(table[i - 1][j], table[i][j - 1])) +
                                  1;
                    maxLen = max(maxLen, table[i][j]);
                }
            }
        }

        return maxLen * maxLen;
    };
};

/**
 * Dynamic Programming (Memoization)
 *
 * Have to use lots of temporary variables, need to be optimized
 */
class Solution3 {
  public:
    int maximalSquare(vector<vector<char>> &matrix) {
        int n = matrix.size(), m = n > 0 ? matrix[0].size() : 0;
        // add dummy row and cloum to avoid corner case
        vector<vector<int>> memo(n + 1, vector<int>(m + 1, 0));
        vector<vector<bool>> memoed(n + 1, vector<bool>(m + 1, false));

        int maxLen = 0;
        dp(matrix, memo, memoed, maxLen, n, m);
        return maxLen * maxLen;
    }

    int dp(vector<vector<char>> &matrix, vector<vector<int>> &memo,
           vector<vector<bool>> &memoed, int &maxLen, int i, int j) {

        // find in memo
        if (memoed[i][j]) {
            return memo[i][j];
        }

        // corner case
        if (i == 0 || j == 0) {
            return memo[i][j];
        }

        int currentValue =
            min(dp(matrix, memo, memoed, maxLen, i, j - 1),
                min(dp(matrix, memo, memoed, maxLen, i - 1, j),
                    dp(matrix, memo, memoed, maxLen, i - 1, j - 1))) +
            1;
        if (matrix[i - 1][j - 1] == '1') {
            memo[i][j] += currentValue;
            maxLen = max(maxLen, memo[i][j]);
        }
        memoed[i][j] = true;

        return memo[i][j];
    }
};

int main() {
    Solution solution;
    Solution2 solution2;
    Solution3 solution3;
    vector<vector<char>> matrix;

    matrix = {{'1', '0', '1', '0', '0'},
              {'1', '0', '1', '1', '1'},
              {'1', '1', '1', '1', '1'},
              {'1', '0', '0', '1', '0'}};
    cout << solution.maximalSquare(matrix) << endl;
    cout << solution2.maximalSquare(matrix) << endl;
    cout << solution3.maximalSquare(matrix) << endl;

    matrix = {{'0', '1'}, {'1', '0'}};
    cout << solution.maximalSquare(matrix) << endl;
    cout << solution2.maximalSquare(matrix) << endl;
    cout << solution3.maximalSquare(matrix) << endl;

    matrix = {{'1', '0'}};
    cout << solution.maximalSquare(matrix) << endl;
    cout << solution2.maximalSquare(matrix) << endl;
    cout << solution3.maximalSquare(matrix) << endl;

    return 0;
}