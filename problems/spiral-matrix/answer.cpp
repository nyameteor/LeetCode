#include <iostream>
#include <queue>
#include <vector>

using namespace std;

class Solution {
public:
    vector<int> spiralOrder(vector<vector<int>> &matrix) {
        int c = 0, r = 0;
        int m = matrix.size();
        int n = matrix[0].size();

        vector<int> res;
        while (r < m && c < n) {
            // top: first row
            for (int i = c; i < n; i++) {
                res.push_back(matrix[c][i]);
            }
            r++;

            // right: last column
            for (int i = r; i < m; i++) {
                res.push_back(matrix[i][n - 1]);
            }
            n--;

            // bottom: last row
            for (int i = n - 1; i >= c && r < m; i--) {
                res.push_back(matrix[m - 1][i]);
            }
            m--;

            // left: first column
            for (int i = m - 1; i >= r && c < n; i--) {
                res.push_back(matrix[i][c]);
            }
            c++;
        }

        return res;
    }
};

void printVector(vector<int> num) {
    for (int x : num) {
        cout << x << " ";
    }
    cout << endl;
}

int main() {
    Solution solution;
    vector<vector<int>> matrix;

    matrix = {{1, 2, 3}, {4, 5, 6}, {7, 8, 9}};
    printVector(solution.spiralOrder(matrix));

    matrix = {{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}};
    printVector(solution.spiralOrder(matrix));

    matrix = {{1, 2, 3}};
    printVector(solution.spiralOrder(matrix));

    matrix = {{1}, {2}, {3}};
    printVector(solution.spiralOrder(matrix));

    return 0;
}