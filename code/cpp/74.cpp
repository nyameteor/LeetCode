#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    bool searchMatrix(vector<vector<int>> &matrix, int target) {
        // pointer: low, high, mid
        int l, h, m;
        int rowSize = matrix.size(), colSize = matrix[0].size();

        // search in row
        int row = -1;
        l = 0, h = rowSize - 1;
        while (l <= h) {
            m = l + (h - l) / 2;

            if ((matrix[m][0] <= target && m <= rowSize - 2 &&
                 target < matrix[m + 1][0]) ||
                (matrix[m][0] <= target && m == rowSize - 1)) {
                row = m;
                break;
            } else if (matrix[m][0] < target) {
                l = m + 1;
            } else if (matrix[m][0] > target) {
                h = m - 1;
            }
        }

        if (row == -1) {
            return false;
        }

        // search in col
        int col = -1;
        l = 0, h = colSize - 1;
        while (l <= h) {
            m = l + (h - l) / 2;

            if (matrix[row][m] == target) {
                col = m;
                break;
            } else if (matrix[row][m] < target) {
                l = m + 1;
            } else if (matrix[row][m] > target) {
                h = m - 1;
            }
        }

        if (col == -1) {
            return false;
        } else {
            return true;
        }
    }
};

int main() {
    Solution solution;
    vector<vector<int>> matrix;

    matrix = {{1, 3, 5, 7}, {10, 11, 16, 20}, {23, 30, 34, 60}};
    cout << solution.searchMatrix(matrix, 3) << endl;
    cout << solution.searchMatrix(matrix, 13) << endl;
    cout << solution.searchMatrix(matrix, 30) << endl;
}