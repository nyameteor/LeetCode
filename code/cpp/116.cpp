#include <iostream>
#include <queue>
using namespace std;

class Node {
public:
    int val;
    Node *left;
    Node *right;
    Node *next;

    Node() : val(0), left(NULL), right(NULL), next(NULL) {}

    Node(int _val) : val(_val), left(NULL), right(NULL), next(NULL) {}

    Node(int _val, Node *_left, Node *_right, Node *_next)
        : val(_val), left(_left), right(_right), next(_next) {}
};

/**
 * BFS - Left to Right
 */
class Solution {
public:
    Node *connect(Node *root) {
        // edge case
        if (!root) {
            return NULL;
        }

        queue<Node *> q;
        q.push(root);
        while (!empty(q)) {
            int sz = size(q);
            Node *cur, *next;
            for (int i = 0; i < sz; i++) {
                cur = q.front();
                q.pop();
                // if current node is the last one of its level,
                // then the next node will be NULL;
                // else the next node will be the next front of queue.
                next = (i < sz - 1) ? q.front() : NULL;
                // connect the nodes
                cur->next = next;
                // The given tree is a pefect binary tree,
                // so if current node is not leaf,
                // then it has both left and right child.
                if (cur->left) {
                    q.push(cur->left);
                    q.push(cur->right);
                }
            }
        }

        return root;
    }
};

int main() {

    Solution solution;

    return 0;
}