#include <iostream>
#include <string>
#include <unordered_map>
#include <vector>
using namespace std;

// Sliding Window, Vector(with scheme)
class Solution {
  public:
    vector<int> findAnagrams(string s, string p) {
        int p_size = size(p);
        int s_size = size(s);

        // edge case
        if (p_size > s_size) {
            return vector<int>{};
        }

        vector<int> pFreq(26, 0); // frequency of character in window
        vector<int> wFreq(26, 0); // frequency of character in pattern

        // init frequency
        for (int i = 0; i < p_size; i++) {
            pFreq[p[i] - 'a']++;
            wFreq[s[i] - 'a']++;
        }

        vector<int> res;

        // check frequency in each window
        for (int i = 0; i <= s_size - p_size; i++) {
            if (pFreq == wFreq) {
                res.push_back(i);
            }
            if (i < s_size - p_size) {
                wFreq[s[i] - 'a']--;
                wFreq[s[i + p_size] - 'a']++;
            }
        }

        return res;
    }
};

// Sliding Window, Hash Table
class Solution2 {
  public:
    vector<int> findAnagrams(string s, string p) {
        int p_size = size(p);
        int s_size = size(s);

        if (p_size > s_size) {
            return vector<int>{};
        }

        unordered_map<char, int> wFreq;
        unordered_map<char, int> pFreq;

        for (int i = 0; i < p_size; i++) {
            pFreq[p[i]]++;
            wFreq[s[i]]++;
        }

        vector<int> res;

        for (int i = 0; i <= s_size - p_size; i++) {
            if (isMapDataEqual(pFreq, wFreq)) {
                res.push_back(i);
            }
            if (i < s_size - p_size) {
                wFreq[s[i]]--;
                wFreq[s[i + p_size]]++;
            }
        }

        return res;
    }

    bool isMapDataEqual(unordered_map<char, int> &m1,
                        unordered_map<char, int> &m2) {
        bool equal = true;
        for (auto p : m1) {
            if (m1[p.first] != m2[p.first]) {
                equal = false;
                break;
            }
        }
        return equal;
    }
};

void printResult(vector<int> v) {
    for (auto x : v) {
        cout << x << ' ';
    }
    cout << endl;
}

int main() {
    Solution solution;
    string s, p;

    // 0, 6
    s = "cbaebabacd";
    p = "abc";
    printResult(solution.findAnagrams(s, p));

    // 0, 1, 2
    s = "abab";
    p = "ab";
    printResult(solution.findAnagrams(s, p));

    s = "abacbabc";
    p = "abc";
    printResult(solution.findAnagrams(s, p));

    return 0;
}