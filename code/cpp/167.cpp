#include <cstdio>
#include <iostream>
#include <set>
#include <stack>
#include <string>
#include <vector>

using namespace std;

vector<int> twoSum(vector<int> &numbers, int target) {
    int l = 0, r = numbers.size() - 1;
    while (l < numbers.size() && r >= 0 && l < r) {
        if (numbers[l] + numbers[r] == target) {
            break;
        } else if (numbers[l] + numbers[r] < target) {
            l++;
        } else if (numbers[l] + numbers[r] > target) {
            r--;
        }
    }
    return vector<int>{l + 1, r + 1};
}

int main() {
    int target = 9;
    vector<int> nums = {2, 7, 11, 15};
    vector<int> ans = twoSum(nums, target);
    for (int i = 0; i < ans.size(); i++) {
        printf("%d ", ans[i]);
    }

    return 0;
}