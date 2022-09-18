#include <queue>
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
 * Depth First Search
 */
class Solution {
public:
    int maxDepth(TreeNode *root) {
        int max_depth = 0;
        dfs(max_depth, 0, root);

        return max_depth;
    }

    void dfs(int &max_depth, int cur_depth, TreeNode *&root) {
        if (!root) {
            max_depth = max(max_depth, cur_depth);
            return;
        }
        dfs(max_depth, cur_depth + 1, root->left);
        dfs(max_depth, cur_depth + 1, root->right);
    }
};

/**
 * Bredth First Search
 */
class Solution2 {
public:
    int maxDepth(TreeNode *root) {
        if (!root) {
            return 0;
        }

        int level = 0;
        queue<TreeNode *> q;

        q.push(root);
        while (!q.empty()) {
            int N = q.size();
            for (int i = 0; i < N; i++) {
                TreeNode *cur = q.front();
                q.pop();

                if (cur->left) {
                    q.push(cur->left);
                }
                if (cur->right) {
                    q.push(cur->right);
                }
            }
            level++;
        }

        return level;
    }
};
