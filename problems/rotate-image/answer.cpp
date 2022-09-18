#include <vector>
using namespace std;

/**
 * Rotate Groups of Four Cells
 */
class Solution {
public:
    void rotate(vector<vector<int>> &matrix) {
        int N = size(matrix);
        int n = N - 1;

        for (int i = 0; i < N / 2; i++) {
            for (int j = i; j < n - i; j++) {
                int tmp = matrix[i][j];
                matrix[i][j] = matrix[n - j][i];
                matrix[n - j][i] = matrix[n - i][n - j];
                matrix[n - i][n - j] = matrix[j][n - i];
                matrix[j][n - i] = tmp;
            }
        }
    }
};

/**
 * Reverse on Diagonal and then Reverse Left to Right
 */
class Solution2 {
public:
    void rotate(vector<vector<int>> &matrix) {
        transpose(matrix);
        reflect(matrix);
    }

    void transpose(vector<vector<int>> &matrix) {
        int N = size(matrix);
        for (int i = 0; i < N; i++) {
            for (int j = i; j < N; j++) {
                int tmp = matrix[i][j];
                matrix[i][j] = matrix[j][i];
                matrix[j][i] = tmp;
            }
        }
    }

    void reflect(vector<vector<int>> &matrix) {
        int N = size(matrix);
        for (int i = 0; i < N; i++) {
            for (int j = 0; j < N / 2; j++) {
                int tmp = matrix[i][j];
                matrix[i][j] = matrix[i][N - 1 - j];
                matrix[i][N - 1 - j] = tmp;
            }
        }
    }
};