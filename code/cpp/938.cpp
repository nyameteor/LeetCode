#include <iostream>
#include <queue>
#include <vector>
using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode() : val(0), left(nullptr), right(nullptr) {}
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
    TreeNode(int x, TreeNode *left, TreeNode *right)
        : val(x), left(left), right(right) {}
};

/**
 * BFS, Optimized for BST
 */
class Solution {
public:
    int rangeSumBST(TreeNode *root, int low, int high) {
        queue<TreeNode *> q;
        q.push(root);

        int sum = 0;
        while (!q.empty()) {
            auto node = q.front();
            q.pop();
            int val = node->val;
            if (val >= low && val <= high) {
                sum += val;
            }
            if (val > low && node->left) { // pruning
                q.push(node->left);
            }
            if (val < high && node->right) { // pruning
                q.push(node->right);
            }
        }
        return sum;
    }
};

/**
 * DFS, Optimized for BST
 */
class Solution2 {
public:
    int rangeSumBST(TreeNode *root, int low, int high) {
        if (!root) {
            return 0;
        }

        int val = root->val;
        return (val >= low && val <= high ? val : 0) +
               (val > low ? rangeSumBST(root->left, low, high) : 0) + // pruning
               (val < high ? rangeSumBST(root->right, low, high)
                           : 0); // pruning
    }
};

/**
 * Todo: Morris Traversal
 */
