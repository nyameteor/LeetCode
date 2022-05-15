#include <iostream>
#include <queue>
#include <vector>
using namespace std;

/*
 * BFS Traversal.
 *
 * - Traverse paths with BFS.
 * - Using a new array to mark whether an index has been visited:
 *  - if already visited, stop further exploration of this path.
 */
class Solution {
public:
    bool canReach(vector<int> &arr, int start) {
        vector<bool> visited(arr.size(), false);
        queue<int> q;

        q.push(start);
        while (!q.empty()) {
            int i = q.front();
            q.pop();
            if (arr[i] == 0) { // target reached
                return true;
            } else {
                int i_a = i + arr[i];
                int i_b = i - arr[i];
                if (i_a < arr.size() && !visited[i_a]) { // try forward jump
                    q.push(i_a);
                }
                if (i_b >= 0 && !visited[i_b]) { // try backward jump
                    q.push(i_b);
                }
            }
            visited[i] = true; // mark current index as visited
        }

        return false;
    }
};

/*
 * Solution improvement.
 *
 * mark current index as visited with a more clever way:
 *  - mark current index as visited by negating
 *  - compare arr[i] with 0 to determine whether index `i` is visited
 *
 * Refer: https://leetcode.com/discuss/topic/1619031
 */

int main() {
    Solution solution;
    vector<int> arr;

    arr = {4, 2, 3, 0, 3, 1, 2};
    cout << solution.canReach(arr, 5) << endl;

    arr = {4, 2, 3, 0, 3, 1, 2};
    cout << solution.canReach(arr, 0) << endl;

    arr = {3, 0, 2, 1, 2};
    cout << solution.canReach(arr, 0) << endl;

    return 0;
}