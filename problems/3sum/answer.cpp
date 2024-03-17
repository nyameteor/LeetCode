#include <algorithm>
#include <cassert>
#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    vector<vector<int>> threeSum(vector<int>& nums) {
        vector<vector<int>> res;

        sort(nums.begin(), nums.end());

        int i = 0;
        while (i < nums.size()) {
            int l = i + 1;
            int r = nums.size() - 1;
            while (l < r) {
                int sum = nums[i] + nums[l] + nums[r];
                if (sum < 0) {
                    l++;
                } else if (sum > 0) {
                    r--;
                } else {
                    res.push_back(vector<int>{nums[i], nums[l], nums[r]});
                    while (l + 1 < r && nums[l + 1] == nums[l]) {
                        l++;
                    }
                    while (l < r - 1 && nums[r - 1] == nums[r]) {
                        r--;
                    }
                    l++;
                    r--;
                }
            }
            while (i + 1 < nums.size() && nums[i] == nums[i + 1]) {
                i++;
            }
            i++;
        }
        return res;
    }
};

int main() {
    Solution solution;
    vector<int> nums;
    vector<vector<int>> result;
    vector<vector<int>> answer;

    nums = {-1, 0, 1, 2, -1, -4};
    answer = {{-1, -1, 2}, {-1, 0, 1}};
    result = solution.threeSum(nums);
    assert(result == answer);

    nums = {0, 0, 0, 0, 0, 0, 0, 0, 0};
    answer = {{0, 0, 0}};
    result = solution.threeSum(nums);
    assert(result == answer);

    nums = {0};
    answer = {};
    result = solution.threeSum(nums);
    assert(result == answer);
}