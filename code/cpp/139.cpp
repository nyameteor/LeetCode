#include <iostream>
#include <map>
#include <string>
#include <vector>

using namespace std;

class Solution {
public:
  bool wordBreak(string s, vector<string> &wordDict) {
    map<string, bool> memo = {};
    return dp(s, wordDict, memo);
  }

  bool dp(string s, vector<string> &wordDict, map<string, bool> &memo) {
    // find in memo
    if (memo.find(s) != memo.end())
      return memo.at(s);

    // s equal to ""
    if (!s.compare(""))
      return true;

    for (int i = 0; i < wordDict.size(); i++) {
      // if s start with wordDict[i]
      if (s.rfind(wordDict[i], 0) == 0) {
        string reminder = s;
        reminder.replace(0, wordDict[i].length(), "");
        bool reminderResult = dp(reminder, wordDict, memo);
        if (reminderResult == true) {
          memo.insert(make_pair(s, true));
          return memo.at(s);
        } else {
          continue;
        }
      }
    }

    memo.insert(make_pair(s, false));
    return memo.at(s);
  }
};

int main() {
  Solution solution;
  vector<string> wordDict = {"leet", "code"};
  cout << solution.wordBreak("leetcode", wordDict);
  wordDict = {"apple", "pen"};
  cout << solution.wordBreak("applepenapple", wordDict);
  wordDict = {"cats", "dog", "sand", "and", "cat"};
  cout << solution.wordBreak("acatsandog", wordDict);
  wordDict = {"cats", "dog", "sand", "and", "cat"};
  cout << solution.wordBreak("acatsandandcatdogcatssanddog", wordDict);
}