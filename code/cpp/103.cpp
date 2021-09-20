#include <iostream>
#include <queue>
#include <stack>
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

class Solution {
  public:
    // Two stacks
    vector<vector<int>> zigzagLevelOrder(TreeNode *root) {
        stack<TreeNode *> s1;
        stack<TreeNode *> s2;
        vector<vector<int>> res;
        int depth = 0;

        s2.push(root);
        while (!s1.empty() || !s2.empty()) {
            // part of result
            vector<int> res_part;

            // even depth: use stack 1 to get right -> left order
            if (depth % 2 == 0) {
                while (!s2.empty()) {
                    TreeNode *node = s2.top();
                    s2.pop();
                    if (node != nullptr) {
                        res_part.push_back(node->val);
                        s1.push(node->left);
                        s1.push(node->right);
                    }
                }
            }

            // odd depth: use stack 2 to get left -> right order
            if (depth % 2 == 1) {
                while (!s1.empty()) {
                    TreeNode *node = s1.top();
                    s1.pop();
                    if (node != nullptr) {
                        res_part.push_back(node->val);
                        s2.push(node->right);
                        s2.push(node->left);
                    }
                }
            }

            if (res_part.size() > 0) {
                res.push_back(res_part);
            }
            depth += 1;
        }

        return res;
    }
};

int main() {
    Solution Solution;
    vector<vector<int>> res;
    TreeNode *d = new TreeNode(15);
    TreeNode *e = new TreeNode(7);
    TreeNode *b = new TreeNode(9);
    TreeNode *c = new TreeNode(20, d, e);
    TreeNode *a = new TreeNode(3, b, c);

    res = Solution.zigzagLevelOrder(a);

    for (int i = 0; i < res.size(); i++) {
        for (int j = 0; j < res[i].size(); j++) {
            cout << res[i][j] << " ";
        }
        cout << endl;
    }
}