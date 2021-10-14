#include <algorithm>
#include <iostream>
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

class Solution {
  public:
    bool isValidBST(TreeNode *root) {
        vector<int> path;
        inOrder(root, path);

        // check if path are sorted
        for (int i = 0; i < path.size() - 1; i++) {
            if (path[i] >= path[i + 1])
                return false;
        }
        return true;
    }

    void inOrder(TreeNode *root, vector<int> &path) {
        if (root == nullptr) {
            return;
        }
        inOrder(root->left, path);
        path.push_back(root->val);
        inOrder(root->right, path);
    }
};

TreeNode *createTree(vector<int> v) {
    int n = v.size();

    if (n <= 0)
        return nullptr;

    TreeNode **tree = new TreeNode *[n];

    for (int i = 0; i < n; i++) {
        if (v[i] == 0) {
            tree[i] = nullptr;
            continue;
        }
        tree[i] = new TreeNode(v[i]);
    }
    int pos = 1;
    for (int i = 0; i < n && pos < n; i++) {
        if (tree[i]) {
            tree[i]->left = tree[pos++];
            if (pos < n) {
                tree[i]->right = tree[pos++];
            }
        }
    }
    return tree[0];
}

int main() {
    Solution solution;

    cout << solution.isValidBST(createTree({2, 1, 3})) << endl;
    cout << solution.isValidBST(createTree({5, 1, 4, 0, 0, 3, 6})) << endl;
    cout << solution.isValidBST(createTree({2, 2, 2})) << endl;

    return 0;
}