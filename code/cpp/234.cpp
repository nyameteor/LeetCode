#include <iostream>
#include <queue>
#include <stack>
#include <vector>

using namespace std;

struct ListNode {
    int val;
    ListNode *next;
    ListNode() : val(0), next(nullptr) {}
    ListNode(int x) : val(x), next(nullptr) {}
    ListNode(int x, ListNode *next) : val(x), next(next) {}
};

// Reverse the second half of the linked list
class Solution {
public:
    bool isPalindrome(ListNode *head) {
        ListNode *pMid = findMiddle(head);
        ListNode *pRev = reverseLinkedList(pMid);

        while (head != pMid) {
            if (head->val != pRev->val) {
                return false;
            }
            head = head->next;
            pRev = pRev->next;
        }

        return true;
    }

    // find middle node using two pointers
    ListNode *findMiddle(ListNode *head) {
        ListNode *slow = head, *fast = head;
        while (fast && fast->next) {
            slow = slow->next;
            fast = fast->next->next;
        }
        return slow;
    }

    ListNode *reverseLinkedList(ListNode *head) {
        ListNode *p = nullptr;

        while (head) {
            ListNode *q = head->next;
            head->next = p;
            p = head;
            head = q;
        }

        return p;
    }
};

// Use an array to simulate
class Solution2 {
public:
    bool isPalindrome(ListNode *head) {
        ListNode *p = head;
        int length = 0;

        // get the length of linked list
        while (p) {
            length++;
            p = p->next;
        }

        // copy linked list to vector
        vector<int> list(length, 0);
        p = head;
        int i = 0;
        while (p) {
            list[i] = p->val;
            i++;
            p = p->next;
        }

        for (int i = 0, j = list.size() - 1; i < j; i++, j--) {
            if (list[i] != list[j]) {
                return false;
            }
        }

        return true;
    }
};

int main() {
    Solution solution;
    Solution2 solution2;

    // 1 -> 2 -> 2 -> 1
    ListNode *d = new ListNode(1);
    ListNode *c = new ListNode(2, d);
    ListNode *b = new ListNode(2, c);
    ListNode *a = new ListNode(1, b);
    cout << solution2.isPalindrome(a) << endl;
    cout << solution.isPalindrome(a) << endl;

    // 1 -> 2
    ListNode *f = new ListNode(2);
    ListNode *e = new ListNode(1, f);
    cout << solution2.isPalindrome(e) << endl;
    cout << solution.isPalindrome(e) << endl;
}