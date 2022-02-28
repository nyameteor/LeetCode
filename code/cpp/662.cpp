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
    int widthOfBinaryTree(TreeNode *root) {

        if (!root) {
            return 0;
        }

        unsigned long long max_width = 0;

        queue<pair<TreeNode *, unsigned long long>> q;
        q.push({root, 1});
        while (!empty(q)) {
            int level_size = size(q);

            unsigned long long left_pos = q.front().second;
            unsigned long long cur_pos = left_pos;

            for (int i = 0, j = 0; i < level_size; i++) {
                auto cur = q.front();
                q.pop();

                auto node = cur.first;
                cur_pos = cur.second;
                if (node->left) {
                    q.push({node->left, 2 * cur_pos});
                }
                if (node->right) {
                    q.push({node->right, 2 * cur_pos + 1});
                }
            }

            // now, current pos is the rightmost in this level
            max_width = max(max_width, cur_pos - left_pos + 1);
        }

        return max_width;
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
    vector<int> nodes;

    // 2
    nodes = {
        0, 0, 0,  -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1,
        0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1,
        0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1,
        0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1,
        0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1,
        0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1,
        0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1,
        0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1,
        0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1,
        0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1,
        0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1,
        0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1,
        0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1, -1,
        0, 0, -1, -1, 0, 0, -1, -1, 0, 0, -1};
    cout << solution.widthOfBinaryTree(createTree(nodes)) << endl;

    // 2147483645
    nodes = {1,  1,  1,  1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1,
             1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,
             -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1,
             1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,
             -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1,
             1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,
             -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1, 1,  -1, -1, 1,  1,  -1,
             1,  -1, 1,  -1, 1,  -1, 1,  -1, -1, -1, 1};
    cout << solution.widthOfBinaryTree(createTree(nodes)) << endl;

    return 0;
}