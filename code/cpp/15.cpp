#include <algorithm>
#include <cassert>
#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    vector<vector<int>> threeSum(vector<int> &nums) {
        vector<vector<int>> ans;

        // sort nums to use two pointers
        sort(nums.begin(), nums.end());

        // edge case
        if (nums.size() < 3) {
            return ans;
        }

        // Use `while` instead of `for` to make
        // the structure of loop and sub-loop more similar
        int i = 0;
        while (i < nums.size() - 2) {
            // left pointer
            int l = i + 1;
            // right pointer
            int r = nums.size() - 1;

            while (l < r) {
                int sum = nums[i] + nums[l] + nums[r];
                if (sum < 0) {
                    while (l < r && nums[l] == nums[l + 1])
                        l++;
                    l++;
                } else if (sum > 0) {
                    while (l < r && nums[r] == nums[r - 1])
                        r--;
                    r--;
                } else if (sum == 0) {
                    vector<int> triplet = {nums[i], nums[l], nums[r]};
                    ans.push_back(triplet);
                    // optimize: skip same element
                    while (l < r && nums[l] == nums[l + 1])
                        l++;
                    while (l < r && nums[r] == nums[r - 1])
                        r--;

                    l++;
                    r--;
                }
            }
            while (i < nums.size() - 3 && nums[i] == nums[i + 1])
                i++;
            i++;
        }
        return ans;
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