#include <iostream>
#include <string>
#include <vector>
using namespace std;

class Solution {
public:
    vector<vector<string>> solveNQueens(int n) {
        vector<pair<int, int>> positions;
        vector<vector<string>> res;
        for (int j = 0; j < n; j++) {
            solver(res, positions, n, 0, j);
        }

        return res;
    }

    void solver(vector<vector<string>> &res, vector<pair<int, int>> &positions,
                int n, int i, int j) {
        if (i >= n) {
            return;
        }

        if (!isPosValid(positions, i, j)) {
            return;
        }

        positions.push_back({i, j});

        if (size(positions) == n) {
            res.push_back(convertRes(positions));
        }

        for (int k = 0; k < n; k++) {
            solver(res, positions, n, i + 1, k);
        }

        positions.pop_back();
    }

    bool isPosValid(vector<pair<int, int>> &positions, int i, int j) {
        for (auto pos : positions) {
            int x = pos.first;
            int y = pos.second;

            // is same row
            if (i == x) {
                return false;
            }

            // is same column
            if (j == y) {
                return false;
            }

            // is same diagonal
            if (abs(i - x) == abs(j - y)) {
                return false;
            }
        }

        return true;
    }

    vector<string> convertRes(vector<pair<int, int>> &positions) {
        int n = size(positions);

        vector<string> res(n);
        for (auto pos : positions) {
            int i = pos.first;
            int j = pos.second;

            string s = "";
            for (int k = 0; k < n; k++) {
                if (k == j) {
                    s.push_back('Q');
                } else {
                    s.push_back('.');
                }
            }

            res[i] = s;
        }

        return res;
    }
};

void printResult(vector<vector<string>> vv) {
    for (auto v : vv) {
        for (auto str : v) {
            cout << str << ' ';
        }
        cout << endl;
    }
}

int main() {
    Solution solution;

    printResult(solution.solveNQueens(4));

    printResult(solution.solveNQueens(8));
}