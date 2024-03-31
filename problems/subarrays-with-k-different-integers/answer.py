from typing import List


class Solution:
    def subarraysWithKDistinct(self, nums: List[int], k: int) -> int:
        return self.slidingWindowAtMost(nums, k) - self.slidingWindowAtMost(nums, k-1)

    def slidingWindowAtMost(self, nums: List[int], k: int) -> int:
        res = 0
        freqs = {}
        i, j = 0, 0
        while j < len(nums):
            freqs[nums[j]] = freqs.get(nums[j], 0) + 1
            while i <= j and len(freqs.keys()) > k:
                freqs[nums[i]] -= 1
                if freqs[nums[i]] == 0:
                    del freqs[nums[i]]
                i += 1
            res += j + 1 - i
            j += 1
        return res


if __name__ == "__main__":
    s = Solution()

    nums = [1, 2, 1, 2, 3]
    k = 2
    # Output: 7
    print(s.subarraysWithKDistinct(nums, k))

    nums = [1, 2, 1, 3, 4]
    k = 3
    # Output: 3
    print(s.subarraysWithKDistinct(nums, k))
