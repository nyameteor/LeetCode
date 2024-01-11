#include <stdio.h>
#include <stdlib.h>

struct TreeNode {
    int val;
    struct TreeNode* left;
    struct TreeNode* right;
};

void inorder(struct TreeNode* root, int* vals, int* valsSize) {
    if (root == NULL) {
        return;
    }
    inorder(root->left, vals, valsSize);
    *(vals + *valsSize) = root->val;
    *valsSize += 1;
    inorder(root->right, vals, valsSize);
}

struct TreeNode* buildBST(int* vals, int l, int r) {
    if (l > r) {
        return NULL;
    }
    int m = l + (r - l) / 2;
    struct TreeNode* res = (struct TreeNode*)malloc(sizeof(struct TreeNode));
    res->val = *(vals + m);
    res->left = buildBST(vals, l, m - 1);
    res->right = buildBST(vals, m + 1, r);
    return res;
}

struct TreeNode* balanceBST(struct TreeNode* root) {
    int max_size = 10000;
    int* vals = (int*)malloc(max_size * sizeof(int));
    int valsSize = 0;
    inorder(root, vals, &valsSize);
    return buildBST(vals, 0, valsSize - 1);
}

int main() {
    struct TreeNode* root;

    root = &(struct TreeNode){
        .val = 1,
        .right =
            &(struct TreeNode){
                .val = 2,
                .right =
                    &(struct TreeNode){
                        .val = 3,
                        .right =
                            &(struct TreeNode){
                                .val = 4,
                            },
                    },
            },
    };
    struct TreeNode* res = balanceBST(root);

    return 0;
}