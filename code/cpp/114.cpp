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

/* Morris Traversal */
class Solution {
  public:
    void flatten(TreeNode *root) {
        if (root == nullptr) {
            return;
        }

        // move left sub-tree to right
        TreeNode *old_right = root->right;
        root->right = root->left;
        root->left = nullptr;

        // insert right sub-tree to rightmost node
        TreeNode *rightmost = root;
        while (rightmost->right != nullptr) {
            rightmost = rightmost->right;
        }
        rightmost->right = old_right;

        // recursion
        flatten(root->right);
    }
};

void printResultTree(TreeNode *root) {
    while (root != nullptr) {
        cout << root->val << ' ';
        root = root->right;
    }
    cout << endl;
}

int main() {
    Solution solution;

    TreeNode *f = new TreeNode(6);
    TreeNode *e = new TreeNode(5, nullptr, f);
    TreeNode *d = new TreeNode(4);
    TreeNode *c = new TreeNode(3);
    TreeNode *b = new TreeNode(2, c, d);
    TreeNode *a = new TreeNode(1, b, e);

    solution.flatten(a);
    printResultTree(a);

    return 0;
}