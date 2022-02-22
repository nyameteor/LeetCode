#include <algorithm>
#include <iostream>
#include <unordered_map>
#include <vector>
using namespace std;

/**
 * Use Hashmap to count element occurrences efficiently.
 */
class Solution {
  public:
    int majorityElement(vector<int> &nums) {
        unordered_map<int, int> freq;
        for (int i = 0; i < size(nums); i++) {
            freq[nums[i]]++;
        }
        vector<pair<int, int>> elems(freq.begin(), freq.end());
        int N = size(nums);
        for (int i = 0; i < size(elems); i++) {
            if (elems[i].second > N / 2) {
                return elems[i].first;
            }
        }

        return -1;
    }
};

/**
 * Majority vote algorithm
 *
 * Refer:
 * https://en.wikipedia.org/wiki/Boyer%E2%80%93Moore_majority_vote_algorithm
 */
class Solution2 {
  public:
    int majorityElement(vector<int> &nums) {
        int count = 0;
        int candidate = -1;

        for (int num : nums) {
            if (count == 0) {
                candidate = num;
            }
            if (num == candidate) {
                count++;
            } else {
                count--;
            }
        }

        return candidate;
    }
};

int main() {
    Solution solution;
    vector<int> nums;

    nums = {3, 2, 3};
    cout << solution.majorityElement(nums) << endl;

    nums = {2, 2, 1, 1, 1, 2, 2};
    cout << solution.majorityElement(nums) << endl;

    return 0;
}