#include <cmath>
#include <queue>
#include <vector>
using namespace std;

class Solution {
public:
    int minCostConnectPoints(vector<vector<int>> &points) {
        // min-heap to store minimum cost edge at top
        // pair: {node, cost}
        auto cmp = [](const auto &l, const auto &r) {
            return l.second > r.second;
        };
        priority_queue<pair<int, int>, vector<pair<int, int>>, decltype(cmp)>
            heap(cmp);

        int N = size(points);

        // mark if a node is known
        vector<bool> known(N);

        // add first selected node
        heap.push({0, 0});

        // sum of cost of the minimum spanning tree
        int allCost = 0;

        // count the used edges
        int edgeUsed = 0;

        // loop while used edges < N
        while (edgeUsed < N) {
            // get the node with minimum cost
            auto pair = heap.top();
            heap.pop();

            int curNode = pair.first;
            int curCost = pair.second;

            // if curNode is known, then it can't be selected
            if (known[curNode]) {
                continue;
            }
            // mark curNode as known
            known[curNode] = true;

            edgeUsed++;

            allCost += curCost;

            // add candidate node(nextNode),
            // where curNode is known and nextNode is unknown
            for (int nextNode = 0; nextNode < N; nextNode++) {
                if (!known[nextNode]) {
                    int nextCost =
                        abs(points[curNode][0] - points[nextNode][0]) +
                        abs(points[curNode][1] - points[nextNode][1]);
                    heap.push({nextNode, nextCost});
                }
            }
        }

        return allCost;
    }
};