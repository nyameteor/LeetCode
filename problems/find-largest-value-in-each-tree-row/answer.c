#include <limits.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

struct Node {
    void* data;
    struct Node* next;
};

struct Queue {
    struct Node* head;
    struct Node* tail;
    int size;
};

struct Queue* initQueue() {
    struct Queue* q = (struct Queue*)malloc(sizeof(struct Queue));
    q->head = NULL;
    q->tail = NULL;
    q->size = 0;
    return q;
}

bool isQueueEmpty(struct Queue* q) {
    return q->size == 0;
}

void enqueue(struct Queue* q, void* data) {
    struct Node* cur = (struct Node*)malloc(sizeof(struct Node));
    cur->data = data;
    cur->next = NULL;
    if (q->tail == NULL) {
        q->head = cur;
        q->tail = cur;
    } else {
        q->tail->next = cur;
        q->tail = cur;
    }
    q->size += 1;
}

void* dequeue(struct Queue* q) {
    if (isQueueEmpty(q)) {
        return NULL;
    }
    struct Node* first = q->head;
    q->head = q->head->next;
    if (q->head == NULL) {
        q->tail = NULL;
    }
    q->size -= 1;
    return first->data;
}

void freeQueue(struct Queue* q) {
    for (struct Node* p = q->head; p != NULL;) {
        struct Node* tmp = p;
        p = p->next;
        free(tmp);
    }
    free(q);
}

struct TreeNode {
    int val;
    struct TreeNode* left;
    struct TreeNode* right;
};

int* largestValues(struct TreeNode* root, int* returnSize) {
    int* res = NULL;
    if (root == NULL) {
        *returnSize = 0;
        return res;
    }

    res = (int*)malloc(10000 * sizeof(int));
    struct Queue* q = initQueue();
    enqueue(q, root);
    int level = 0;
    while (!isQueueEmpty(q)) {
        int size = q->size;
        int maxVal = INT_MIN;
        for (int i = 0; i < size; i++) {
            struct TreeNode* cur = (struct TreeNode*)dequeue(q);
            if (cur->val > maxVal) {
                maxVal = cur->val;
            }
            if (cur->left != NULL) {
                enqueue(q, cur->left);
            }
            if (cur->right != NULL) {
                enqueue(q, cur->right);
            }
        }
        res[level] = maxVal;
        level += 1;
    }
    freeQueue(q);
    *returnSize = level;
    return res;
}

void printResult(int* res, int resSize) {
    for (int i = 0; i < resSize; i++) {
        printf("%d ", res[i]);
    }
    printf("\n");
}

int main() {
    struct TreeNode* root;
    int returnSize;
    int* res;

    // [1,3,9]
    root = &(struct TreeNode){
        .val = 1,
        .left =
            &(struct TreeNode){
                .val = 3,
                .left =
                    &(struct TreeNode){
                        .val = 5,
                    },
                .right =
                    &(struct TreeNode){
                        .val = 3,
                    },
            },
        .right =
            &(struct TreeNode){
                .val = 2,
                .right =
                    &(struct TreeNode){
                        .val = 9,
                    },
            },
    };
    res = largestValues(root, &returnSize);
    printResult(res, returnSize);

    // [1,3]
    root = &(struct TreeNode){
        .val = 1,
        .left =
            &(struct TreeNode){
                .val = 2,
            },
        .right =
            &(struct TreeNode){
                .val = 3,
            },
    };
    res = largestValues(root, &returnSize);
    printResult(res, returnSize);

    return 0;
}