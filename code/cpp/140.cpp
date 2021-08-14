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

  wordDict = {"cat", "cats", "and", "sand", "dog"};
  ans = solution.wordBreak("catsanddog", wordDict);
  for (int i = 0; i < ans.size(); i++) {
    cout << ans[i] << ' ';
  }
  cout << endl;

  wordDict = {"apple", "pen", "applepen", "pine", "pineapple"};
  ans = solution.wordBreak("pineapplepenapple", wordDict);
  for (int i = 0; i < ans.size(); i++) {
    cout << ans[i] << ' ';
  }
  cout << endl;

  wordDict = {"cats", "dog", "sand", "and", "cat"};
  ans = solution.wordBreak("catsandog", wordDict);
  for (int i = 0; i < ans.size(); i++) {
    cout << ans[i] << ' ';
  }
  cout << endl;
}