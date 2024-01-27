#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

struct TreeNode {
    int val;
    struct TreeNode* left;
    struct TreeNode* right;
};

void helper(struct TreeNode* root, int* valCounts, int* res) {
    if (root == NULL) {
        return;
    }

    valCounts[root->val] += 1;

    if (root->left == NULL && root->right == NULL) {
        int oddCount = 0;
        for (int i = 1; i <= 9; i++) {
            if (valCounts[i] % 2 == 1) {
                oddCount += 1;
            }
        }
        if (oddCount <= 1) {
            *res += 1;
        }
    }

    helper(root->left, valCounts, res);
    helper(root->right, valCounts, res);

    valCounts[root->val] -= 1;
}

int pseudoPalindromicPaths(struct TreeNode* root) {
    int* valCounts = (int*)malloc(10 * sizeof(int));
    for (int i = 0; i < 10; i++) {
        valCounts[i] = 0;
    }

    int res = 0;
    helper(root, valCounts, &res);

    free(valCounts);
    return res;
}

int main() {
    struct TreeNode* root;
    int res;

    /* Output: 2 */
    root = &(struct TreeNode){
        .val = 2,
        .left =
            &(struct TreeNode){
                .val = 3,
                .left =
                    &(struct TreeNode){
                        .val = 3,
                    },
                .right =
                    &(struct TreeNode){
                        .val = 1,
                    },
            },
        .right =
            &(struct TreeNode){
                .val = 1,
                .right =
                    &(struct TreeNode){
                        .val = 1,
                    },
            },
    };
    res = pseudoPalindromicPaths(root);
    printf("%d\n", res);

    /* Output: 1 */
    root = &(struct TreeNode){
        .val = 2,
        .left =
            &(struct TreeNode){
                .val = 1,
                .left =
                    &(struct TreeNode){
                        .val = 1,
                    },
                .right =
                    &(struct TreeNode){
                        .val = 3,
                        .right =
                            &(struct TreeNode){
                                .val = 1,
                            },
                    },
            },
        .right =
            &(struct TreeNode){
                .val = 1,
            },
    };
    res = pseudoPalindromicPaths(root);
    printf("%d\n", res);

    return 0;
}