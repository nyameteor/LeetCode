#include <iostream>
#include <stack>
#include <vector>
using namespace std;

/**
 * Monotonic Stack
 *
 * Time: O(n), Space: O(n)
 */
class Solution {
  public:
    vector<int> dailyTemperatures(vector<int> &temperatures) {
        int size = temperatures.size();
        vector<int> res(size);
        stack<int> st;

        for (int i = size - 1; i >= 0; i--) {
            // keep stack monotonic decreasing
            while (!st.empty() && !(temperatures[i] < temperatures[st.top()])) {
                st.pop();
            }
            // get distance between current index and next `greater` index
            res[i] = !st.empty() ? st.top() - i : 0;
            // put current index to stack
            st.push(i);
        }

        return res;
    }
};

/**
 * Brute Force
 *
 * Time: O(n^2), Space: O(1)
 */
class Solution2 {
  public:
    vector<int> dailyTemperatures(vector<int> &temperatures) {
        vector<int> res;

        int i, j;
        int size = temperatures.size();
        for (i = 0; i < size; i++) {
            for (j = i; j < size; j++) {
                if (temperatures[j] > temperatures[i]) {
                    break;
                }
            }
            if (j < size) {
                res.push_back(j - i);
            } else {
                res.push_back(0);
            }
        }

        return res;
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

    nums = {73, 74, 75, 71, 69, 72, 76, 73};
    printResult(solution.dailyTemperatures(nums));

    nums = {30, 40, 50, 60};
    printResult(solution.dailyTemperatures(nums));

    nums = {71, 71, 71, 71, 71};
    printResult(solution.dailyTemperatures(nums));

    nums = {89, 62, 70, 58, 47, 47, 46, 76, 100, 70};
    printResult(solution.dailyTemperatures(nums));

    return 0;
}