#include <cassert>
#include <iostream>
#include <map>
#include <string>
#include <vector>

using namespace std;

class Solution {
  public:
    bool isAnagram(string s, string t) {
        // initialize all elements to 0
        int record[26] = {0};

        for (int i = 0; i < s.size(); i++) {
            record[s[i] - 'a']++;
        }
        for (int i = 0; i < t.size(); i++) {
            record[t[i] - 'a']--;
        }
        for (int i = 0; i < 26; i++) {
            // s or t has more character than another
            if (record[i] != 0)
                return false;
        }
        return true;
    }
};

int main() {
    Solution solution = Solution();
    assert(solution.isAnagram("anagram", "nagaram") == true);
    assert(solution.isAnagram("rat", "car") == false);
    assert(solution.isAnagram("ab", "abb") == false);
}