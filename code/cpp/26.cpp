#include <iostream>
#include <string>
#include <vector>

using namespace std;

class Solution {
public:
  int removeDuplicates(vector<int> &nums) {
    // edge case
    if (nums.size() == 0)
      return 0;
    if (nums.size() == 1)
      return 1;

    int i = 0;
    int j = 0;
    while (i < nums.size() || j < nums.size()) {
      while (j + 1 < nums.size() && nums[j] == nums[j + 1]) {
        j = j + 1;
      }
      if (j + 1 > nums.size() - 1)
        return i + 1;
      i = i + 1;
      j = j + 1;
      nums[i] = nums[j];
    }
    return i + 1;
  }
};

int main() {
  Solution solution;
  vector<int> nums;

  nums = {1, 1, 2};
  cout << solution.removeDuplicates(nums) << endl;
  for (int i = 0; i < nums.size(); i++) {
    cout << nums[i] << ' ';
  }
  cout << endl;

  nums = {0, 0, 1, 1, 1, 2, 2, 3, 3, 4};
  cout << solution.removeDuplicates(nums) << endl;
  for (int i = 0; i < nums.size(); i++) {
    cout << nums[i] << ' ';
  }
  cout << endl;

  nums = {1, 1};
  cout << solution.removeDuplicates(nums) << endl;
  for (int i = 0; i < nums.size(); i++) {
    cout << nums[i] << ' ';
  }
  cout << endl;

  nums = {1, 2, 2, 2, 2};
  cout << solution.removeDuplicates(nums) << endl;
  for (int i = 0; i < nums.size(); i++) {
    cout << nums[i] << ' ';
  }
  cout << endl;
}