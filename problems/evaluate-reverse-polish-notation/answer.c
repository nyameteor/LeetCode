#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

struct Node {
    char* data;
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

void stackPush(struct Stack* s, char* data) {
    struct Node* tmp = (struct Node*)malloc(sizeof(struct Node));
    tmp->data = data;
    tmp->next = s->head;
    s->head = tmp;
    s->size += 1;
}

char* stackPop(struct Stack* s) {
    if (stackEmpty(s)) {
        printf("Error: stack is empty.");
        exit(1);
    }
    struct Node* tmp = s->head;
    s->head = s->head->next;
    s->size -= 1;
    char* data = tmp->data;
    free(tmp);
    return data;
}

char* stackPeek(struct Stack* s) {
    if (stackEmpty(s)) {
        printf("Error: stack is empty.");
        exit(1);
    }
    return s->head->data;
}

void freeStack(struct Stack* s) {
    struct Node* p = s->head;
    while (p != NULL) {
        struct Node* tmp = p;
        p = p->next;
        free(tmp);
    }
    free(s);
}

char* int_to_str(int i) {
    int max_len = 12;
    char* s = (char*)malloc(max_len * sizeof(char));
    sprintf(s, "%d", i);
    return s;
}

int evalRPN(char** tokens, int tokensSize) {
    struct Stack* s = stackInit();
    for (int i = 0; i < tokensSize; i++) {
        char* token = tokens[i];
        if (!strcmp(token, "+") || !strcmp(token, "-") || !strcmp(token, "*") ||
            !strcmp(token, "/")) {
            int num2 = atoi(stackPop(s));
            int num1 = atoi(stackPop(s));
            int num = 0;
            if (token[0] == '+') {
                num = num1 + num2;
            } else if (token[0] == '-') {
                num = num1 - num2;
            } else if (token[0] == '*') {
                num = num1 * num2;
            } else {
                num = num1 / num2;
            }
            stackPush(s, int_to_str(num));
        } else {
            stackPush(s, token);
        }
    }
    int res = atoi(stackPop(s));
    freeStack(s);
    return res;
}

int main() {
    char** tokens;
    int tokenSize;
    int res;

    /* Output: 9 */
    tokens = (char*[]){"2", "1", "+", "3", "*"};
    tokenSize = 5;
    res = evalRPN(tokens, tokenSize);
    printf("%d\n", res);

    /* Output: 6 */
    tokens = (char*[]){"4", "13", "5", "/", "+"};
    tokenSize = 5;
    res = evalRPN(tokens, tokenSize);
    printf("%d\n", res);

    /* Output: 22 */
    tokens = (char*[]){"10", "6", "9",  "3", "+", "-11", "*",
                       "/",  "*", "17", "+", "5", "+"};
    tokenSize = 13;
    res = evalRPN(tokens, tokenSize);
    printf("%d\n", res);

    return 0;
}