use std::cell::RefCell;
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
    pub fn inorder_traversal(root: Option<Rc<RefCell<TreeNode>>>) -> Vec<i32> {
        let mut result = Vec::new();
        Self::dfs(root, &mut result);
        result
    }

    fn dfs(node: Option<Rc<RefCell<TreeNode>>>, result: &mut Vec<i32>) {
        let Some(node) = node else {
            return;
        };

        let (left, val, right) = {
            let root = node.borrow();
            (root.left.clone(), root.val, root.right.clone())
        };

        Self::dfs(left, result);
        result.push(val);
        Self::dfs(right, result);
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
    expected: Vec<i32>,
}

fn main() {
    let test_cases = vec![
        TestCase {
            input: tree!(1, None, tree!(2, tree!(3), None)),
            expected: vec![1, 3, 2],
        },
        TestCase {
            input: tree!(
                1,
                tree!(2, tree!(4), tree!(5, tree!(6), tree!(7))),
                tree!(3, None, tree!(8, tree!(9), None))
            ),
            expected: vec![4, 2, 6, 5, 7, 1, 3, 9, 8],
        },
        TestCase {
            input: None,
            expected: vec![],
        },
        TestCase {
            input: tree!(1),
            expected: vec![1],
        },
    ];

    for tc in test_cases {
        let actual = Solution::inorder_traversal(tc.input.clone());
        assert_eq!(actual, tc.expected);
    }
}
