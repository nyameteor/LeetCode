#include <algorithm>
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
 * Natural merge sort
 *
 * Refer: https://en.wikipedia.org/wiki/Merge_sort#Natural_merge_sort
 *
 * TODO: optimization
 */
class Solution {
  public:
    ListNode *sortList(ListNode *head) {
        if (!head) {
            return nullptr;
        }

        vector<ListNode *> nodes = selectIncrease(head);
        while (true) {
            if (size(nodes) == 1) {
                break;
            }
            vector<ListNode *> newNodes;
            for (int i = 0; i < size(nodes); i += 2) {
                if (i + 1 < size(nodes)) {
                    newNodes.push_back(merge(nodes[i], nodes[i + 1]));
                } else {
                    newNodes.push_back(nodes[i]);
                }
            }
            nodes = newNodes;
        }
        return nodes[0];
    }

    vector<ListNode *> selectIncrease(ListNode *A) {
        vector<ListNode *> nodes;
        ListNode *p = A;
        ListNode *dummy = new ListNode(INT_MAX, A);
        ListNode *d = dummy;
        int prev = dummy->val;
        int cur = p->val;
        while (p) {
            cur = p->val;
            if (prev > cur) {
                d->next = nullptr;
                nodes.push_back(p);
            }
            prev = cur;
            d = p;
            p = p->next;
        }

        return nodes;
    }

    ListNode *merge(ListNode *A, ListNode *B) {
        ListNode *p1 = A;
        ListNode *p2 = B;
        ListNode *dummy = new ListNode(-1);
        ListNode *p3 = dummy;

        while (p1 && p2) {
            if (p1->val <= p2->val) {
                p3->next = p1;
                p1 = p1->next;
                p3 = p3->next;
            } else {
                p3->next = p2;
                p2 = p2->next;
                p3 = p3->next;
            }
        }

        if (p1 && !p2) {
            while (p1) {
                p3->next = p1;
                p1 = p1->next;
                p3 = p3->next;
            }
        } else if (!p1 && p2) {
            while (p2) {
                p3->next = p2;
                p2 = p2->next;
                p3 = p3->next;
            }
        }

        return dummy->next;
    }
};

/**
 * Use array to simulate
 */
class Solution2 {
  public:
    ListNode *sortList(ListNode *head) {
        vector<int> v;

        ListNode *p = head;
        while (p) {
            v.push_back(p->val);
            p = p->next;
        }
        sort(v.begin(), v.end());

        p = head;
        int i = 0;
        while (p) {
            p->val = v[i++];
            p = p->next;
        }

        return head;
    }
};

ListNode *createLinkedList(vector<int> nums) {
    ListNode *dummy = new ListNode(-1);
    ListNode *p = dummy;
    for (auto x : nums) {
        p->next = new ListNode(x);
        p = p->next;
    }
    return dummy->next;
}

void printLinkedList(ListNode *head) {
    ListNode *p = head;
    while (p) {
        cout << p->val << ' ';
        p = p->next;
    }
    cout << endl;
}

int main() {
    Solution solution;
    vector<int> nums;

    nums = {3, 4, 2, 1, 7, 5, 8, 9, 0, 6};
    printLinkedList(solution.sortList(createLinkedList(nums)));

    nums = {-1, 5, 3, 4, 0};
    printLinkedList(solution.sortList(createLinkedList(nums)));

    nums = {};
    printLinkedList(solution.sortList(createLinkedList(nums)));
}