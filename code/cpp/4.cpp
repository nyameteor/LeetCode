#include <iostream>
#include <vector>
using namespace std;

/**
 * Refer: https://leetcode.com/problems/median-of-two-sorted-arrays/discuss/2471
 *
 * Time: O(log(M+N))
 * Space: O(1)
 */
class Solution {
  public:
    double findMedianSortedArrays(vector<int> &nums1, vector<int> &nums2) {
        if (size(nums1) > size(nums2)) {
            // make sure nums1 is shorter
            // swap nums1 and nums2 also works
            return findMedianSortedArrays(nums2, nums1);
        }

        int N1 = size(nums1);
        int N2 = size(nums2);

        int lo = 0;
        int hi = N1;
        while (lo <= hi) {
            int C1 = (lo + hi) / 2;      // try cut form the middle
            int C2 = (N1 + N2) / 2 - C1; // C1 + C2 = (N1 + N2) / 2

            // C = 0, means there are no leaft half
            int L1 = C1 == 0 ? INT32_MIN : nums1[C1 - 1];
            int L2 = C2 == 0 ? INT32_MIN : nums2[C2 - 1];
            // C = N, means there are no right half
            int R1 = C1 == N1 ? INT32_MAX : nums1[C1];
            int R2 = C2 == N2 ? INT32_MAX : nums2[C2];

            if (L1 > R2) {
                hi = C1 - 1;
            } else if (L2 > R1) {
                lo = C1 + 1;
            } else {
                if ((N1 + N2) % 2) {
                    // N1 + N2 is odd
                    return min(R1, R2);
                } else {
                    // N1 + N2 is even
                    return (max(L1, L2) + min(R1, R2)) / 2.0;
                }
            }
        }
        return -1;
    }
};

int main() {
    Solution solution;
    vector<int> nums1, nums2;

    // 2.5
    nums1 = {1, 2};
    nums2 = {3, 4};
    cout << solution.findMedianSortedArrays(nums1, nums2) << endl;

    // 3
    nums1 = {1, 2, 3, 4, 5, 6};
    nums2 = {1, 2, 3, 4, 5};
    cout << solution.findMedianSortedArrays(nums1, nums2) << endl;

    return 0;
}