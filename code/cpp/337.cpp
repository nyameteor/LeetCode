#include <iostream>
#include <unordered_map>
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
 * Recursion
 *
 * Time: O(n^2)
 */
class Solution {
  public:
    int rob(TreeNode *root) { return dfs(root); }

    int dfs(TreeNode *root) {
        if (!root) {
            return 0;
        } else if (!root->left && !root->right) {
            return root->val;
        } else {
            return max(
                root->val +
                    (root->left
                         ? (dfs(root->left->left) + dfs(root->left->right))
                         : 0) +
                    (root->right
                         ? (dfs(root->right->left) + dfs(root->right->right))
                         : 0),
                dfs(root->left) + dfs(root->right));
        }
    }
};

/**
 * Recursion + Memoization
 *
 * Time: O(n)
 */
class Solution2 {
  public:
    int rob(TreeNode *root) {
        unordered_map<TreeNode *, int> memo;
        return dfs(root, memo);
    }

    int dfs(TreeNode *root, unordered_map<TreeNode *, int> &memo) {
        if (memo[root]) {
            return memo[root];
        }

        if (!root) {
            memo[root] = 0;
        } else if (!root->left && !root->right) {
            memo[root] = root->val;
        } else {
            memo[root] = max(root->val +
                                 (root->left ? (dfs(root->left->left, memo) +
                                                dfs(root->left->right, memo))
                                             : 0) +
                                 (root->right ? (dfs(root->right->left, memo) +
                                                 dfs(root->right->right, memo))
                                              : 0),
                             dfs(root->left, memo) + dfs(root->right, memo));
        }

        return memo[root];
    }
};
