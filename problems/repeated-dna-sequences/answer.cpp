#include <iostream>
#include <unordered_map>
#include <vector>
using namespace std;

class Solution {
public:
    vector<string> findRepeatedDnaSequences(string s) {
        vector<string> res;

        return res;
    }
};

/**
 * Brute Force
 *
 * Idea: Use hash table to count the number of occurrences of each possible
 * sequences, if a sequence occurs more than once, then add it to result.
 *
 * Time: O(n), Space: O(n)
 */
class Solution2 {
public:
    vector<string> findRepeatedDnaSequences(string s) {
        vector<string> res;
        unordered_map<string, int> m;

        if (s.size() <= 10) {
            return res;
        }

        for (int i = 0; i <= s.size() - 10; i++) {
            string seq = s.substr(i, 10);
            if (m.find(seq) == end(m)) {
                m.insert({seq, 1});
            } else {
                if (m.at(seq) == 1) {
                    m.at(seq) = 2;
                    res.push_back(seq);
                } else {
                    continue;
                }
            }
        }

        return res;
    }
};

void printResult(vector<string> v) {
    for (auto x : v) {
        cout << x << ' ';
    }
    cout << endl;
}

int main() {
    Solution2 solution;
    string s;

    s = "AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT";
    printResult(solution.findRepeatedDnaSequences(s));

    s = "AAAAAAAAAAAAA";
    printResult(solution.findRepeatedDnaSequences(s));

    s = "A";
    printResult(solution.findRepeatedDnaSequences(s));

    return 0;
}