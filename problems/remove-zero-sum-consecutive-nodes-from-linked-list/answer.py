from typing import Optional


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next

    def __repr__(self) -> str:
        return f'({self.val}, {self.next})'


class Solution:
    def removeZeroSumSublists(self, head: Optional[ListNode]) -> Optional[ListNode]:
        dummy = ListNode(0, head)
        prefix = 0
        seen = {0: dummy}
        while head:
            prefix += head.val
            seen[prefix] = head
            head = head.next

        head = dummy
        prefix = 0
        while head:
            prefix += head.val
            head.next = seen[prefix].next
            head = head.next

        return dummy.next


if __name__ == "__main__":
    solution = Solution()

    head = ListNode(1, ListNode(2, ListNode(-3, ListNode(3, ListNode(1)))))
    res = solution.removeZeroSumSublists(head)
    # [3,1]
    print(res)

    head = ListNode(1, ListNode(2, ListNode(3, ListNode(-3, ListNode(4)))))
    res = solution.removeZeroSumSublists(head)
    # [1,2,4]
    print(res)

    head = ListNode(1, ListNode(2, ListNode(3, ListNode(-3, ListNode(-2)))))
    res = solution.removeZeroSumSublists(head)
    # [1]
    print(res)
