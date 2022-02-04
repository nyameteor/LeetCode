#include <iostream>
#include <unordered_map>
#include <unordered_set>
#include <vector>
using namespace std;

class Solution {
  public:
    int fourSumCount(vector<int> &nums1, vector<int> &nums2, vector<int> &nums3,
                     vector<int> &nums4) {
        unordered_map<int, int> m;
        for (auto a : nums1) {
            for (auto b : nums2) {
                m[a + b]++;
            }
        }

        int count = 0;
        for (auto c : nums3) {
            for (auto d : nums4) {
                if (m.find(-c - d) != end(m)) {
                    count += m[-c - d];
                }
            }
        }
        return count;
    }
};

int main() {
    Solution solution;
    vector<int> nums1, nums2, nums3, nums4;

    nums1 = {1, 2};
    nums2 = {-2, -1};
    nums3 = {-1, 2};
    nums4 = {0, 2};
    cout << solution.fourSumCount(nums1, nums2, nums3, nums4) << endl;

    return 0;
}