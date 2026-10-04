use std::collections::HashMap;

struct Solution;

impl Solution {
    pub fn group_anagrams(strs: Vec<String>) -> Vec<Vec<String>> {
        let mut groups: HashMap<[u8; 26], Vec<String>> = HashMap::new();

        for s in strs {
            let mut freq: [u8; 26] = [0; 26];
            for c in s.bytes() {
                freq[(c - b'a') as usize] += 1;
            }
            groups.entry(freq).or_default().push(s);
        }

        groups.into_values().collect()
    }
}

struct TestCase {
    input: Vec<String>,
    expected: Vec<Vec<String>>,
}

fn strings(xs: Vec<&str>) -> Vec<String> {
    xs.into_iter().map(String::from).collect()
}

fn normalize(mut groups: Vec<Vec<String>>) -> Vec<Vec<String>> {
    for group in &mut groups {
        group.sort();
    }
    groups.sort();
    groups
}

fn main() {
    let test_cases = vec![
        TestCase {
            input: strings(vec!["eat", "tea", "tan", "ate", "nat", "bat"]),
            expected: vec![
                strings(vec!["bat"]),
                strings(vec!["nat", "tan"]),
                strings(vec!["ate", "eat", "tea"]),
            ],
        },
        TestCase {
            input: strings(vec![""]),
            expected: vec![strings(vec![""])],
        },
        TestCase {
            input: strings(vec!["a"]),
            expected: vec![strings(vec!["a"])],
        },
    ];

    for tc in test_cases {
        let actual = Solution::group_anagrams(tc.input.clone());
        assert_eq!(
            normalize(actual),
            normalize(tc.expected),
            "input: {:?}",
            tc.input
        );
    }
}
