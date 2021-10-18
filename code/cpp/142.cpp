#include <iostream>
#include <vector>
using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

/**
 * Two Pointers (Brute Force)
 *
 * Idea:
 *  1. Using the cycle-chcking algorithm.
 *  2. Checking node in linked list one by one, if slow and fast pointers meet
 *      at where they start checking, that place is the cycle's start-point
 */
class Solution {
  public:
    ListNode *detectCycle(ListNode *head) {
        while (head != nullptr) {
            ListNode *p = head, *q = head;
            while (p != nullptr && q != nullptr && q->next != nullptr) {
                p = p->next;
                q = q->next->next;
                if (p == q) {
                    break;
                }
            }

            if (p == nullptr || q == nullptr || q->next == nullptr) {
                // no cycle in linked list
                return nullptr;
            } else if (p == head && q == head) {
                // meet at begin node, means it's the cycle begins
                return p;
            } else {
                head = head->next;
            }
        }

        return nullptr;
    }
};

/**
 * Two Pointers (Two Pass)
 *
 * Idea:
 *  1. Using the cycle-chcking algorithm.
 *  2. Once p(slow) and q(fast) meet, reset p to head, and move p & q with same
 *      speed until they meet again, the place is the cycle's start-point
 *
 * Theory: todo
 */
class Solution2 {
  private:
    ListNode *p, *q;

  public:
    ListNode *detectCycle(ListNode *head) {
        if (!hasCycle(head)) {
            return nullptr;
        }

        p = head;

        while (p != q) {
            p = p->next;
            q = q->next;
        }

        return p;
    }

    bool hasCycle(ListNode *head) {
        p = head, q = head;
        while (p != nullptr && q != nullptr) {
            p = p->next;

            if (q->next == nullptr) {
                return false;
            }

            q = q->next->next;

            if (p == q) {
                return true;
            }
        }

        return false;
    }
};

ListNode *createLinkedListWithCycle(vector<int> v, int pos) {
    int n = v.size();

    if (n <= 0) {
        return nullptr;
    }

    ListNode **list = new ListNode *[n];

    for (int i = 0; i < n; i++) {
        list[i] = new ListNode(v[i]);
    }
    for (int i = 0; i < n - 1; i++) {
        list[i]->next = list[i + 1];
    }
    if (pos != -1) {
        // connect tail to $(pos)th node
        list[n - 1]->next = list[pos];
    } // -1 means no cycle

    return list[0];
}

void printNodePos(ListNode *head, ListNode *node) {
    int pos = 0;

    while (head != nullptr) {
        if (head == node) {
            cout << pos << endl;
            return;
        }
        pos++;
        head = head->next;
    }

    // not found
    cout << nullptr << endl;
}

int main() {
    Solution solution;
    Solution2 solution2;
    ListNode *head, *res;

    // 1
    head = createLinkedListWithCycle({3, 2, 0, -4}, 1);
    printNodePos(head, solution.detectCycle(head));
    printNodePos(head, solution2.detectCycle(head));

    // 0
    head = createLinkedListWithCycle({1, 2}, 0);
    printNodePos(head, solution.detectCycle(head));
    printNodePos(head, solution2.detectCycle(head));

    // null
    head = createLinkedListWithCycle({1}, -1);
    printNodePos(head, solution.detectCycle(head));
    printNodePos(head, solution2.detectCycle(head));

    // 0
    head = createLinkedListWithCycle({1}, 0);
    printNodePos(head, solution.detectCycle(head));
    printNodePos(head, solution2.detectCycle(head));

    return 0;
}