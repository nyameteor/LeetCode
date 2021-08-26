#include <cassert>
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
    int result;

    nums = {1, 1, 2};
    assert(solution.removeDuplicates(nums) == 2);

    nums = {0, 0, 1, 1, 1, 2, 2, 3, 3, 4};
    assert(solution.removeDuplicates(nums) == 5);

    nums = {1, 1};
    assert(solution.removeDuplicates(nums) == 1);

    nums = {1, 2, 2, 2, 2};
    assert(solution.removeDuplicates(nums) == 2);
}