#include <cstdio>
#include <iostream>
#include <map>
#include <vector>

using namespace std;

class Solution {
public:
    int climbStairs(int n) {
        map<int, int> memo;
        return dp(n, memo);
    }

    int dp(int n, map<int, int> &memo) {
        // find in memo
        if (memo.find(n) != memo.end())
            return memo.at(n);
        if (n == 0)
            return 1;
        if (n < 0)
            return 0;

        memo.insert(make_pair(n, dp(n - 1, memo) + dp(n - 2, memo)));
        return memo.at(n);
    }
};

int main() {
    Solution solution;

    // 2
    cout << solution.climbStairs(2) << endl;

    // 3
    cout << solution.climbStairs(3) << endl;

    // 1836311903
    cout << solution.climbStairs(45) << endl;
}