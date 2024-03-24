#include <stdlib.h>

struct ListNode {
    int val;
    struct ListNode* next;
};

struct ListNode* findMiddle(struct ListNode* head) {
    struct ListNode* slow = head;
    struct ListNode* fast = head;
    while (fast && fast->next) {
        slow = slow->next;
        fast = fast->next->next;
    }
    return slow;
}

struct ListNode* reverseList(struct ListNode* head) {
    struct ListNode* prev = NULL;
    struct ListNode* curr = head;
    struct ListNode* next;
    while (curr) {
        next = curr->next;
        curr->next = prev;
        prev = curr;
        curr = next;
    }
    return prev;
}

void reorderList(struct ListNode* head) {
    struct ListNode* mid = findMiddle(head);
    struct ListNode* p = head;
    struct ListNode* q = reverseList(mid);
    struct ListNode* next;
    int i = 0;
    while (p != mid && q != NULL) {
        if (i % 2 == 0) {
            next = p->next;
            p->next = q;
            p = next;
        } else {
            next = q->next;
            q->next = p;
            q = next;
        }
        i++;
    }
}
