#include <iostream>
#include <unordered_map>
#include <vector>
using namespace std;

/**
 * Idea: find target combination in k-combinations
 * C(n,k) = P(n,k)/k! = n!/k!(n-k)! (k = 2, n = sizeof(nums))
 *
 * Refer: https://en.wikipedia.org/wiki/Combination
 */

/**
 * Use nested loops to find target combination in k-combinations.
 */
class Solution {
public:
    vector<int> twoSum(vector<int> &nums, int target) {
        vector<int> res(2, 0);

        for (int i = 0; i < nums.size() - 1; i++) {
            for (int j = i + 1; j < nums.size(); j++) {
                if (nums[i] + nums[j] == target) {
                    res[0] = i;
                    res[1] = j;
                    return res;
                }
            }
        }

        return res;
    }
};

/**
 * Use map to find target combination:
 *  1. Traverse the input array one by one
 *  2. we checking if `key = target - nums[i]` is in the map:
 *      if we found it, it means we found the combination;
 *      if not, put `(nums[i], i)` into the map
 *
 * This implementation is a bit of tricky, but not difficult to understand
 */
class Solution2 {
public:
    vector<int> twoSum(vector<int> &nums, int target) {
        vector<int> res(2, 0);
        unordered_map<int, int> m;

        for (int i = 0; i < nums.size(); i++) {
            int second = target - nums[i];
            if (m.find(second) != end(m)) {
                res[0] = m.at(second);
                res[1] = i;
                break;
            } else {
                m.insert(make_pair(nums[i], i));
            }
        }
        return res;
    }
};

void printVector(vector<int> v) {
    for (int x : v) {
        cout << x << ' ';
    }
    cout << endl;
}

int main() {
    Solution solution;
    Solution2 solution2;
    vector<int> nums;

    nums = {2, 7, 11, 15};
    printVector(solution.twoSum(nums, 9));
    printVector(solution2.twoSum(nums, 9));

    nums = {3, 2, 4};
    printVector(solution.twoSum(nums, 6));
    printVector(solution2.twoSum(nums, 6));

    nums = {2, 3, 3};
    printVector(solution.twoSum(nums, 6));
    printVector(solution2.twoSum(nums, 6));

    return 0;
}