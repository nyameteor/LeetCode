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
};

/**
 * Recursive, Preorder traverse
 *
 * Time: O(H), H is the height of BST.
 * Space: O(H), for recursive stack
 */
class Solution {
  public:
    TreeNode *deleteNode(TreeNode *root, int key) {
        if (!root) {
            return nullptr;
        }

        // searching for target
        if (root->val == key) {
            root = deleteAndMerge(root);
        } else if (root->val > key) {
            root->left = deleteNode(root->left, key);
        } else if (root->val < key) {
            root->right = deleteNode(root->right, key);
        }

        return root;
    }

    TreeNode *deleteAndMerge(TreeNode *node) {
        auto toDelete = node;
        // node has less than two children
        if (!node->left || !node->right) {
            node = node->left ? node->left : node->right;
        }
        // node has both two children
        else {
            auto prev = node->right, cur = node->right;
            while (cur->left) {
                prev = cur;      // keep the previous node of current
                cur = cur->left; // finding smallest in right subtree
            }
            node->val = cur->val; // replace node by smallest
            if (prev->left == cur) {
                prev->left = cur->right;
            } else {
                prev->right = cur->right;
            }
            delete cur;
        }

        delete toDelete;

        return node;
    }
};

TreeNode *createTree(vector<int> v) {
    int n = v.size();

    if (n <= 0)
        return nullptr;

    TreeNode **tree = new TreeNode *[n];

    for (int i = 0; i < n; i++) {
        // use -1 to represent nullptr
        if (v[i] == -1) {
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
    vector<int> nums;

    nums = {5, 3, 6, 2, 4, -1, 7};
    solution.deleteNode(createTree(nums), 3);
}