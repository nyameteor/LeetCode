use std::collections::VecDeque;

struct Solution;

impl Solution {
    pub fn max_sliding_window(nums: Vec<i32>, k: i32) -> Vec<i32> {
        let k = k as usize;
        let mut result = Vec::with_capacity(nums.len() - k + 1);
        let mut deque = VecDeque::new();

        for i in 0..nums.len() {
            if deque.front().is_some_and(|&j| j + k <= i) {
                deque.pop_front();
            }

            while deque.back().is_some_and(|&j| nums[j] <= nums[i]) {
                deque.pop_back();
            }

            deque.push_back(i);

            if i >= k - 1 {
                result.push(nums[deque[0]]);
            }
        }

        result
    }
}

struct TestCase {
    input: (Vec<i32>, i32),
    expected: Vec<i32>,
}

fn main() {
    let test_cases = vec![
        TestCase {
            input: (vec![1, 3, -1, -3, 5, 3, 6, 7], 3),
            expected: vec![3, 3, 5, 5, 6, 7],
        },
        TestCase {
            input: (vec![1], 1),
            expected: vec![1],
        },
        TestCase {
            input: (vec![3, 1, 1, 3], 3),
            expected: vec![3, 3],
        },
        TestCase {
            input: (vec![3, 1, 2], 2),
            expected: vec![3, 2],
        },
        TestCase {
            input: (vec![1, 2, 3, 4], 2),
            expected: vec![2, 3, 4],
        },
        TestCase {
            input: (vec![9, 8, 7, 1], 3),
            expected: vec![9, 8],
        },
        TestCase {
            input: (vec![2, 2, 2, 2], 3),
            expected: vec![2, 2],
        },
        TestCase {
            input: (vec![5, 4, 3, 2, 1], 5),
            expected: vec![5],
        },
        TestCase {
            input: (vec![4, -2, 5], 1),
            expected: vec![4, -2, 5],
        },
    ];

    for tc in test_cases {
        let (nums, k) = tc.input;
        let actual = Solution::max_sliding_window(nums, k);
        assert_eq!(actual, tc.expected);
    }
}
