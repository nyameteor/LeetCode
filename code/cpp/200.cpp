#include <iostream>
#include <vector>
using namespace std;

/**
 * Union Find
 */
class Solution {
  private:
    vector<int> s;

  public:
    int numIslands(vector<vector<char>> &grid) {
        int m = grid.size();
        int n = grid[0].size();

        // transform element index
        // note the &. It means we are capturing all of the enclosing variables
        // by reference
        // Refer: https://stackoverflow.com/a/26903867
        auto transform = [&](int i, int j) { return i * n + j; };

        // init disjoint sets, initally root is itself
        s = vector<int>(m * n);
        for (int i = 0; i < m * n; i++) {
            s[i] = i;
        }

        // union the lands
        int cntWater = 0;
        for (int i = 0; i < m; i++) {
            for (int j = 0; j < n; j++) {
                if (grid[i][j] == '0') {
                    cntWater++;
                    continue;
                }
                // only need to check right and down
                if (grid[i][j] == '1' && j + 1 < n && grid[i][j + 1] == '1') {
                    unionSets(transform(i, j), transform(i, j + 1));
                }
                if (grid[i][j] == '1' && i + 1 < m && grid[i + 1][j] == '1') {
                    unionSets(transform(i, j), transform(i + 1, j));
                }
            }
        }

        // number of islands = number of disjoint sets - number of water
        int cntDisjoint = 0;
        for (int i = 0; i < s.size(); i++) {
            if (s[i] == i) {
                cntDisjoint++;
            }
        }
        return cntDisjoint - cntWater;
    }

    int find(int x) {
        if (s[x] != x) {
            s[x] = find(s[x]);
        }
        return s[x];
    }

    void unionSets(int x, int y) {
        int root1 = find(x);
        int root2 = find(y);

        if (root1 == root2) {
            return;
        } else {
            s[root2] = root1;
        }
    }
};

int main() {
    Solution solution;
    vector<vector<char>> grid;

    grid = {
        {'1', '1', '1', '1', '0'},
        {'1', '1', '0', '1', '0'},
        {'0', '0', '0', '0', '0'},
        {'0', '0', '0', '0', '0'},
    };
    cout << solution.numIslands(grid) << endl;

    grid = {
        {'1', '1', '0', '0', '0'},
        {'1', '1', '0', '0', '0'},
        {'0', '0', '1', '0', '0'},
        {'0', '0', '0', '1', '1'},
    };
    cout << solution.numIslands(grid) << endl;

    grid = {{'1', '1', '1'}, {'0', '1', '0'}, {'1', '1', '1'}};
    cout << solution.numIslands(grid) << endl;
}