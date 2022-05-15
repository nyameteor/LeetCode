#include <algorithm>
#include <iostream>
#include <vector>
using namespace std;

/**
 * Sort & Try Adjacent Pair of Elements
 *
 * Time: O(NlogN)
 * Space: O(logN)
 * (When using quick sort as sorting algorithm)
 */
class Solution {
public:
    vector<vector<int>> minimumAbsDifference(vector<int> &arr) {
        vector<vector<int>> res;
        sort(arr.begin(), arr.end());

        int minDiff = INT32_MAX;
        for (int i = 0; i < arr.size() - 1; i++) {
            int currDiff = arr[i + 1] - arr[i];
            if (currDiff < minDiff) {
                minDiff = currDiff;
                res.clear();
                res.push_back({arr[i], arr[i + 1]});
            } else if (currDiff == minDiff) {
                res.push_back({arr[i], arr[i + 1]});
            } else {
                continue;
            }
        }
        return res;
    }
};

/**
 * Counting Sort
 *
 * Refer:
 * https://leetcode.com/problems/minimum-absolute-difference/discuss/1637287
 */