#include <algorithm>
#include <cassert>
#include <vector>
using namespace std;

class Solution {
public:
    int threeSumMulti(vector<int> &arr, int target) {
        int MOD = 1000000007;
        int N = size(arr);
        long res = 0;

        sort(arr.begin(), arr.end());

        for (int i = 0; i < N; i++) {

            int T = target - arr[i];
            int j = i + 1;
            int k = N - 1;

            while (j < k) {
                if (arr[j] + arr[k] < T) {
                    j++;
                } else if (arr[j] + arr[k] > T) {
                    k--;
                } else if (arr[j] != arr[k]) {
                    int left = 1;
                    int right = 1;

                    while (j + 1 < k && arr[j] == arr[j + 1]) {
                        left++;
                        j++;
                    }
                    while (k - 1 > j && arr[k] == arr[k - 1]) {
                        right++;
                        k--;
                    }

                    res += left * right;
                    res = res % MOD;

                    j++;
                    k--;
                } else {
                    res += (k - j + 1) * (k - j) / 2;
                    res = res % MOD;

                    break;
                }
            }
        }

        return (int)res;
    }
};

int main() {
    Solution solution;
    vector<int> arr;
    int target;
    int res;

    arr = {1, 1, 2, 2, 3, 3, 4, 4, 5, 5};
    target = 8;
    res = 20;
    assert(solution.threeSumMulti(arr, target) == res);

    return 0;
}