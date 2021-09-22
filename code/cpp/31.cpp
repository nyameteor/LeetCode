#include <algorithm>
#include <iostream>
#include <vector>

using namespace std;

class Solution {
  public:
    void nextPermutation(vector<int> &nums) {
        int i, j;
        for (i = nums.size() - 1; i > 0; i--) {
            if (nums[i - 1] < nums[i]) {
                break;
            }
        }

        // corner case, e.g. 4 3 2 1
        if (i == 0) {
            sort(nums.begin(), nums.end());
            return;
        }

        for (j = nums.size() - 1; j >= i; j--) {
            if (nums[j] > nums[i - 1]) {
                break;
            }
        }

        // swap
        int tmp = nums[i - 1];
        nums[i - 1] = nums[j];
        nums[j] = tmp;

        sort(nums.begin() + i, nums.end());
    }
};

void printVector(vector<int> &num) {
    for (int i = 0; i < num.size(); i++) {
        cout << num[i] << " ";
    }
    cout << endl;
}

int main() {
    Solution solution;
    vector<int> nums;

    nums = {1, 2, 3, 4};
    solution.nextPermutation(nums);
    printVector(nums);

    nums = {4, 3, 2, 1};
    solution.nextPermutation(nums);
    printVector(nums);

    nums = {1, 5, 1};
    solution.nextPermutation(nums);
    printVector(nums);

    nums = {1, 1, 1, 1};
    solution.nextPermutation(nums);
    printVector(nums);
}