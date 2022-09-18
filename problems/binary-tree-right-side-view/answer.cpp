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
    TreeNode(int x, TreeNode *left, TreeNode *right)
        : val(x), left(left), right(right) {}
};

class Solution {
public:
    vector<int> rightSideView(TreeNode *root) {
        vector<int> res;
        dfs(root, res, 0);

        return res;
    }

    void dfs(TreeNode *root, vector<int> &path, int depth) {
        if (root == nullptr) {
            return;
        }
        if (path.size() == depth) {
            path.push_back(root->val);
        }
        dfs(root->right, path, depth + 1);
        dfs(root->left, path, depth + 1);
    }
};

class Solution2 {
public:
    vector<int> rightSideView(TreeNode *root) {
        queue<TreeNode *> q;
        vector<int> res;
        int level = 0;

        // corner case
        if (root == nullptr) {
            return res;
        }

        // BFS
        q.push(root);
        while (!q.empty()) {
            // track level with size of level
            int levelSize = q.size();
            while (levelSize--) {
                TreeNode *node = q.front();
                q.pop();

                // visit node
                if (res.size() == level) {
                    res.push_back(node->val);
                }

                if (node->right != nullptr) {
                    q.push(node->right);
                }
                if (node->left != nullptr) {
                    q.push(node->left);
                }
            }
            level++;
        }

        return res;
    }
};

void printVector(vector<int> num) {
    for (int x : num) {
        cout << x << " ";
    }
    cout << endl;
}

int main() {
    Solution solution;
    Solution solution2;
    TreeNode *e = new TreeNode(5);
    TreeNode *d = new TreeNode(4);
    TreeNode *c = new TreeNode(3);
    TreeNode *b = new TreeNode(2, d, e);
    TreeNode *a = new TreeNode(1, b, c);

    // 1 3 5
    printVector(solution.rightSideView(a));
    printVector(solution2.rightSideView(a));
}