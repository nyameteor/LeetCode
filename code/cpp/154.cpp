#include <cmath>
#include <iostream>
#include <vector>
using namespace std;

/**
 * Binary Search
 *
 * Time: O(logn) ~ O(n)
 */
class Solution {
  public:
    int findMin(vector<int> &nums) {
        int l = 0, r = nums.size() - 1, m;
        int res = 0;

        while (l < r) {
            m = floor((l + r) / 2);
            if (nums[m] > nums[r]) {
                // [l, m] is ordered, search [m, r]
                l = m + 1;
            } else if (nums[l] > nums[m]) {
                // [m, r] is ordered, search [l, m]
                r = m;
            } else {
                // Case: 2 1 2 2, 2 2 1 2
                // In this condition, we can't decide which interval is ordered,
                // so we can let r--, this step will reduce the interval and
                // not lose minimum or make index out of bounds
                r--;
            }
        }

        return nums[l];
    }
};

int main() {
    Solution solution;
    vector<int> nums;

    nums = {4, 5, 5, 0, 0, 1, 1, 1, 2, 4};
    cout << solution.findMin(nums) << endl;

    nums = {3, 3, 1, 3};
    cout << solution.findMin(nums) << endl;

    nums = {1, 3, 3};
    cout << solution.findMin(nums) << endl;

    nums = {3, 3, 4};
    cout << solution.findMin(nums) << endl;

    nums = {0, 0, 0, 0, 0};
    cout << solution.findMin(nums) << endl;

    nums = {1};
    cout << solution.findMin(nums) << endl;

    return 0;
}