#include <iostream>
using namespace std;

/**
 * Number Theory: Digital Root
 *
 * The digital root of a natural number in a given radix is the value obtained
 * by an iterative process of summing digits, on each iteration using the result
 * from the previous iteration to compute a digit sum. The process continues
 * until a single-digit number is reached.
 *
 * Refer: https://en.wikipedia.org/wiki/Digital_root#Congruence_formula
 */
class Solution {
public:
    int addDigits(int num) {
        if (num == 0) {
            return 0;
        } else {
            return 1 + ((num - 1) % (10 - 1));
        }
    }
};

/**
 * Naive Implementation
 */
class Solution {
public:
    int addDigits(int num) {
        while (num >= 10) {
            int tmpNum = num;
            int newNum = 0;
            while (tmpNum > 0) {
                newNum += tmpNum % 10;
                tmpNum = tmpNum / 10;
            }
            num = newNum;
        }
        return num;
    }
};

int main() {
    Solution solution;
    int num;

    num = 38;
    cout << solution.addDigits(num) << endl;

    num = 0;
    cout << solution.addDigits(num) << endl;

    return 0;
}