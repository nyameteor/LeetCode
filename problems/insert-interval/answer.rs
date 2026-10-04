use std::cmp;

struct Solution;

impl Solution {
    pub fn insert(intervals: Vec<Vec<i32>>, new_interval: Vec<i32>) -> Vec<Vec<i32>> {
        let mut left = Vec::<Vec<i32>>::new();
        let mut right = Vec::<Vec<i32>>::new();
        let mut merged = new_interval.clone();

        for interval in intervals {
            if interval[1] < new_interval[0] {
                left.push(interval);
            } else if interval[0] > new_interval[1] {
                right.push(interval);
            } else {
                merged[0] = cmp::min(interval[0], merged[0]);
                merged[1] = cmp::max(interval[1], merged[1]);
            }
        }

        left.push(merged);
        left.append(&mut right);
        left
    }
}

struct TestCase {
    input: (Vec<Vec<i32>>, Vec<i32>),
    expected: Vec<Vec<i32>>,
}

fn main() {
    let test_cases = vec![
        TestCase {
            input: (vec![vec![1, 3], vec![6, 9]], vec![2, 5]),
            expected: vec![vec![1, 5], vec![6, 9]],
        },
        TestCase {
            input: (
                vec![
                    vec![1, 2],
                    vec![3, 5],
                    vec![6, 7],
                    vec![8, 10],
                    vec![12, 16],
                ],
                vec![4, 8],
            ),
            expected: vec![vec![1, 2], vec![3, 10], vec![12, 16]],
        },
    ];

    for tc in test_cases {
        let (intervals, new_intervals) = tc.input.clone();
        let actual = Solution::insert(intervals, new_intervals);
        assert_eq!(actual, tc.expected, "input: {:?}", tc.input);
    }
}
