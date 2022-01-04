#include <iostream>
#include <unordered_set>
using namespace std;

class Solution {
  public:
    int smallestRepunitDivByK(int k) {
        int n = 1;
        int i = 1;
        int R;                   // remainder
        unordered_set<int> seen; // store seen remainders
        while (true) {
            R = n % k;
            if (R == 0) {
                return i;
            }
            if (seen.count(R)) {
                return -1;
            }
            seen.insert(R); // mark remainder as seen
            n = R * 10 + 1;
            i++;
        }
        return -1;
    }
};

int main() {
    Solution solution;

    for (int i = 1; i < 50; i++) {
        cout << solution.smallestRepunitDivByK(i) << endl;
    }
}