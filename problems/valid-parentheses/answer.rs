struct Solution;

impl Solution {
    pub fn is_valid(s: String) -> bool {
        let mut stack: Vec<char> = Vec::new();
        for c in s.chars() {
            match c {
                '(' | '[' | '{' => stack.push(c),
                ')' | ']' | '}' => {
                    let expected = match c {
                        ')' => '(',
                        ']' => '[',
                        '}' => '{',
                        _ => unreachable!(),
                    };
                    if stack.pop() != Some(expected) {
                        return false;
                    }
                }
                _ => {
                    return false;
                }
            }
        }
        stack.is_empty()
    }
}

struct TestCase {
    input: String,
    expected: bool,
}

fn main() {
    let test_cases = vec![
        TestCase {
            input: "()".to_string(),
            expected: true,
        },
        TestCase {
            input: "()[]{}".to_string(),
            expected: true,
        },
        TestCase {
            input: "(]".to_string(),
            expected: false,
        },
        TestCase {
            input: "([])".to_string(),
            expected: true,
        },
        TestCase {
            input: "([)]".to_string(),
            expected: false,
        },
    ];

    for tc in test_cases {
        let actual = Solution::is_valid(tc.input.clone());
        assert_eq!(actual, tc.expected, "input: {}", tc.input);
    }
}
