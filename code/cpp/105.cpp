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
    TreeNode *buildTree(vector<int> &preorder, vector<int> &inorder) {
        if (size(preorder) == 0) {
            return nullptr;
        }

        TreeNode *node = new TreeNode(preorder[0]);

        if (size(preorder) == 1) {
            return node;
        }

        // find the pivot in inorder
        int i;
        for (i = 0; i < size(inorder); i++) {
            if (preorder[0] == inorder[i]) {
                break;
            }
        }

        vector<int> v1;
        vector<int> v2;
        v1 = {begin(preorder) + 1, begin(preorder) + i + 1};
        v2 = {begin(inorder), begin(inorder) + i};
        node->left = buildTree(v1, v2);

        v1 = {begin(preorder) + i + 1, end(preorder)};
        v2 = {begin(inorder) + i + 1, end(inorder)};
        node->right = buildTree(v1, v2);

        return node;
    }
};

class Solution2 {
public:
    TreeNode *buildTree(vector<int> &preorder, vector<int> &inorder) {
        return buildHelper(preorder, 0, size(preorder), inorder, 0,
                           size(inorder));
    }

private:
    TreeNode *buildHelper(vector<int> &preorder, int b1, int e1,
                          vector<int> &inorder, int b2, int e2) {

        if (b1 >= e1) {
            return nullptr;
        }

        TreeNode *node = new TreeNode(preorder[b1]);

        int i;
        for (i = 0; i < e2 - b2; i++) {
            if (preorder[b1] == inorder[i + b2]) {
                break;
            }
        }

        node->left =
            buildHelper(preorder, b1 + 1, b1 + i + 1, inorder, b2, b2 + i);

        node->right =
            buildHelper(preorder, b1 + i + 1, e1, inorder, b2 + i + 1, e2);

        return node;
    }
};