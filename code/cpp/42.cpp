#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    // Two Pointer
    // Time Limit Exceeded
    int trap(vector<int> &height) {
        // corner case
        if (height.size() <= 2) {
            return 0;
        }

        int res = 0;
        // head and tail cannot trap water
        for (int i = 1; i < height.size() - 1; i++) {
            int lmax = 0, rmax = 0;

            // find highest column in left
            for (int j = i - 1; j >= 0; j--) {
                lmax = max(lmax, height[j]);
            }

            // find highest column in right
            for (int k = i + 1; k < height.size(); k++) {
                rmax = max(rmax, height[k]);
            }

            int h = min(lmax, rmax) - height[i];
            if (h > 0) {
                res += h;
            }
        }
        return res;
    }

    // Dynamic Programming
    int trap2(vector<int> &height) {
        int size = height.size();

        // corner case
        if (size <= 2) {
            return 0;
        }

        vector<int> maxLeft(size), maxRight(size);
        // tabulation
        maxLeft[0] = height[0];
        for (int i = 1; i < size; i++) {
            maxLeft[i] = max(maxLeft[i - 1], height[i]);
        }

        // tabulation
        maxRight[size - 1] = height[size - 1];
        for (int i = size - 2; i >= 0; i--) {
            maxRight[i] = max(maxRight[i + 1], height[i]);
        }

        int res = 0;
        // head and tail cannot trap water
        for (int i = 1; i < size - 1; i++) {
            int h = min(maxLeft[i], maxRight[i]) - height[i];
            if (h > 0) {
                res = res + h;
            }
        }
        return res;
    }
};

int main() {
    Solution solution;
    vector<int> height;

    height = {0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1};
    cout << solution.trap2(height) << endl;

    height = {4, 2, 0, 3, 2, 5};
    cout << solution.trap2(height) << endl;
}