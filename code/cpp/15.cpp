#include <algorithm>
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

    for (int i = 0; i < nums.size() - 2; i++) {
      // left pointer
      int l = i + 1;
      // right pointer
      int r = nums.size() - 1;

      while (l < r) {
        int sum = nums[i] + nums[l] + nums[r];
        if (sum < 0) {
          l++;
        } else if (sum > 0) {
          r--;
        } else if (sum == 0) {
          vector<int> triplet = {nums[i], nums[l], nums[r]};
          ans.push_back(triplet);
          // optimize: skip the same value
          while (l < r && nums[l] == nums[l + 1])
            l = l + 1;
          while (l < r && nums[r] == nums[r - 1])
            r = r - 1;

          l++;
          r--;
        }
      }
      // optimize: skip the same value
      while (i < nums.size() - 3 && nums[i] == nums[i + 1])
        i = i + 1;
    }
    return ans;
  }
};

int main() {
  Solution solution;
  vector<int> nums;
  vector<vector<int>> result;

  // [[-1,-1,2],[-1,0,1]]
  nums = {-1, 0, 1, 2, -1, -4};
  result = solution.threeSum(nums);
  for (int i = 0; i < result.size(); i++) {
    for (int j = 0; j < result[i].size(); j++) {
      cout << result[i][j] << " ";
    }
    cout << endl;
  }

  // [[0,0,0]]
  nums = {0, 0, 0, 0, 0, 0, 0, 0, 0};
  result = solution.threeSum(nums);
  for (int i = 0; i < result.size(); i++) {
    for (int j = 0; j < result[i].size(); j++) {
      cout << result[i][j] << " ";
    }
    cout << endl;
  }

  // []
  nums = {0};
  result = solution.threeSum(nums);
  for (int i = 0; i < result.size(); i++) {
    for (int j = 0; j < result[i].size(); j++) {
      cout << result[i][j] << " ";
    }
    cout << endl;
  }
}