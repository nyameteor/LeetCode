#include <iostream>
#include <vector>
using namespace std;

class Solution {
  public:
    bool isPalindrome(int x) {}
};

/**
 * Idea: Convert to array
 *
 * Time: O(n), Space: O(n)
 */
class Solution2 {
  public:
    bool isPalindrome(int x) {
        if (x < 0) {
            return false;
        }

        vector<int> v = intToArray(x);
        int i = 0, j = v.size() - 1;
        while (i <= j) {
            if (v[i] != v[j]) {
                return false;
            }
            i++;
            j--;
        }
        return true;
    }

    vector<int> intToArray(int x) {
        vector<int> v;
        while (x != 0) {
            v.push_back(x % 10);
            x /= 10;
        }
        return v;
    }
};

int main() {
    Solution solution;
    cout << solution.isPalindrome(121) << endl;
    cout << solution.isPalindrome(10) << endl;

    return 0;
}