#include <cassert>
#include <iostream>
#include <map>
#include <string>
#include <vector>

using namespace std;

class Solution {
public:
  vector<string> wordBreak(string s, vector<string> &wordDict) {
    map<string, vector<string>> memo = {};
    return dp(s, wordDict, memo);
  }

  vector<string> dp(string s, vector<string> &wordDict,
                    map<string, vector<string>> &memo) {
    vector<string> ans;

    // find in memo
    if (memo.find(s) != memo.end())
      return memo.at(s);

    // s equal to ""
    if (!s.compare("")) {
      ans = {""};
      return ans;
    }

    for (int i = 0; i < wordDict.size(); i++) {
      // if s start with wordDict[i]
      if (s.rfind(wordDict[i], 0) == 0) {
        string reminder = s.substr(wordDict[i].size());
        vector<string> reminderResult = dp(reminder, wordDict, memo);

        vector<string> empty;
        if (reminderResult != empty) {
          vector<string> sResult;

          for (int j = 0; j < reminderResult.size(); j++) {
            string insertStr = reminderResult[j];
            // edge case
            if (reminderResult[j] != "") {
              insertStr.insert(0, " ");
            }
            insertStr.insert(0, wordDict[i]);
            sResult.push_back(insertStr);
          }

          // appending vector "sResult" to vector "ans"
          ans.insert(ans.end(), sResult.begin(), sResult.end());
        }
      }
    }

    memo.insert(make_pair(s, ans));
    return memo.at(s);
  }
};

int main() {
  Solution solution;
  vector<string> wordDict;
  vector<string> ans;
  vector<string> res;

  wordDict = {"cat", "cats", "and", "sand", "dog"};
  ans = {"cat sand dog", "cats and dog"};
  res = solution.wordBreak("catsanddog", wordDict);
  assert(res == ans);

  wordDict = {"apple", "pen", "applepen", "pine", "pineapple"};
  ans = {"pine apple pen apple", "pine applepen apple", "pineapple pen apple"};
  res = solution.wordBreak("pineapplepenapple", wordDict);
  assert(res == ans);

  wordDict = {"cats", "dog", "sand", "and", "cat"};
  ans = {};
  res = solution.wordBreak("catsandog", wordDict);
  assert(res == ans);
}