#include <algorithm>
#include <cmath>
#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    int threeSumClosest(vector<int> &nums, int target) {
        // edge case
        if (nums.size() == 3) {
            return nums[0] + nums[1] + nums[2];
        }

        // sort beforce using two pointers
        sort(nums.begin(), nums.end());

        // big enough initial value
        int cloest_sum = 999999;

        for (int i = 0; i < nums.size() - 2; i++) {
            // left pointer
            int l = i + 1;
            // right pointer
            int r = nums.size() - 1;

            while (l < r) {
                int sum = nums[i] + nums[l] + nums[r];
                int diff = sum - target;
                int distance = abs(sum - target);
                int cloest_distance = abs(cloest_sum - target);

                if (diff < 0) {
                    // update ans
                    if (cloest_distance > distance)
                        cloest_sum = sum;
                    l++;
                } else if (diff > 0) {
                    if (cloest_distance > distance)
                        cloest_sum = sum;
                    r--;
                } else if (diff == 0) {
                    return target;
                }
            }
        }
        return cloest_sum;
    }
};

int main() {
    Solution solution;

    // 3 <= nums.length <= 1000
    // -1000 <= nums[i] <= 1000
    // -10^4 <= target <= 10^4
    vector<int> nums;

    // 2
    nums = {-1, 2, 1, -4};
    cout << solution.threeSumClosest(nums, 1) << endl;

    // 0
    nums = {0, 2, 1, -3};
    cout << solution.threeSumClosest(nums, 1) << endl;

    // 3
    nums = {1, 1, 1, 1};
    cout << solution.threeSumClosest(nums, -100) << endl;

    // 0
    nums = {0, 0, 0};
    cout << solution.threeSumClosest(nums, 1) << endl;
}