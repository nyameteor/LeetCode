#include <iostream>
#include <vector>
using namespace std;

/**
 * Dutch Flag Algorithm by Dijkstra
 */
class Solution {
  public:
    void sortColors(vector<int> &nums) {
        int i = 0, j = 0, k = nums.size() - 1;

        while (j <= k) {
            if (nums[j] < 1) {
                swap(nums[i], nums[j]);
                i++;
                j++;
            } else if (nums[j] > 1) {
                swap(nums[j], nums[k]);
                k--;
            } else {
                j++;
            }
        }
    }

    void swap(int &a, int &b) {
        int temp = a;
        a = b;
        b = temp;
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
    vector<int> nums;

    nums = {2, 0, 2, 1, 1, 0};
    solution.sortColors(nums);
    printVector(nums);

    return 0;
}