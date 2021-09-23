#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

class Solution {
  public:
    int maxSubArray(vector<int> &nums) {
        int maxSum = INT32_MIN;
        int sum = 0;
        for (int i = 0; i < nums.size(); i++) {
            sum = sum + nums[i];
            maxSum = max(maxSum, sum);

            // greedy, discard continuous sum
            if (sum < 0) {
                sum = 0;
            }
        }

        return maxSum;
    }
};

int main() {
    Solution solution;
    vector<int> nums;

    // 6
    nums = {-2, 1, -3, 4, -1, 2, 1, -5, 4};
    cout << solution.maxSubArray(nums) << endl;

    // 1
    nums = {1};
    cout << solution.maxSubArray(nums) << endl;

    // 23
    nums = {5, 4, -1, 7, 8};
    cout << solution.maxSubArray(nums) << endl;
}