#include <iostream>
#include <vector>
using namespace std;

/**
 * Todo: Optimize
 */
class Solution {
  public:
    int numPairsDivisibleBy60(vector<int> &time) {
        vector<int> R(60, 0); // R: reminders
        for (auto t : time) {
            R[t % 60]++;
        }
        int res = 0;
        for (int i = 1; i <= 29; i++) {
            if (R[i] != 0) {
                res += R[i] * R[60 - i];
            } else {
                continue;
            }
        }
        res += combinations(R[0], 2);
        res += combinations(R[30], 2);
        return res;
    }

    int combinations(unsigned n, unsigned k) {
        if (k > n)
            return 0;
        if (k * 2 > n)
            k = n - k;
        if (k == 0)
            return 1;

        int res = n;
        for (int i = 2; i <= k; ++i) {
            res *= (n - i + 1);
            res /= i;
        }
        return res;
    }
};

int main() {
    Solution solution;
    vector<int> time;

    time = {30, 20, 150, 100, 40};
    cout << solution.numPairsDivisibleBy60(time) << endl;

    time = {60, 60, 60};
    cout << solution.numPairsDivisibleBy60(time) << endl;

    time = {20, 30, 90, 30, 100};
    cout << solution.numPairsDivisibleBy60(time) << endl;

    return 0;
}