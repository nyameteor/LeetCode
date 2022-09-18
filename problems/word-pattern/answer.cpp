#include <iostream>
#include <string>
#include <unordered_map>
using namespace std;

class Solution {
public:
    bool wordPattern(string pattern, string s) {
        unordered_map<char, string> mapA; // pattern -> word
        unordered_map<string, char> mapB; // word -> pattern
        int i = 0, j = 0;
        for (; i < size(s); i++) {
            string word;
            while (i < size(s) && !isspace(s[i])) {
                word.push_back(s[i]);
                i++;
            }
            char x = pattern[j];
            j++;

            // check bijection between pattern and word
            // check pattern -> word
            if (mapA.find(x) != end(mapA)) {
                if (mapA[x] != word) {
                    return false;
                }
            } else {
                mapA[x] = word;
            }

            // check word -> pattern
            if (mapB.find(word) != end(mapB)) {
                if (mapB[word] != x) {
                    return false;
                }
            } else {
                mapB[word] = x;
            }
        }
        if (j < size(pattern)) { // not all patterns are matched
            return false;
        } else {
            return true;
        }
    }
};

int main() {
    Solution solution;
    string pattern;
    string s;

    pattern = "abba";
    s = "dog cat cat dog";
    cout << solution.wordPattern(pattern, s) << endl;

    pattern = "abba";
    s = "dog cat cat fish";
    cout << solution.wordPattern(pattern, s) << endl;

    pattern = "abba";
    s = "dog dog dog dog";
    cout << solution.wordPattern(pattern, s) << endl;

    pattern = "jquery";
    s = "jquery";
    cout << solution.wordPattern(pattern, s) << endl;

    return 0;
}