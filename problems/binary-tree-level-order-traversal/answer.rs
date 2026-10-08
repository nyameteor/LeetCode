use std::cell::RefCell;
use std::collections::VecDeque;
use std::rc::Rc;

// Definition for a binary tree node.
#[derive(Debug, PartialEq, Eq)]
pub struct TreeNode {
    pub val: i32,
    pub left: Option<Rc<RefCell<TreeNode>>>,
    pub right: Option<Rc<RefCell<TreeNode>>>,
}

impl TreeNode {
    #[inline]
    pub fn new(val: i32) -> Self {
        TreeNode {
            val,
            left: None,
            right: None,
        }
    }
}

struct Solution;

impl Solution {
    pub fn level_order(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<Vec<i32>> {
        let mut result = Vec::new();

        let Some(root) = root else { return result };

        let mut queue = VecDeque::from([root]);

        while !queue.is_empty() {
            let level_size = queue.len();
            let mut level_result = Vec::<i32>::new();

            for _ in 0..level_size {
                let node = queue.pop_front().unwrap();
                let node = node.borrow();

                level_result.push(node.val);
                if let Some(left) = node.left.clone() {
                    queue.push_back(left);
                }
                if let Some(right) = node.right.clone() {
                    queue.push_back(right);
                }
            }
            result.push(level_result);
        }

        result
    }
}

macro_rules! tree {
    ($val:expr) => {
        Some(Rc::new(RefCell::new(TreeNode {
            val: $val,
            left: None,
            right: None,
        })))
    };
    ($val:expr, $left:expr, $right:expr) => {
        Some(Rc::new(RefCell::new(TreeNode {
            val: $val,
            left: $left,
            right: $right,
        })))
    };
}

struct TestCase {
    input: Option<Rc<RefCell<TreeNode>>>,
    expected: Vec<Vec<i32>>,
}

fn main() {
    let test_cases = vec![
        TestCase {
            input: tree!(3, tree!(9), tree!(20, tree!(15), tree!(7))),
            expected: vec![vec![3], vec![9, 20], vec![15, 7]],
        },
        TestCase {
            input: tree!(1),
            expected: vec![vec![1]],
        },
        TestCase {
            input: None,
            expected: vec![],
        },
    ];

    for tc in test_cases {
        let actual = Solution::level_order(tc.input);
        assert_eq!(actual, tc.expected);
    }
}
