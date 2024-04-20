class Solution:
    def longestPalindrome(self, s: str) -> str:
        self.res = ""
        N = len(s)

        def helper(l: int, r: int):
            while l >= 0 and r < N and s[l] == s[r]:
                if r + 1 - l > len(self.res):
                    self.res = s[l : r + 1]
                l -= 1
                r += 1

        for i in range(N):
            helper(i, i)
            helper(i, i + 1)

        return self.res


if __name__ == "__main__":
    solution = Solution()

    s = "babad"
    res = solution.longestPalindrome(s)
    # bab
    print(res)

    s = "cbbd"
    # bb
    res = solution.longestPalindrome(s)
    print(res)

    s = "ccc"
    # ccc
    res = solution.longestPalindrome(s)
    print(res)

    s = "a"
    # a
    res = solution.longestPalindrome(s)
    print(res)

    s = "abbcccbbbcaaccbababcbcabca"
    # bbcccbb
    res = solution.longestPalindrome(s)
    print(res)
