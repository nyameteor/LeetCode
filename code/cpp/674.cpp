#include <iostream>
#include <vector>
using namespace std;

/**
 * Sliding Window
 *
 * Idea: 
 */
class Solution {
  public:
    int findLengthOfLCIS(vector<int> &nums) {
        int res = 0;

        for (int l = 0, r = 0; r < nums.size(); r++) {
            if (r > 0 && nums[r - 1] >= nums[r]) {
                l = r;
            }
            res = max(res, r - l + 1);
        }

        return res;
    }
};

int main() {
    Solution solution;
    vector<int> nums;

    nums = {1, 3, 5, 4, 7};
    cout << solution.findLengthOfLCIS(nums) << endl;

    nums = {2, 2, 2, 2, 2};
    cout << solution.findLengthOfLCIS(nums) << endl;

    nums = {1, 2, 3, 4, 5};
    cout << solution.findLengthOfLCIS(nums) << endl;
}