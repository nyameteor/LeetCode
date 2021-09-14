#include <iostream>
#include <vector>

using namespace std;

class Solution {
  public:
    // Ordinary way, find the path and comapre the path.
    TreeNode *lowestCommonAncestor(TreeNode *root, TreeNode *p, TreeNode *q) {
        vector<TreeNode *> path_p, path_q;
        findPath(root, p, path_p);
        findPath(root, q, path_q);

        TreeNode *res;
        for (int i = 0; i < path_p.size() && i < path_q.size(); i++) {
            if (path_p[i]->val == path_q[i]->val) {
                res = path_p[i];
            }
        }
        return res;
    }

    bool findPath(TreeNode *root, TreeNode *target, vector<TreeNode *> &path) {
        if (root == target) {
            path.push_back(target);
            return true;
        }
        if (root == NULL) {
            return false;
        }

        path.push_back(root);
        if (findPath(root->left, target, path))
            return true;
        if (findPath(root->right, target, path))
            return true;
        path.pop_back();

        return false;
    }
};

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(NULL), right(NULL) {}
};