#include <iostream>
#include <vector>
using namespace std;

struct TreeNode {
    int val;
    TreeNode *left;
    TreeNode *right;
    TreeNode(int x) : val(x), left(nullptr), right(nullptr) {}
};

/**
 * Time: O(n), Space: O(1)
 */
class Solution {
  public:
    int sumNumbers(TreeNode *root) { return dfs(root, 0); }

    int dfs(TreeNode *root, int cur) {
        if (root == nullptr) {
            return 0;
        }

        cur = cur * 10 + root->val;

        if (root->left == nullptr && root->right == nullptr) {
            return cur;
        }

        return dfs(root->left, cur) + dfs(root->right, cur);
    }
};

/**
 * Idea: same as first solution, but use more intermediate variables,
 * for me it's straightforward.
 *
 * FIXME: this solution(4ms) is much slower than first(0ms), need test.
 *
 * Time: O(n), Space: O(1)
 */
class Solution2 {
  public:
    int sumNumbers(TreeNode *root) {
        int pathSum = 0, allSum = 0;
        dfs(root, pathSum, allSum);

        return allSum;
    }

    void dfs(TreeNode *root, int pathSum, int &allSum) {
        if (root == nullptr) {
            return;
        }

        pathSum = pathSum * 10 + root->val;

        if (root->left == nullptr && root->right == nullptr) {
            allSum += pathSum;
        }

        dfs(root->left, pathSum, allSum);
        dfs(root->right, pathSum, allSum);
    }
};

/**
 * Idea: store every path when traverse the tree, and calculate it's sum
 *
 * Time: O(n), Space: O(n)
 */
class Solution3 {
  public:
    int sumNumbers(TreeNode *root) {
        vector<int> path;
        int sum = 0;

        dfs(root, path, sum);

        return sum;
    }

    void dfs(TreeNode *root, vector<int> &path, int &sum) {
        if (root == nullptr) {
            return;
        }

        if (root->left == nullptr && root->right == nullptr) {
            sum += pathSum(path) * 10 + root->val;
        }

        path.push_back(root->val);

        dfs(root->left, path, sum);
        dfs(root->right, path, sum);

        path.pop_back();
    }

    int pathSum(vector<int> path) {
        int sum = 0;
        for (int i = 0; i < path.size(); i++) {
            sum = sum * 10 + path[i];
        }
        return sum;
    }
};

TreeNode *createTree(vector<int> v) {
    int n = v.size();

    if (n <= 0)
        return nullptr;

    TreeNode **tree = new TreeNode *[n];

    for (int i = 0; i < n; i++) {
        // use -1 to present nullptr
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
    Solution2 solution;

    cout << solution.sumNumbers(createTree({1, 2, 3})) << endl;
    cout << solution.sumNumbers(createTree({4, 9, 0, 5, 1})) << endl;
    cout << solution.sumNumbers(createTree({1, -1, 5})) << endl;

    return 0;
}