#include <iostream>
#include <vector>
using namespace std;

class Solution {
  public:
    int numTrees(int n) {
        vector<int> a((n + 1), 0);
        a[0] = 1;
        for (int i = 1; i <= n; i++) {
            for (int j = 1; j <= i; j++) {
                a[i] += a[j - 1] * a[i - j];
            }
        }
        return a[n];
    }
};

int main() {
    Solution solution;
    cout << solution.numTrees(1) << endl;
    cout << solution.numTrees(3) << endl;
    cout << solution.numTrees(4) << endl;
    cout << solution.numTrees(8) << endl;

    return 0;
}