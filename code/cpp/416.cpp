#include <algorithm>
#include <iostream>
#include <unordered_map>
#include <vector>
using namespace std;

/**
 * Recursion
 *
 * Time: O(2^n), Space: O(1)
 */
class Solution {
  public:
    bool canPartition(vector<int> &nums) {
        bool res = false;
        dfs(nums, 0, 0, 0, res);
        return res;
    }

    void dfs(vector<int> &nums, int sum1, int sum2, int i, bool &res) {
        if (i == nums.size()) {
            if (sum1 == sum2) {
                res = true;
                return;
            } else {
                return;
            }
        }
        dfs(nums, sum1 + nums[i], sum2, i + 1, res);
        dfs(nums, sum1, sum2 + nums[i], i + 1, res);
    }
};

/**
 * Recursion
 *
 * Another implementation of recursion.
 */
class Solution1 {
  public:
    bool canPartition(vector<int> &nums) { return dfs(nums, 0, 0, 0); }

    bool dfs(vector<int> &nums, int sum1, int sum2, int i) {
        if (i == nums.size()) {
            if (sum1 == sum2) {
                return true;
            } else {
                return false;
            }
        }
        return dfs(nums, sum1 + nums[i], sum2, i + 1) || dfs(nums, sum1, sum2 + nums[i], i + 1);
    }
};

/**
 * Recursion + Memoization
 */
class Solution2 {
  public:
    bool canPartition(vector<int> &nums) {
        unordered_map<int, bool> memo;
        return dfs(nums, 0, 0, 0, memo);
    }

    bool dfs(vector<int> &nums, int sum1, int sum2, int i, unordered_map<int, bool> &memo) {
        if (memo.find(sum1) != end(memo)) {
            return memo[sum1];
        }
        if (i == nums.size()) {
            if (sum1 == sum2) {
                memo[sum1] = true;
            } else {
                memo[sum1] = false;
            }
            return memo[sum1];
        }
        memo[sum1] = dfs(nums, sum1 + nums[i], sum2, i + 1, memo) ||
                     dfs(nums, sum1, sum2 + nums[i], i + 1, memo);
        return memo[sum1];
    }
};

int main() {
    Solution2 solution;
    vector<int> nums;

    nums = {1, 5, 11, 5};
    cout << solution.canPartition(nums) << endl;

    nums = {1, 2, 3, 5};
    cout << solution.canPartition(nums) << endl;

    return 0;
}