from typing import List


class Solution(object):
    def wordBreak(self, s: str, wordDict: List[str]) -> bool:
        memo = {}
        return self.dp(s, wordDict, memo)

    def dp(self, s: str, wordDict: List[str], memo) -> bool:
        if s in memo:
            return memo[s]
        if s == "":
            return True

        for word in wordDict:
            if(s.startswith(word)):
                y = s[len(word):]
                if self.dp(y, wordDict, memo) == True:
                    memo[s] = True
                    return memo[s]
            else:
                # if match failed, skip the recursion
                continue

        memo[s] = False
        return memo[s]

    def wordBreakBruteForce(self, s: str, wordDict: List[str]) -> bool:
        if s == "":
            return True

        for word in wordDict:
            if s.startswith(word):
                y = s[len(word):]
                if self.wordBreakBruteForce(y, wordDict) == True:
                    return True
            else:
                # if match failed, skip the recursion
                continue

        return False


if __name__ == "__main__":
    s = Solution()
    print(s.wordBreak("a", ["b"]) == False)
    print(s.wordBreak("ccbb", ["bc", "cb"]) == False)
    print(s.wordBreak("leetcode", ["leet", "code"]) == True)
    print(s.wordBreak("applepenapple", ["apple", "pen"]) == True)
    print(s.wordBreak("catsandog", [
          "cats", "dog", "sand", "and", "cat"]) == False)
