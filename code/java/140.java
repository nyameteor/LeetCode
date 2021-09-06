import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.stream.Collectors;

class Solution {
    private StringBuilder builder = new StringBuilder();

    public List<String> wordBreak(String s, List<String> wordDict) {
        HashMap<String, List<String>> memo = new HashMap<>();
        return dp(s, wordDict, memo);
    }

    private List<String> dp(String s, List<String> wordDict, HashMap<String, List<String>> memo) {
        if (memo.containsKey(s))
            return memo.get(s);
        if (s.equals(""))
            return List.of("");

        List<String> ans = new ArrayList<>();

        for (String word : wordDict) {
            if (s.startsWith(word)) {
                String reminder = s.substring(word.length());
                List<String> reminderResults = dp(reminder, wordDict, memo);

                if (!reminderResults.equals(List.of())) {
                    List<String> sResults = reminderResults.stream().map(e -> resultAdd(e, word))
                            .collect(Collectors.toList());
                    ans.addAll(sResults);
                }
            } else {
                continue;
            }
        }

        memo.put(s, ans);
        return memo.get(s);
    }

    private String resultAdd(String e, String word) {
        // empty string builder
        builder.delete(0, builder.length());

        // edge case
        if (e.equals("")) {
            // Use StringBuilder inside a loop body when concatenating multiple strings
            return builder.append(word).append(e).toString();
        } else {
            return builder.append(word).append(" ").append(e).toString();
        }
    }

    public static void main(String[] args) {
        Solution solution = new Solution();

        // [cat sand dog, cats and dog]
        System.out.println(solution.wordBreak("catsanddog", List.of("cat", "cats", "and", "sand", "dog")));

        // [pine apple pen apple, pine applepen apple, pineapple pen apple]
        System.out.println(
                solution.wordBreak("pineapplepenapple", List.of("apple", "pen", "applepen", "pine", "pineapple")));

        // []
        System.out.println(solution.wordBreak("catsandog", List.of("cat", "cats", "and", "sand", "dog")));
    }
}