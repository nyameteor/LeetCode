#include <stdio.h>
#include <stdlib.h>

struct TreeNode {
    int val;
    struct TreeNode* left;
    struct TreeNode* right;
};

struct TreeNode* copyTree(struct TreeNode* cur) {
    if (cur == NULL) {
        return NULL;
    }
    struct TreeNode* res = (struct TreeNode*)malloc(sizeof(struct TreeNode));
    res->val = cur->val;
    res->left = copyTree(cur->left);
    res->right = copyTree(cur->right);
    return res;
}

void insertHelper(struct TreeNode* cur, struct TreeNode* pre, int val) {
    if (cur == NULL || cur->val < val) {
        struct TreeNode* new =
            (struct TreeNode*)malloc(sizeof(struct TreeNode));
        new->val = val;
        new->left = cur;
        new->right = NULL;
        pre->right = new;
        return;
    }
    insertHelper(cur->right, cur, val);
}

struct TreeNode* insertIntoMaxTree(struct TreeNode* root, int val) {
    root = copyTree(root);
    struct TreeNode* dummy = (struct TreeNode*)malloc(sizeof(struct TreeNode));
    dummy->val = -1;
    dummy->left = NULL;
    dummy->right = root;
    insertHelper(root, dummy, val);
    return dummy->right;
}

int main() {
    struct TreeNode* root;
    int val;
    struct TreeNode* res;

    root = &(struct TreeNode){
        .val = 4,
        .left =
            &(struct TreeNode){
                .val = 1,
            },
        .right =
            &(struct TreeNode){
                .val = 3,
                .left =
                    &(struct TreeNode){
                        .val = 2,
                    },
            },
    };
    val = 5;
    res = insertIntoMaxTree(root, val);

    root = &(struct TreeNode){
        .val = 5,
        .left =
            &(struct TreeNode){
                .val = 2,
                .right =
                    &(struct TreeNode){
                        .val = 1,
                    },
            },
        .right =
            &(struct TreeNode){
                .val = 4,
            },
    };
    val = 3;
    res = insertIntoMaxTree(root, val);

    root = &(struct TreeNode){
        .val = 5,
        .left =
            &(struct TreeNode){
                .val = 2,
                .right =
                    &(struct TreeNode){
                        .val = 1,
                    },
            },
        .right =
            &(struct TreeNode){
                .val = 3,
            },
    };
    val = 4;
    res = insertIntoMaxTree(root, val);

    return 0;
}