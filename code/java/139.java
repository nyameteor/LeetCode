import java.util.HashMap;
import java.util.List;
import java.util.Map;

class Solution {
    public boolean wordBreak(String s, List<String> wordDict) {
        Map<String, Boolean> memo = new HashMap<>();
        return dp(s, wordDict, memo);
    }

    private boolean dp(String s, List<String> wordDict, Map<String, Boolean> memo) {
        if (memo.containsKey(s))
            return memo.get(s);
        if (s.equals(""))
            return true;

        for (String word : wordDict) {
            String y;
            if (s.startsWith(word)) {
                y = s.substring(word.length());
            } else {
                continue;
            }
            if (dp(y, wordDict, memo) == true) {
                memo.put(s, true);
                return memo.get(s);
            }
        }

        memo.put(s, false);
        return memo.get(s);
    }

    public boolean bruteForce(String s, List<String> wordDict) {
        if (s.equals(""))
            return true;

        for (String word : wordDict) {
            String y;
            if (s.startsWith(word)) {
                y = s.substring(word.length());
            } else {
                continue;
            }
            if (bruteForce(y, wordDict) == true) {
                return true;
            }
        }

        return false;
    }

    public static void main(String[] args) {
        Solution solution = new Solution();
        System.out.println(solution.wordBreak("leetcode", List.of("leet", "code")) == true);
        System.out.println(solution.wordBreak("applepenapple", List.of("apple", "pen")) == true);
        System.out.println(solution.wordBreak("catsandog", List.of("cats", "dog", "sand", "and", "cat")) == false);
        System.out.println(solution.wordBreak("ccbb", List.of("bc", "cb")) == false);
        System.out.println(solution.wordBreak("cars", List.of("car", "ca", "rs")) == true);
    }
}