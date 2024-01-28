#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>

struct Node {
    int data;
    struct Node* next;
};

struct Stack {
    struct Node* head;
    int size;
};

struct Stack* stackInit() {
    struct Stack* s = (struct Stack*)malloc(sizeof(struct Stack));
    s->head = NULL;
    s->size = 0;
    return s;
}

bool stackEmpty(struct Stack* s) {
    return s->size == 0;
}

void stackPush(struct Stack* s, int data) {
    struct Node* tmp = (struct Node*)malloc(sizeof(struct Node));
    tmp->data = data;
    tmp->next = s->head;
    s->head = tmp;
    s->size += 1;
}

int stackPop(struct Stack* s) {
    if (stackEmpty(s)) {
        printf("Error: stack is empty");
        exit(1);
    }
    struct Node* tmp = s->head;
    s->head = s->head->next;
    s->size -= 1;
    int data = tmp->data;
    free(tmp);
    return data;
}

int stackPeek(struct Stack* s) {
    if (stackEmpty(s)) {
        printf("Error: stack is empty");
        exit(1);
    }
    return s->head->data;
}

void freeStack(struct Stack* s) {
    for (struct Node* p = s->head; p != NULL;) {
        struct Node* tmp = p;
        p = p->next;
        free(tmp);
    }
    free(s);
}

bool sameSign(int a, int b) {
    return (a ^ b) >= 0;
}

int abs(int a) {
    return a >= 0 ? a : -a;
}

void helper(int* asteroids, int asteroidsSize, int i, struct Stack* s) {
    if (i >= asteroidsSize) {
        return;
    }
    int curData = asteroids[i];
    if (stackEmpty(s)) {
        stackPush(s, curData);
        helper(asteroids, asteroidsSize, i + 1, s);
        return;
    }
    int headData = stackPeek(s);
    if (sameSign(headData, curData) || headData < 0 && curData > 0) {
        stackPush(s, curData);
        helper(asteroids, asteroidsSize, i + 1, s);
    } else if (abs(curData) < abs(headData)) {
        helper(asteroids, asteroidsSize, i + 1, s);
    } else if (abs(curData) == abs(headData)) {
        stackPop(s);
        helper(asteroids, asteroidsSize, i + 1, s);
    } else {
        stackPop(s);
        helper(asteroids, asteroidsSize, i, s);
    }
}

int* asteroidCollision(int* asteroids, int asteroidsSize, int* returnSize) {
    struct Stack* s = stackInit();
    helper(asteroids, asteroidsSize, 0, s);

    int resSize = s->size;
    int* res = (int*)malloc(resSize * sizeof(int));
    for (int i = resSize - 1; i >= 0; i--) {
        res[i] = stackPop(s);
    }
    *returnSize = resSize;

    freeStack(s);
    return res;
}

void printArray(int* arr, int size) {
    for (int i = 0; i < size; i++) {
        printf("%d ", arr[i]);
    }
    printf("\n");
}

int main() {
    int* asteroids;
    int asteroidsSize;
    int* res;
    int returnSize;

    /* Output: [5,10] */
    asteroids = (int[]){5, 10, -5};
    asteroidsSize = 3;
    res = asteroidCollision(asteroids, asteroidsSize, &returnSize);
    printArray(res, returnSize);
    free(res);

    /* Output: [] */
    asteroids = (int[]){8, -8};
    asteroidsSize = 2;
    res = asteroidCollision(asteroids, asteroidsSize, &returnSize);
    printArray(res, returnSize);
    free(res);

    /* Output: [10] */
    asteroids = (int[]){10, 2, -5};
    asteroidsSize = 3;
    res = asteroidCollision(asteroids, asteroidsSize, &returnSize);
    printArray(res, returnSize);
    free(res);

    return 0;
}