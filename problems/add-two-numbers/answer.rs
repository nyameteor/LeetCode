// Definition for singly-linked list.
#[derive(PartialEq, Eq, Clone, Debug)]
pub struct ListNode {
    pub val: i32,
    pub next: Option<Box<ListNode>>,
}

impl ListNode {
    #[inline]
    fn new(val: i32) -> Self {
        ListNode { next: None, val }
    }
}

struct Solution;

impl Solution {
    pub fn add_two_numbers(
        l1: Option<Box<ListNode>>,
        l2: Option<Box<ListNode>>,
    ) -> Option<Box<ListNode>> {
        let mut dummy = Box::new(ListNode::new(-1));
        let mut tail = &mut dummy;
        let mut p1 = l1.as_ref();
        let mut p2 = l2.as_ref();
        let mut carry = 0;

        while p1.is_some() || p2.is_some() || carry > 0 {
            let mut sum = carry;

            if let Some(node) = p1 {
                sum += node.val;
                p1 = node.next.as_ref();
            }

            if let Some(node) = p2 {
                sum += node.val;
                p2 = node.next.as_ref();
            }

            carry = sum / 10;

            tail.next = Some(Box::new(ListNode::new(sum % 10)));
            tail = tail.next.as_mut().unwrap();
        }

        dummy.next
    }
}

macro_rules! list {
    ($val:expr) => {
        Some(Box::new(ListNode {
            val: $val,
            next: None,
        }))
    };
    ($val:expr, $next:expr) => {
        Some(Box::new(ListNode {
            val: $val,
            next: $next,
        }))
    };
}

struct TestCase {
    input: (Option<Box<ListNode>>, Option<Box<ListNode>>),
    expected: Option<Box<ListNode>>,
}

fn main() {
    let test_cases = vec![
        TestCase {
            input: (list!(2, list!(4, list!(3))), list!(5, list!(6, list!(4)))),
            expected: list!(7, list!(0, list!(8))),
        },
        TestCase {
            input: (list!(0), list!(0)),
            expected: list!(0),
        },
        TestCase {
            input: (
                list!(
                    9,
                    list!(9, list!(9, list!(9, list!(9, list!(9, list!(9))))))
                ),
                list!(9, list!(9, list!(9, list!(9)))),
            ),
            expected: list!(
                8,
                list!(
                    9,
                    list!(9, list!(9, list!(0, list!(0, list!(0, list!(1))))))
                )
            ),
        },
    ];

    for tc in test_cases {
        let (l1, l2) = tc.input;
        let actual = Solution::add_two_numbers(l1, l2);

        assert_eq!(actual, tc.expected);
    }
}
