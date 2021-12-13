class Solution {
  public:
    int sumOfLeftLeaves(TreeNode *root) {
        int sum = 0;
        dfs(root, sum);
        return sum;
    }

    void dfs(TreeNode *root, int &sum) {
        if (!root) {
            return;
        }

        if (root->left && !root->left->left && !root->left->right) {
            sum += root->left->val;
        }

        dfs(root->left, sum);
        dfs(root->right, sum);
    }
};

/**
 * Optimized with checking if the current node is a left child.
 *
 * Refer: https://leetcode.com/problems/sum-of-left-leaves/discuss/1558055
 */
class Solution {
  public:
    int sumOfLeftLeaves(TreeNode *root) {
        int sum = 0;
        dfs(root, sum, false);
        return sum;
    }

    // Use a bool parameter `isLeft` to keep track of whether the
    // current node is a left child or not.
    void dfs(TreeNode *root, int &sum, bool isLeft) {
        if (!root) {
            return;
        }

        if (!root->left && !root->right && isLeft) {
            sum += root->val;
        }

        dfs(root->left, sum, true);
        dfs(root->right, sum, false);
    }
};

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};
