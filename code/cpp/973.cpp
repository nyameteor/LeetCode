#include <cmath>
#include <iostream>
#include <queue>
#include <unordered_map>
#include <vector>
using namespace std;

/**
 * Min-Heap
 */
class Solution {
public:
    vector<vector<int>> kClosest(vector<vector<int>> &points, int k) {
        // {index in points, dinstance to origin}
        auto cmp = [](pair<int, int> left, pair<int, int> right) {
            return left.second > right.second;
        };
        priority_queue<pair<int, int>, vector<pair<int, int>>, decltype(cmp)> q(
            cmp);

        for (int i = 0; i < points.size(); i++) {
            q.push({i, distance_to_origin(points[i])});
        }

        vector<vector<int>> res;
        for (int i = 0; i < k; i++) {
            auto cur = q.top();
            q.pop();
            res.push_back(points[cur.first]);
        }
        return res;
    }

    double distance_to_origin(vector<int> &point) {
        return point[0] * point[0] + point[1] * point[1];
    }
};

void printResult(vector<vector<int>> vv) {
    for (int i = 0; i < vv.size(); i++) {
        for (int j = 0; j < vv[i].size(); j++) {
            cout << vv[i][j] << ' ';
        }
        cout << endl;
    }
    cout << endl;
}

int main() {
    Solution solution;
    vector<vector<int>> points;
    int k;

    points = {{1, 3}, {-2, 2}};
    k = 1;
    printResult(solution.kClosest(points, k));

    points = {{3, 3}, {5, -1}, {-2, 4}};
    k = 2;
    printResult(solution.kClosest(points, k));

    return 0;
}