#include <algorithm>
#include <cstdlib>
#include <ctime>
#include <iostream>
#include <vector>

using namespace std;

class Solution {
public:
    int rand10() {
        int row, col, res;

        // rejection sampling
        // use `do while` to do at least once
        do {
            row = rand7();
            col = rand7();
            res = (row - 1) * 7 + col;
        } while (res > 40);

        return 1 + (res - 1) % 10;
    }

    // Ref: https://www.cplusplus.com/reference/cstdlib/rand/
    int rand7() {

        /* generate secret number between 1 and 10: */
        return rand() % 7 + 1;
    }
};

int main() {
    Solution solution;

    /* initialize random seed: */
    srand(time(NULL));

    cout << solution.rand10() << endl;
    cout << solution.rand10() << endl;
    cout << solution.rand10() << endl;
}