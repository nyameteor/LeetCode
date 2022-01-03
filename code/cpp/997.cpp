#include <iostream>
#include <vector>
using namespace std;

/**
 * Count degree of vertex
 */
class Solution {
  public:
    int findJudge(int n, vector<vector<int>> &trust) {
        vector<int> indegree(n + 1, 0);
        vector<int> outdegree(n + 1, 0);
        for (auto t : trust) {
            outdegree[t[0]]++;
            indegree[t[1]]++;
        }
        for (int i = 1; i <= n; i++) {
            if (indegree[i] == n - 1 && outdegree[i] == 0) {
                return i;
            }
        }
        return -1;
    }
};

/**
 * Count degree of vertex
 *
 * Combine the indegree and outdegree,
 * `indegree - outdegree = N - 1` become the judge.
 */
class Solution2 {
  public:
    int findJudge(int n, vector<vector<int>> &trust) {
        vector<int> count(n + 1, 0);
        for (auto t : trust) {
            count[t[0]]--;
            count[t[1]]++;
        }
        for (int i = 1; i <= n; i++) {
            if (count[i] == n - 1) {
                return i;
            }
        }
        return -1;
    }
};

int main() {
    Solution solution;
    int n;
    vector<vector<int>> trust;

    n = 2;
    trust = {{1, 2}};
    cout << solution.findJudge(n, trust) << endl;

    n = 3;
    trust = {{1, 3}, {2, 3}};
    cout << solution.findJudge(n, trust) << endl;

    n = 3;
    trust = {{1, 3}, {2, 3}, {3, 1}};
    cout << solution.findJudge(n, trust) << endl;

    return 0;
}