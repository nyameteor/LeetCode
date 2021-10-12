#include <algorithm>
#include <iostream>
#include <vector>
using namespace std;

/**
 * Sorting, easy to understand
 */
class Solution {
  public:
    vector<vector<int>> merge(vector<vector<int>> &intervals) {
        vector<vector<int>> res;
        sort(intervals.begin(), intervals.end(),
             [](const auto &left, const auto &right) {
                 return left[0] < right[0];
             });

        for (auto current : intervals) {
            // corner case
            if (res.size() == 0) {
                res.push_back(current);
                continue;
            }

            vector<int> previous = res.back();
            if (previous[1] >= current[0]) {
                // merge intervals
                res.pop_back();
                if (previous[1] >= current[1]) {
                    res.push_back({previous[0], previous[1]});
                } else {
                    res.push_back({previous[0], current[1]});
                }
            } else {
                // if not merged, simply append it
                res.push_back(current);
            }
        }

        return res;
    }
};

/**
 * Sorting, with simplified steps
 */
class Solution2 {
  public:
    vector<vector<int>> merge(vector<vector<int>> &intervals) {
        sort(intervals.begin(), intervals.end());

        vector<vector<int>> merged;
        for (auto interval : intervals) {
            // if the list of merged intervals is empty or if the current
            // interval does not overlap with the previous, simply append it.
            if (merged.empty() || merged.back()[1] < interval[0]) {
                merged.push_back(interval);
            }
            // otherwise, there is overlap, so we merge the current and previous
            // intervals.
            else {
                merged.back()[1] = max(merged.back()[1], interval[1]);
            }
        }
        return merged;
    }
};

void printResult(vector<vector<int>> res) {
    for (auto x : res) {
        for (auto y : x) {
            cout << y << ' ';
        }
    }
    cout << endl;
}

int main() {
    Solution solution;
    Solution2 solution2;
    vector<vector<int>> intervals;

    intervals = {{1, 3}, {2, 6}, {8, 10}, {15, 18}};
    printResult(solution.merge(intervals));
    printResult(solution2.merge(intervals));

    intervals = {{1, 4}, {4, 5}};
    printResult(solution.merge(intervals));
    printResult(solution2.merge(intervals));

    intervals = {{1, 9}, {2, 6}, {4, 10}};
    printResult(solution.merge(intervals));
    printResult(solution2.merge(intervals));
}