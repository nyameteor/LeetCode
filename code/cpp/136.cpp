#include <iostream>
#include <unordered_set>
#include <vector>
using namespace std;

class Solution {
  public:
    int singleNumber(vector<int> &nums) {
        unordered_set<int> set;
        for (auto num : nums) {
            if (set.count(num)) {
                set.erase(num);
            } else {
                set.insert(num);
            }
        }
        return *set.begin(); // return first element
    }
};

class Solution2 {
  public:
    int singleNumber(vector<int> &nums) {
        int res = 0;
        for (auto num : nums) {
            res = res ^ num;
        }
        return res;
    }
};

int main() {
    Solution solution;
    vector<int> nums;

    nums = {2, 2, 1};
    cout << solution.singleNumber(nums) << endl;

    nums = {4, 1, 2, 1, 2};
    cout << solution.singleNumber(nums) << endl;

    return 0;
}