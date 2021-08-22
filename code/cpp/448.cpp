#include <cassert>
#include <cmath>
#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
  vector<int> findDisappearedNumbers(vector<int> &nums) {
    for (int i = 0; i < nums.size(); i++) {
      int key = abs(nums[i]) - 1;
      if (nums[key] > 0)
        nums[key] = -nums[key];
    }

    vector<int> ans;
    for (int i = 0; i < nums.size(); i++) {
      if (nums[i] > 0) {
        ans.push_back(i + 1);
      }
    }
    return ans;
  }
};

int main() {
  Solution solution;
  vector<int> nums;
  vector<int> result;
  vector<int> answer;

  // [5, 6]
  nums = {4, 3, 2, 7, 8, 2, 3, 1};
  answer = {5, 6};
  result = solution.findDisappearedNumbers(nums);
  assert(result == answer);

  // [2]
  nums = {1, 1};
  answer = {2};
  result = solution.findDisappearedNumbers(nums);
  assert(result == answer);
}