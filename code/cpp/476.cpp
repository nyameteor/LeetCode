#include <cmath>
#include <iostream>
using namespace std;

/**
 * Flip Bits from Right to Left - Extra Space
 */
class Solution {
  public:
    int findComplement(int num) {
        int res = 0;
        int i = 0; // position of bit
        while (num) {
            // flip the rightmost bit, calculate and add its value
            res += (num ^ 1) * pow(2, i);
            num = num >> 1; // right shift
            i++;
        }
        return res;
    }
};

/**
 * Flip Bits from Right to Left - In-place
 */
class Solution2 {
  public:
    int findComplement(int num) {
        int i = 1;
        while (i <= num) {
            num ^= i;
            i <<= 1;
        }
        return num;
    }
};

int main() {
    Solution solution;

    cout << solution.findComplement(5) << endl;
    cout << solution.findComplement(2) << endl;
    cout << solution.findComplement(1) << endl;

    cout << (8 ^ 1) << endl;

    return 0;
}