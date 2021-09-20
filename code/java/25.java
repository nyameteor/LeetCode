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
            ListNode[] result = reverse(h, t);
            h = result[0];
            t = result[1];
            p.next = h;
            t.next = n;

            p = t;
            h = p.next;
        }

        return root.next;
    }

    // reverse the linked list in a particular range.
    private ListNode[] reverse(ListNode head, ListNode tail) {
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

        // a -> b -> c -> d -> e
        ListNode e = new ListNode(5);
        ListNode d = new ListNode(4, e);
        ListNode c = new ListNode(3, d);
        ListNode b = new ListNode(2, c);
        ListNode a = new ListNode(1, b);

        ListNode h = solution.reverseKGroup(a, 2);

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