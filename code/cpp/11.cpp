#include <vector>
using namespace std;

/**
 * Two Pointers + Greedy
 *
 * In leetcode it's slower than without greedy, don't know why.
 */
class Solution {
public:
    int maxArea(vector<int> &height) {
        int l = 0, r = size(height) - 1;
        int max_area = 0;
        while (l < r) {
            int lh = height[l];
            int rh = height[r];

            max_area = max(max_area, (r - l) * min(lh, rh));

            if (lh < rh) {
                while (l < r && height[l] <= lh) {
                    l++;
                }
            } else {
                while (l < r && height[r] <= rh) {
                    r--;
                }
            }
        }

        return max_area;
    }
};

/**
 * Two Pointers
 */
class Solution2 {
public:
    int maxArea(vector<int> &height) {
        int l = 0, r = size(height) - 1;
        int max_area = 0;
        while (l < r) {
            int lh = height[l];
            int rh = height[r];

            max_area = max(max_area, (r - l) * min(lh, rh));

            if (lh < rh) {
                l++;
            } else {
                r--;
            }
        }

        return max_area;
    }
};