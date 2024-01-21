#include <stdio.h>
#include <stdlib.h>

struct TreeNode {
    int val;
    struct TreeNode* left;
    struct TreeNode* right;
};

int* newArr(int size, int initVal) {
    int* res = (int*)malloc(size * sizeof(int));
    for (int i = 0; i < size; i++) {
        res[i] = initVal;
    }
    return res;
}

int* helper(struct TreeNode* root, int distance, int* pairsNum) {
    if (root == NULL) {
        return newArr(distance + 1, 0);
    }
    if (root->left == NULL && root->right == NULL) {
        int* res = newArr(distance + 1, 0);
        res[1] += 1;
        return res;
    }

    int* leftDists = helper(root->left, distance, pairsNum);
    int* rightDists = helper(root->right, distance, pairsNum);

    int* res = newArr(distance + 1, 0);
    for (int i = 1; i < distance + 1; i++) {
        for (int j = 1; j < distance + 1; j++) {
            if (i + j <= distance) {
                *pairsNum += leftDists[i] * rightDists[j];
            } else {
                break;
            }
        }
    }

    for (int i = 1; i < distance + 1; i++) {
        res[i] = leftDists[i - 1] + rightDists[i - 1];
    }

    return res;
}

int countPairs(struct TreeNode* root, int distance) {
    int pairsNum = 0;
    helper(root, distance, &pairsNum);
    return pairsNum;
}

int main() {
    struct TreeNode* root;
    int distance;
    int res;

    // 1
    root = &(struct TreeNode){
        .val = 1,
        .left =
            &(struct TreeNode){
                .val = 2,
                .right =
                    &(struct TreeNode){
                        .val = 4,
                    },
            },
        .right =
            &(struct TreeNode){
                .val = 3,
            },
    };
    distance = 3;
    res = countPairs(root, distance);
    printf("%d\n", res);

    // 2
    root = &(struct TreeNode){
        .val = 1,
        .left =
            &(struct TreeNode){
                .val = 2,
                .left =
                    &(struct TreeNode){
                        .val = 4,
                    },
                .right =
                    &(struct TreeNode){
                        .val = 5,
                    },
            },
        .right =
            &(struct TreeNode){
                .val = 3,
                .left =
                    &(struct TreeNode){
                        .val = 6,
                    },
                .right =
                    &(struct TreeNode){
                        .val = 7,
                    },
            },
    };
    distance = 3;
    res = countPairs(root, distance);
    printf("%d\n", res);

    // 1
    root = &(struct TreeNode){
        .val = 7,
        .left =
            &(struct TreeNode){
                .val = 1,
                .left =
                    &(struct TreeNode){
                        .val = 6,
                    },
            },
        .right =
            &(struct TreeNode){
                .val = 4,
                .left =
                    &(struct TreeNode){
                        .val = 5,
                    },
                .right =
                    &(struct TreeNode){
                        .val = 3,
                        .right =
                            &(struct TreeNode){
                                .val = 2,
                            },
                    },
            },
    };
    distance = 3;
    res = countPairs(root, distance);
    printf("%d\n", res);

    // 3
    root = &(struct TreeNode){
        .val = 15,
        .left =
            &(struct TreeNode){
                .val = 66,
                .left =
                    &(struct TreeNode){
                        .val = 97,
                        .right =
                            &(struct TreeNode){
                                .val = 54,
                            },
                    },
                .right =
                    &(struct TreeNode){
                        .val = 60,
                        .right =
                            &(struct TreeNode){
                                .val = 49,
                                .right =
                                    &(struct TreeNode){
                                        .val = 90,
                                    },
                            },
                    },
            },
        .right =
            &(struct TreeNode){
                .val = 55,
                .left =
                    &(struct TreeNode){
                        .val = 12,
                        .right =
                            &(struct TreeNode){
                                .val = 9,
                            },
                    },
                .right =
                    &(struct TreeNode){
                        .val = 56,
                    },
            },
    };
    distance = 5;
    res = countPairs(root, distance);
    printf("%d\n", res);

    return 0;
}