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
    bool isCompleteTree(TreeNode *root) {
        bool isEnd = false;
        queue<TreeNode *> q;

        q.push(root);
        while (!q.empty()) {
            auto p = q.front();
            q.pop();
            if (!p) {
                isEnd = true;
                continue;
            } else if (isEnd) {
                return false;
            }
            q.push(p->left);
            q.push(p->right);
        }
        return true;
    }
};

int main() {
    Solution solution;
    TreeNode *f = new TreeNode(7);
    TreeNode *e = new TreeNode(5);
    TreeNode *d = new TreeNode(4);
    TreeNode *c = new TreeNode(3, nullptr, f);
    TreeNode *b = new TreeNode(2, d, e);
    TreeNode *a = new TreeNode(1, b, c);

    cout << solution.isCompleteTree(a) << endl;
}