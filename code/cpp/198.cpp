#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    int rob(vector<int> &nums) {
        int n = nums.size();
        vector<int> memo(n, -1);
        return dfs(n, nums, memo);
    }

    int dfs(int n, vector<int> &nums, vector<int> &memo) {
        if (n <= 0) {
            return 0;
        } else if (n == 1) {
            return nums[n - 1];
        }

        if (memo[n - 1] != -1) {
            return memo[n - 1];
        }

        memo[n - 1] =
            max(dfs(n - 2, nums, memo) + nums[n - 1], dfs(n - 1, nums, memo));
        return memo[n - 1];
    }
};

int main() {
    Solution solution;
    vector<int> nums;

    // 4
    nums = {1, 2, 3, 1};
    cout << solution.rob(nums) << endl;

    // 12
    nums = {2, 7, 9, 3, 1};
    cout << solution.rob(nums) << endl;

    return 0;
}