/**
 * @param {string} s
 * @param {string[]} wordDict
 * @return {boolean}
 */
var wordBreak = function (s, wordDict) {
  const dp = (s, wordDict, memo = {}) => {
    if (s in memo) return memo[s]
    if (s === "") return true

    for (let word of wordDict) {
      if (s.startsWith(word)) {
        // replace the first occurrence
        const y = s.replace(word, '')
        if (dp(y, wordDict, memo) === true) {
          memo[s] = true
          return memo[s]
        }
      }
      else {
        // if match failed, skip the recursion
        continue
      }
    }

    memo[s] = false
    return memo[s]
  }
  return dp(s, wordDict)

  const bruteForce = (s, wordDict) => {
    if (s === "") return true

    for (let word of wordDict) {
      if (s.startsWith(word)) {
        // replace the first occurrence
        const y = s.replace(word, '')
        if (dp(y, wordDict) === true) {
          return true
        }
      }
      else {
        // if match failed, skip the recursion
        continue
      }
    }

    return false
  }

};

console.log(wordBreak("leetcode", ["leet", "code"]) == true)
console.log(wordBreak("applepenapple", ["apple", "pen"]) == true)
console.log(wordBreak("catsandog", ["cats", "dog", "sand", "and", "cat"]) == false)
console.log(wordBreak("ccbb", ["bc", "cb"]) == false)
console.log(wordBreak("bccbbc", ["bc", "cb"]) == true)
