#include <iostream>
#include <vector>

using namespace std;

class Solution {
  public:
    int search(vector<int> &nums, int target) {
        int left = 0, right = nums.size() - 1;

        while (left <= right) {
            int mid = left + (right - left) / 2;

            if (target == nums[mid]) {
                return mid;
            }

            if (nums[left] <= nums[mid]) {
                // [left, mid] is in order
                if (nums[left] <= target && target < nums[mid]) {
                    right = mid - 1;
                } else {
                    left = mid + 1;
                }
            } else {
                // [mid + 1, right] is in order
                if (nums[mid] < target && target <= nums[right]) {
                    left = mid + 1;
                } else {
                    right = mid - 1;
                }
            }
        }

        return -1;
    }
};

int main() {
    Solution solution;
    vector<int> nums;

    nums = {4, 5, 6, 7, 0, 1, 2};
    cout << solution.search(nums, 4) << endl;
    cout << solution.search(nums, 1) << endl;
    cout << solution.search(nums, 3) << endl;

    nums = {1, 2};
    cout << solution.search(nums, 1) << endl;

    nums = {1};
    cout << solution.search(nums, 1) << endl;
}