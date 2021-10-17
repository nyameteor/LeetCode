#include <iostream>
#include <vector>
using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode(int x) : val(x), next(NULL) {}
};

/**
 * Two Pointers
 *
 * Use slow(1 step) and fast(2 step) pointers to traversal the linked list at
 * the same time, if these two pointers meet at a specific node, that means
 * the linked list has a cycle in it.
 */
class Solution {
  public:
    bool hasCycle(ListNode *head) {
        ListNode *p = head, *q = head;
        while (p != nullptr && q != nullptr && q->next != nullptr) {
            p = p->next;
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

int main() {
    Solution solution;
    ListNode *head;

    head = createLinkedListWithCycle({3, 2, 0, -4}, 1);
    cout << solution.hasCycle(head) << endl;

    head = createLinkedListWithCycle({1, 2}, 0);
    cout << solution.hasCycle(head) << endl;

    head = createLinkedListWithCycle({1}, -1);
    cout << solution.hasCycle(head) << endl;

    return 0;
}