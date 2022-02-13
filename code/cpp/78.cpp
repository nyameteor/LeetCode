#include <iostream>
#include <unordered_set>
#include <vector>
using namespace std;

/**
 * Depth-First Search (with Backtracking)
 *
 * Time: O(2^n)
 * Space: O(1)
 */
class Solution {
  public:
    vector<vector<int>> subsets(vector<int> &nums) {
        vector<vector<int>> res;
        vector<int> cur;

        dfs(res, cur, nums, 0);

        return res;
    }

    void dfs(vector<vector<int>> &res, vector<int> &cur, vector<int> nums,
             int i) {
        if (i == size(nums)) {
            res.push_back(cur);
            return;
        }
        cur.push_back(nums[i]);
        dfs(res, cur, nums, i + 1);

        cur.pop_back();
        dfs(res, cur, nums, i + 1);
    }
};

/**
 * Depth-First Search
 *
 * Time: O(2^n)
 * Space: O(2^n)
 */
class Solution2 {
  public:
    vector<vector<int>> subsets(vector<int> &nums) {
        vector<vector<int>> res;
        vector<int> subset;

        dfs(res, subset, nums, 0);

        return res;
    }

    void dfs(vector<vector<int>> &res, vector<int> cur, vector<int> nums,
             int i) {
        if (i == size(nums)) {
            res.push_back(cur);
            return;
        }
        dfs(res, cur, nums, i + 1);

        vector<int> cur_copy = cur;
        cur_copy.push_back(nums[i]);
        dfs(res, cur_copy, nums, i + 1);
    }
};

/**
 * Cascading
 */
class Solution3 {
  public:
    vector<vector<int>> subsets(vector<int> &nums) {
        vector<vector<int>> res;
        vector<int> subset;
        res.push_back(subset);
        for (auto num : nums) {
            int N = size(res);
            for (int i = 0; i < N; i++) {
                vector<int> cur = res[i];
                cur.push_back(num);
                res.push_back(cur);
            }
        }

        return res;
    }
};

void printResult(vector<vector<int>> vv) {
    for (auto v : vv) {
        cout << '[';
        for (int i = 0; i < size(v); i++) {
            if (i != size(v) - 1) {
                cout << v[i] << ',';
            } else {
                cout << v[i];
            }
        }
        cout << ']';
    }
    cout << endl;
}

int main() {
    Solution solution;
    vector<int> nums;

    nums = {1, 2, 3};
    printResult(solution.subsets(nums));

    nums = {0};
    printResult(solution.subsets(nums));

    return 0;
}