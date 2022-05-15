#include <iostream>
#include <queue>
#include <utility>
#include <vector>
using namespace std;

/**
 * Breadth-First Search
 *
 * Idea:
 *  1. Push the positions of all initially rotten oranges to the queue; Count
 *  all fresh orange at start;
 *  2. Traversing from rotten oranges, pop queue head, go to it's 4 adjacent
 *  posisionts and if the orange is fresh, then rot it and decrease the count of
 *  fresh oranges;
 *  3. Increase the `time` after repeating step 2, until queue is empty;
 *  4. If there are fresh orange remain then return -1, otherwise retur `time`.
 *
 * Optimized
 */
class Solution {
public:
    int orangesRotting(vector<vector<int>> &grid) {
        queue<pair<int, int>> q;
        int fresh = 0;

        int m = grid.size();
        int n = grid[0].size();
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == 2) {
                    q.push(make_pair(i, j));
                } else if (grid[i][j] == 1) {
                    fresh++;
                }
            }
        }

        int time = 0;
        // top, bottom, right, left
        vector<pair<int, int>> dirs = {{1, 0}, {-1, 0}, {0, 1}, {0, -1}};
        while (!q.empty()) {
            int size = q.size();
            // has made fresh orange become rotten
            bool hasAdded = false;
            for (int i = 0; i < size; i++) {
                auto cur = q.front();
                q.pop();

                for (auto dir : dirs) {
                    int r = cur.first + dir.first;
                    int c = cur.second + dir.second;
                    if (r >= 0 && r < m && c >= 0 && c < n && grid[r][c] == 1) {
                        grid[r][c] = 2;
                        q.push({r, c});
                        fresh--;
                        hasAdded = true;
                    }
                }
            }
            if (hasAdded) {
                time++;
            }
        }

        if (fresh > 0) {
            return -1;
        } else {
            return time;
        }
    }
};

/**
 * Same ieda, but not optimized
 */
class Solution2 {
private:
    struct cell {
        int row;
        int col;
        cell(int r, int c) : row(r), col(c) {}
    };

public:
    int orangesRotting(vector<vector<int>> &grid) {
        queue<cell *> q;
        int fresh = 0;

        int m = grid.size();
        int n = grid[0].size();
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == 2) {
                    q.push(new cell(i, j));
                } else if (grid[i][j] == 1) {
                    fresh++;
                }
            }
        }

        int time = 0;
        while (!q.empty()) {
            int size = q.size();
            bool hasAdded = false;
            for (int i = 0; i < size; i++) {
                cell *c = q.front();
                q.pop();

                // top
                if (c->row - 1 >= 0 && grid[c->row - 1][c->col] == 1) {
                    grid[c->row - 1][c->col] = 2;
                    q.push(new cell(c->row - 1, c->col));
                    hasAdded = true;
                    fresh--;
                }

                // bottom
                if (c->row + 1 < m && grid[c->row + 1][c->col] == 1) {
                    grid[c->row + 1][c->col] = 2;
                    q.push(new cell(c->row + 1, c->col));
                    hasAdded = true;
                    fresh--;
                }

                // left
                if (c->col - 1 >= 0 && grid[c->row][c->col - 1] == 1) {
                    grid[c->row][c->col - 1] = 2;
                    q.push(new cell(c->row, c->col - 1));
                    hasAdded = true;
                    fresh--;
                }

                // right
                if (c->col + 1 < n && grid[c->row][c->col + 1] == 1) {
                    grid[c->row][c->col + 1] = 2;
                    q.push(new cell(c->row, c->col + 1));
                    hasAdded = true;
                    fresh--;
                }
            }
            if (hasAdded) {
                time++;
            }
        }

        if (fresh > 0) {
            return -1;
        }

        return time;
    }
};

int main() {
    Solution solution;
    vector<vector<int>> grid;

    grid = {{2, 1, 1}, {1, 1, 0}, {0, 1, 1}};
    cout << solution.orangesRotting(grid) << endl;

    grid = {{2, 1, 1}, {0, 1, 1}, {1, 0, 1}};
    cout << solution.orangesRotting(grid) << endl;

    grid = {{2, 0}};
    cout << solution.orangesRotting(grid) << endl;

    grid = {{1, 2}};
    cout << solution.orangesRotting(grid) << endl;
}