class Solution {
    public ListNode reverseKGroup(ListNode head, int k) {
        // corner case
        if (k == 1) {
            return head;
        }

        // add dummy head
        ListNode root = new ListNode(-1);
        root.next = head;

        // previous, next
        ListNode p = root;
        ListNode n = p;
        // sublist head, tail
        ListNode h, t;

        t = p;
        while (t != null) {
            for (int i = 0; i < k; i++) {
                t = t.next;
                if (t == null) {
                    return root.next;
                }
            }
            h = p.next;
            n = t.next;
            ListNode[] result = reverseSub(h, t);
            h = result[0];
            t = result[1];
            p.next = h;
            t.next = n;

            p = t;
            h = p.next;
        }

        return root.next;
    }

    private ListNode[] reverseSub(ListNode head, ListNode tail) {
        // current, backfoward, forward
        ListNode b = head;
        ListNode c = b.next, f = b.next;

        while (b != tail) {
            f = c.next;
            c.next = b;
            b = c;
            c = f;
        }

        // new reversed list tail will be original head, head will be original tail
        return new ListNode[] { tail, head };
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        ListNode a = new ListNode(1);
        ListNode b = new ListNode(2);
        ListNode c = new ListNode(3);
        a.next = b;
        b.next = c;
        c.next = null;

        ListNode[] result = solution.reverseSub(a, c);
        ListNode h = result[0];
        // avoid linked list become ring
        a.next = null;

        while (h != null) {
            System.out.println(h.val);
            h = h.next;
        }
    }
}

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