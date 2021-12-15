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

/**
 * Sort by swapping nodes
 */
class Solution {
  public:
    ListNode *insertionSortList(ListNode *head) {
        auto dummy = new ListNode(0, head);
        auto q = dummy, pPrev = head, p = head->next;
        while (p) {
            q = dummy;
            while (q->next != p) {
                if (q->next->val > p->val) {
                    break;
                } else {
                    q = q->next;
                }
            }
            if (q->next == p) {
                pPrev = p;
                p = p->next;
            } else {
                auto pNext = p->next;
                p->next = q->next;
                q->next = p;
                pPrev->next = pNext;
                p = pNext;
            }
        }
        return dummy->next;
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
    Solution solution;
    vector<int> nums;

    nums = {-1, 5, 3, 4, 0};
    printLinkedList(solution.insertionSortList(createLinkedList(nums)));

    nums = {2, 4, 1, 3, 0};
    printLinkedList(solution.insertionSortList(createLinkedList(nums)));

    return 0;
}