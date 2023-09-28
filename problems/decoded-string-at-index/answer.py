class Solution:
    def decodeAtIndex(self, s: str, k: int) -> str:
        tape_len = 0

        for c in s:
            if c.isalpha():
                tape_len += 1
            elif c.isdigit():
                tape_len *= int(c)

        for i in range(len(s)-1, -1, -1):
            c = s[i]
            if c.isalpha():
                if k == 0 or tape_len == k:
                    return c
                tape_len -= 1
            elif c.isdigit():
                tape_len //= int(c)
                k %= tape_len

        return "Error"


if __name__ == '__main__':
    solution = Solution()

    # "o"
    print(solution.decodeAtIndex("leet2code3", 10))

    # "e"
    print(solution.decodeAtIndex("leet2code3", 26))

    # "h"
    print(solution.decodeAtIndex("ha22", 5))

    # "a"
    print(solution.decodeAtIndex("a2345678999999999999999", 1))

    # "a"
    print(solution.decodeAtIndex("abc", 1))

    # "b"
    print(solution.decodeAtIndex("a2b3c4d5e6f7g8h9", 9))

    # "z"
    print(solution.decodeAtIndex("vzpp636m8y", 2920))

    # "y"
    print(solution.decodeAtIndex("y959q969u3hb22odq595", 222280369))
