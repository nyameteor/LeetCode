#include <cmath>
#include <iostream>
#include <unordered_set>
#include <vector>
using namespace std;

/**
 * Detect cycle with hash set
 *
 * Idea: We can also use hash map, but in this
 * problem, we only need to check if the value is in the collection or not, so
 * use set is better.
 *
 * Time: O(n)
 * Space: O(n)
 */
class Solution {
  public:
    bool isHappy(int n) {
        unordered_set<int> s;

        while (true) {
            if (n == 1) {
                return true;
            }

            int sum = 0;
            while (n != 0) {
                int d = n % 10;
                sum += d * d;
                n /= 10;
            }

            if (s.find(sum) != end(s)) {
                return false;
            } else {
                s.insert(sum);
            }
            n = sum;
        }
    }
};

int main() {
    Solution solution;

    cout << solution.isHappy(19) << endl;
    cout << solution.isHappy(2) << endl;

    return 0;
}