class ListNode {
  int val;
  ListNode next;

  ListNode(int x) {
    val = x;
    next = null;
  }
}

class Solution {
  public ListNode getIntersectionNode(ListNode headA, ListNode headB) {
    // edge case
    if (headA == null || headB == null)
      return null;

    ListNode pA = headA, pB = headB;
    boolean toCombine = true;

    while (pA != pB) {
      if (pA.next == null && toCombine)
        toCombine = false;
      else if (pA.next == null && !toCombine)
        return null;
      pA = (pA.next == null) ? headB : pA.next;
      pB = (pB.next == null) ? headA : pB.next;
    }

    return pA;
  }
}