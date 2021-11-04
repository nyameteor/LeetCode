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

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};
