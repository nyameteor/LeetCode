#include <iostream>
#include <vector>
using namespace std;

class Solution {
  private:
    vector<int> order;
    vector<int> indegree;

  public:
    vector<int> findOrder(int numCourses, vector<vector<int>> &prerequisites) {
        vector<vector<int>> G(numCourses);
        indegree = vector<int>(numCourses, 0);
        for (auto &req : prerequisites) { // convert to graph
            G[req[1]].push_back(req[0]);
            indegree[req[0]]++;
        }
        order = vector<int>(0);

        for (int v = 0; v < numCourses; v++) {
            if (indegree[v] == 0) {
                dfs(G, v);
            }
        }

        if (size(order) == numCourses) { // avoid cases that graph has cycle
            return order;
        } else {
            return {};
        }
    }

    void dfs(vector<vector<int>> G, int v) {
        order.push_back(v); // push current vertex in ordering
        indegree[v] = -1;   // mark current vertex as visited
        for (int w : G[v]) {
            indegree[w]--;
            if (indegree[w] == 0) { // if there's a next vertex having 0 indegree
                dfs(G, w);          // then we traverse it.
            }
        }
    }
};

void printResult(vector<int> v) {
    for (auto x : v) {
        cout << x << ' ';
    }
    cout << endl;
}

int main() {
    Solution solution;
    int numCourses;
    vector<vector<int>> prerequisites;

    numCourses = 2;
    prerequisites = {{1, 0}};
    printResult(solution.findOrder(numCourses, prerequisites));

    numCourses = 4;
    prerequisites = {{1, 0}, {2, 0}, {3, 1}, {3, 2}};
    printResult(solution.findOrder(numCourses, prerequisites));

    return 0;
}