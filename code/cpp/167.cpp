#include <iostream>
#include <vector>
using namespace std;

class Solution {
  public:
    vector<int> twoSum(vector<int> &numbers, int target) {
        int l = 0, r = numbers.size() - 1;
        while (l < r) {
            if (numbers[l] + numbers[r] == target) {
                break;
            } else if (numbers[l] + numbers[r] < target) {
                l++;
            } else {
                r--;
            }
        }
        return vector<int>{l + 1, r + 1};
    }
};

void printResult(vector<int> v) {
    for (auto x : v) {
        cout << x << ' ';
    }
    cout << endl;
}

int main() {
    Solution solution;
    vector<int> nums;

    nums = {2, 7, 11, 15};
    printResult(solution.twoSum(nums, 9));
    nums = {5, 25, 75};
    printResult(solution.twoSum(nums, 100));
    nums = {2, 3, 4};
    printResult(solution.twoSum(nums, 6));

    return 0;
}