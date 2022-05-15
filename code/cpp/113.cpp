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

/* Recursion */
class Solution {
public:
    vector<vector<int>> pathSum(TreeNode *root, int targetSum) {
        vector<vector<int>> paths;
        vector<int> path;

        dfs(root, targetSum, path, paths);
        return paths;
    }

    void dfs(TreeNode *root, int targetSum, vector<int> path,
             vector<vector<int>> &paths) {
        if (root == nullptr) {
            return;
        }

        path.push_back(root->val);
        targetSum = targetSum - root->val;

        if (targetSum == 0 && root->left == nullptr && root->right == nullptr) {
            paths.push_back(path);
        }

        dfs(root->left, targetSum, path, paths);
        dfs(root->right, targetSum, path, paths);
    }
};

/* Recursive Backtracking */
class Solution2 {
public:
    vector<vector<int>> pathSum(TreeNode *root, int targetSum) {
        vector<vector<int>> paths;
        vector<int> path;

        dfs(root, targetSum, path, paths);
        return paths;
    }

    void dfs(TreeNode *root, int targetSum, vector<int> &path,
             vector<vector<int>> &paths) {
        if (root == nullptr) {
            return;
        }

        path.push_back(root->val);
        targetSum = targetSum - root->val;

        if (targetSum == 0 && root->left == nullptr && root->right == nullptr) {
            paths.push_back(path);
        }

        dfs(root->left, targetSum, path, paths);
        dfs(root->right, targetSum, path, paths);

        // backtrack, will reuse path
        path.pop_back();
    }
};

void printResult(vector<vector<int>> vv) {
    for (int i = 0; i < vv.size(); i++) {
        for (int j = 0; j < vv[i].size(); j++) {
            cout << vv[i][j] << ' ';
        }
        cout << endl;
    }
}

int main() {
    Solution solution;
    Solution2 solution2;

    TreeNode *e = new TreeNode(2);
    TreeNode *d = new TreeNode(7);
    TreeNode *c = new TreeNode(11, d, e);
    TreeNode *b = new TreeNode(4, c, nullptr);
    TreeNode *a = new TreeNode(5, b, nullptr);

    printResult(solution.pathSum(a, 22));
    printResult(solution2.pathSum(a, 22));

    return 0;
}