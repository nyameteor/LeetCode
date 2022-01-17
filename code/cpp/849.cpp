#include <iostream>
#include <vector>
using namespace std;

class Solution {
  public:
    int maxDistToClosest(vector<int> &seats) {
        int res = 0;
        int sz = size(seats);
        // i = -1 represent no occupied seat have been scanned
        int i = -1, j = 0;
        for (; j < sz; j++) {
            while (j < sz && seats[j] != 1) {
                j++;
            }
            int dist = j - i;
            if (i == -1) {
                // distance between start to first person
                res = max(res, (dist - 1));
            } else if (j < sz) {
                // distance between two person
                res = max(res, dist / 2);
            } else if (j == sz) {
                // distance between last person to end
                res = max(res, (dist - 1));
            }
            i = j;
        }

        return res;
    }
};

int main() {
    Solution solution;
    vector<int> seats;

    seats = {1, 0, 0, 0, 1, 0, 1};
    cout << solution.maxDistToClosest(seats) << endl;

    seats = {1, 0, 0, 0};
    cout << solution.maxDistToClosest(seats) << endl;

    seats = {0, 1};
    cout << solution.maxDistToClosest(seats) << endl;

    seats = {1, 0, 0, 1};
    cout << solution.maxDistToClosest(seats) << endl;

    return 0;
}