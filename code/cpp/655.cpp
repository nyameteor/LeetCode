#include <cmath>
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

class Solution {
  public:
    vector<vector<string>> printTree(TreeNode *root) {
        int height = getHeight(root) - 1;
        int m = height + 1;
        int n = pow(2, height + 1) - 1;
        vector<vector<string>> res(m, vector<string>(n));

        queue<pair<TreeNode *, int>> q;
        int row = 0;
        int column = (n - 1) / 2;
        q.push({root, column});
        while (!q.empty()) {
            int qz = q.size();

            for (int i = 0; i < qz; i++) {
                auto c = q.front();
                q.pop();
                auto node = c.first;
                int column = c.second;
                res[row][column] = to_string(node->val); // int to string

                if (node->left) {
                    q.push({node->left, column - pow(2, height - row - 1)});
                }
                if (node->right) {
                    q.push({node->right, column + pow(2, height - row - 1)});
                }
            }
            row++;
        }

        return res;
    }

    int getHeight(TreeNode *root) {
        if (!root) {
            return 0;
        } else {
            return 1 + max(getHeight(root->left), getHeight(root->right));
        }
    }
};
