#include <cmath>
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
/**
 * Backtracking
 *
 * Idea:
 *  1. count the tree maximum level.
 *  2. count the number of nodes in the last level, if one node is leaf and
 *     current level is maximum, then the node is in the last level.
 *  3. number of nodes = 2^(maxLevel - 1) - 1 + lastLevelCnt
 *
 * Simple and intuitive
 */
class Solution {
  public:
    int countNodes(TreeNode *root) {
        int level = 0, maxLevel = 0, lastLevelCnt = 0;
        dfs(root, level, maxLevel, lastLevelCnt);

        return pow(2, (maxLevel - 1)) - 1 + lastLevelCnt;
    }

    void dfs(TreeNode *root, int &level, int &maxLevel, int &lastLevelCnt) {
        if (!root) {
            return;
        }

        level++;
        maxLevel = max(maxLevel, level);

        // if node is leaf
        if (!root->left && !root->right) {
            // if node is in last level
            if (level == maxLevel) {
                lastLevelCnt++;
            } else {
                // tree pruning
                level--; // backtrack (don't forget it)
                return;
            }
        }

        dfs(root->left, level, maxLevel, lastLevelCnt);
        dfs(root->right, level, maxLevel, lastLevelCnt);

        level--; // backtrack
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
    Solution solution;
    vector<int> tree;

    tree = {1, 2, 3, 4, 5, 6};
    cout << solution.countNodes(createTree(tree)) << endl;

    tree = {1, 2, 3, 4, 5, 6, 7, 8, 9, 10};
    cout << solution.countNodes(createTree(tree)) << endl;

    tree = {1};
    cout << solution.countNodes(createTree(tree)) << endl;

    return 0;
}