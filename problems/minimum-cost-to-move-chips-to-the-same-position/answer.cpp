#include <iostream>
#include <vector>
using namespace std;

class Solution {
public:
    int minCostToMoveChips(vector<int> &position) {
        int evenCnt = 0, oddCnt = 0;
        for (auto p : position) {
            if (p % 2 == 0) {
                evenCnt += 1;
            } else if (p % 2 == 1) {
                oddCnt += 1;
            }
        }
        return min(evenCnt, oddCnt) * 1;
    }
};

/**
 * check parity with `bitand`
 */
class Solution2 {
public:
    int minCostToMoveChips(vector<int> &position) {
        int evenCnt = 0, oddCnt = 0;
        for (auto p : position) {
            p & 1 ? oddCnt++ : evenCnt++;
        }
        return min(evenCnt, oddCnt) * 1;
    }
};
