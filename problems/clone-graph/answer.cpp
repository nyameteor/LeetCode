#include <iostream>
#include <queue>
#include <unordered_map>
#include <unordered_set>
#include <vector>
using namespace std;

class Node {
public:
    int val;
    vector<Node *> neighbors;
    Node() {
        val = 0;
        neighbors = vector<Node *>();
    }
    Node(int _val) {
        val = _val;
        neighbors = vector<Node *>();
    }
    Node(int _val, vector<Node *> _neighbors) {
        val = _val;
        neighbors = _neighbors;
    }
};

/**
 * Clone graph with BFS
 */
class Solution {
public:
    Node *cloneGraph(Node *node) {
        if (!node) {
            return nullptr;
        }

        // create a map, key is source node, value is clone node
        unordered_map<Node *, Node *> cloneMap;

        // clone the root
        Node *cloneNode = new Node(node->val);
        cloneMap[node] = cloneNode;

        // breadth first search
        queue<Node *> q;
        q.push(node);
        while (!q.empty()) {
            Node *c = q.front();
            q.pop();

            // for each neighbor
            for (auto n : c->neighbors) {
                // neighbor not exsisted in cloneMap
                if (cloneMap.find(n) == end(cloneMap)) {
                    // clone the neighbor to a new node
                    Node *newNode = new Node(n->val);
                    cloneMap[n] = newNode;

                    // put the neighbor into the queue
                    q.push(n);
                }
                cloneMap[c]->neighbors.push_back(cloneMap[n]);
            }
        }

        return cloneNode;
    }
};

// Print graph with BFS
void printGraph(Node *node) {
    unordered_set<int> visited;
    queue<Node *> q;

    q.push(node);
    visited.insert(node->val);
    while (!q.empty()) {
        Node *c = q.front();
        q.pop();
        cout << c->val << endl;

        for (auto n : c->neighbors) {
            if (visited.find(n->val) != end(visited)) {
                continue;
            } else {
                q.push(n);
                visited.insert(n->val);
            }
        }
    }
}

Node *createGraph(vector<vector<int>> vv) {
    vector<Node *> nodes;
    // create nodes
    for (int i = 0; i < vv.size(); i++) {
        nodes.push_back(new Node(i + 1));
    }
    // add adjs
    for (int i = 0; i < vv.size(); i++) {
        for (int j = 0; j < vv[i].size(); j++) {
            nodes[i]->neighbors.push_back(nodes[vv[i][j] - 1]);
        }
    }
    return nodes[0];
}

int main() {
    Solution solution;
    vector<vector<int>> adjList;

    adjList = {{2, 4}, {1, 3}, {2, 4}, {1, 3}};
    printGraph(solution.cloneGraph(createGraph(adjList)));

    return 0;
}