#include <cstdint>
#include <queue>
#include <vector>
using namespace std;

class Solution {
public:
    int shortestPathBinaryMatrix(vector<vector<int>> &grid) {

        // edge case
        if (grid[0][0] == 1) {
            return -1;
        }

        int N = size(grid);
        int level = 0;
        vector<vector<bool>> visited(N, vector<bool>(N, false));

        queue<pair<int, int>> q; // {row, col}
        q.push({0, 0});
        visited[0][0] = true;
        while (!empty(q)) {

            level++;

            int len = size(q);
            while (len--) {
                auto cur = q.front();
                q.pop();

                int i = cur.first;
                int j = cur.second;

                if (i == N - 1 && j == N - 1) {
                    return level;
                }

                for (int k = i - 1; k <= i + 1; k++) {
                    for (int l = j - 1; l <= j + 1; l++) {
                        if (isValid(grid, N, k, l, visited)) {
                            q.push({k, l});
                            visited[k][l] = true;
                        }
                    }
                }
            }
        }

        return -1;
    }

    int isValid(vector<vector<int>> &grid, int N, int i, int j,
                vector<vector<bool>> &visited) {
        return (i >= 0 && i < N) && (j >= 0 && j < N) && grid[i][j] == 0 &&
               visited[i][j] == false;
    }
};