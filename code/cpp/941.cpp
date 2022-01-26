#include <iostream>
#include <vector>
using namespace std;

class Solution {
  public:
    bool validMountainArray(vector<int> &arr) {
        int N = size(arr);
        int i = 0;

        // walk up
        for (; i < N - 1 && arr[i] < arr[i + 1];) {
            i++;
        }
        // peak can't be first or last
        if (i == N - 1 || i == 0) {
            return false;
        }

        // wark down
        for (; i < N - 1 && arr[i] > arr[i + 1];) {
            i++;
        }
        return i == N - 1;
    }
};

int main() {
    Solution solution;
    vector<int> arr;

    arr = {2, 1};
    cout << solution.validMountainArray(arr) << endl;

    arr = {3, 5, 5};
    cout << solution.validMountainArray(arr) << endl;

    arr = {0, 3, 2, 1};
    cout << solution.validMountainArray(arr) << endl;

    return 0;
}