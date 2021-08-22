/**
 * Definition for singly-linked list.
 */
class ListNode {
  int val;
  ListNode next;

  ListNode() {
  }

  ListNode(int val) {
    this.val = val;
  }

  ListNode(int val, ListNode next) {
    this.val = val;
    this.next = next;
  }
}

class Solution {
  public ListNode reverseList(ListNode head) {
    if (head == null)
      return null;

    ListNode backword = head;
    ListNode current = head.next;
    ListNode forward = current;

    while (current != null) {
      forward = current.next;

      // reverse
      current.next = backword;

      // edge case, to avoid cycle
      if (backword == head)
        backword.next = null;

      // move pointers
      backword = current;
      current = forward;
    }

    return backword;
  }
}