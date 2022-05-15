#include <iostream>
#include <vector>

using namespace std;

/* dynamic programming */
class Solution {
public:
    int lengthOfLIS(vector<int> &nums) {
        int maxLen = 0;

        // tabulation
        vector<int> table(nums.size(), 1);
        for (int i = 0; i < nums.size(); i++) {
            for (int j = 0; j < i; j++) {
                if (nums[j] < nums[i]) {
                    table[i] = max(table[i], table[j] + 1);
                }
            }
            maxLen = max(maxLen, table[i]);
        }

        return maxLen;
    }
};

int main() {
    Solution solution;
    vector<int> nums;

    // 4
    nums = {10, 9, 2, 5, 3, 7, 101, 18};
    cout << solution.lengthOfLIS(nums) << endl;

    // 4
    nums = {0, 1, 0, 3, 2, 3};
    cout << solution.lengthOfLIS(nums) << endl;

    // 1
    nums = {7, 7, 7, 7, 7, 7, 7};
    cout << solution.lengthOfLIS(nums) << endl;

    // 3
    nums = {4, 10, 4, 3, 8, 9};
    cout << solution.lengthOfLIS(nums) << endl;

    // 6
    nums = {1, 3, 6, 7, 9, 4, 10, 5, 6};
    cout << solution.lengthOfLIS(nums) << endl;

    return 0;
}