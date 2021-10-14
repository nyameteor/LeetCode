#include <algorithm>
#include <iostream>
#include <vector>
using namespace std;

/**
 * Backtracking, only backtrack at path
 */
class Solution {
  public:
    vector<vector<int>> permute(vector<int> &nums) {
        vector<vector<int>> paths;
        vector<int> path;
        dfs(nums, path, paths);
        return paths;
    }

    void dfs(vector<int> &nums, vector<int> &path, vector<vector<int>> &paths) {
        if (nums.size() == 0) {
            paths.push_back(path);
            return;
        }
        for (int i = 0; i < nums.size(); i++) {
            path.push_back(nums[i]);
            // copy a temporary array to store unvisited nums
            vector<int> tmpNums = nums;
            tmpNums.erase(tmpNums.begin() + i);

            dfs(tmpNums, path, paths);

            path.pop_back(); // backtrack at path
        }
    }
};

/**
 * Backtracking, backtrack at path and branch
 *
 * It's a straightforward solution with good performance
 */
class Solution2 {
  public:
    vector<vector<int>> permute(vector<int> &nums) {
        vector<vector<int>> paths;
        vector<int> path;
        vector<bool> visited(nums.size(), false);
        dfs(nums, path, visited, paths);
        return paths;
    }

    void dfs(vector<int> &nums, vector<int> &path, vector<bool> &visited,
             vector<vector<int>> &paths) {
        if (path.size() == nums.size()) {
            paths.push_back(path);
            return;
        }
        for (int i = 0; i < nums.size(); i++) {
            // using a array to check if nums are visited
            if (!visited[i]) {
                visited[i] = true;
                path.push_back(nums[i]);

                dfs(nums, path, visited, paths);

                path.pop_back();    // backtrack at path
                visited[i] = false; // backtrack at branch
            }
        }
    }
};

void printResult(vector<vector<int>> res) {
    for (auto x : res) {
        for (auto y : x) {
            cout << y << ' ';
        }
    }
    cout << endl;
}

int main() {
    Solution solution;
    Solution2 solution2;
    vector<int> nums;

    nums = {1, 2, 3};
    printResult(solution.permute(nums));
    printResult(solution2.permute(nums));

    nums = {0, 1};
    printResult(solution.permute(nums));
    printResult(solution2.permute(nums));

    nums = {1};
    printResult(solution.permute(nums));
    printResult(solution2.permute(nums));
}