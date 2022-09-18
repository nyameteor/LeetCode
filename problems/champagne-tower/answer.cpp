#include <iostream>
#include <vector>
using namespace std;

/**
 * Dynamic Programming
 */
class Solution {
public:
    double champagneTower(int poured, int query_row, int query_glass) {
        vector<double> currCups(1, 0);
        currCups[0] = poured;
        for (int i = 0; i < query_row; i++) {
            vector<double> nextCups((i + 1) + 1, 0);
            for (int j = 0; j <= i; j++) {
                if (currCups[j] > 1) {
                    nextCups[j] += (currCups[j] - 1) / 2.0;
                    nextCups[j + 1] += (currCups[j] - 1) / 2.0;
                }
            }
            currCups = nextCups;
        }

        return currCups[query_glass] > 1 ? 1 : currCups[query_glass];
    }
};

/**
 * Another implementation, one row with in-place
 */
class Solution2 {
public:
    double champagneTower(int poured, int query_row, int query_glass) {
        double cups[query_row + 1];
        fill(cups, cups + query_row + 1, 0);
        cups[0] = poured;
        for (int i = 0; i < query_row; i++) {
            for (int j = i; j >= 0; j--) {
                double overflow = cups[j] > 1 ? cups[j] - 1 : 0;
                cups[j + 1] += overflow / 2;
                cups[j] = overflow / 2;
            }
        }

        return cups[query_glass] > 1 ? 1 : cups[query_glass];
    }
};

int main() {
    Solution solution;
    int poured, query_row, query_glass;

    // 0.625
    poured = 10, query_row = 4, query_glass = 2;
    cout << solution.champagneTower(poured, query_row, query_glass) << endl;

    return 0;
}