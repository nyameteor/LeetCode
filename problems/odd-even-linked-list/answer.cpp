#include <iostream>
#include <vector>
using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

class Solution {
public:
    ListNode *oddEvenList(ListNode *head) {
        if (!head || !head->next) {
            return head;
        }

        ListNode *p = head;
        ListNode *q = head->next, *evenHead = q;
        while (q && q->next) {
            p->next = p->next->next;
            p = p->next;

            q->next = q->next->next;
            q = q->next;
        }
        p->next = evenHead;

        return head;
    }
};

/**
 * Refer: @archit91
 * https://leetcode.com/discuss/topic/1607746
 *
 * Same as first solution, but more elegant
 */
class Solution2 {
    ListNode *oddEvenList(ListNode *head) {
        if (!head || !head->next) {
            return head;
        }

        ListNode *p = head;
        ListNode *q = head->next, *evenHead = q;
        while (q && q->next) {
            p = p->next = q->next;
            q = q->next = p->next;
        }
        p->next = evenHead;

        return head;
    }
};

void printLinkedList(ListNode *head) {
    ListNode *p = head;
    while (p) {
        cout << p->val << ' ';
        p = p->next;
    }
    cout << endl;
}

ListNode *createLinkedList(vector<int> nums) {
    ListNode *dummy = new ListNode(-1);
    ListNode *p = dummy;
    for (auto x : nums) {
        p->next = new ListNode(x);
        p = p->next;
    }
    return dummy->next;
}

int main() {
    vector<int> nums;
    Solution solution;

    nums = {1, 2, 3, 4, 5};
    printLinkedList(solution.oddEvenList(createLinkedList(nums)));

    nums = {2, 1, 3, 5, 6, 4, 7};
    printLinkedList(solution.oddEvenList(createLinkedList(nums)));

    return 0;
}