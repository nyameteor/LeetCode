from typing import List


class Solution:
    def longestOnes(self, nums: List[int], k: int) -> int:
        zeroIdxes = []
        curCnt = 0
        maxCnt = 0
        for i in range(len(nums)):
            if nums[i] == 1:
                curCnt += 1
                maxCnt = max(maxCnt, curCnt)
            elif len(zeroIdxes) < k:
                curCnt += 1
                maxCnt = max(maxCnt, curCnt)
                zeroIdxes.append(i)
            elif k > 0:
                curCnt = i - zeroIdxes.pop(0)
                zeroIdxes.append(i)
            elif k == 0:
                curCnt = 0
        return maxCnt


if __name__ == '__main__':
    solution = Solution()

    nums = [1, 1, 1, 0, 0, 0, 1, 1, 1, 1, 0]
    k = 2
    # 6
    print(solution.longestOnes(nums, k))

    nums = [0, 0, 1, 1, 0, 0, 1, 1, 1, 0, 1, 1, 0, 0, 0, 1, 1, 1, 1]
    k = 3
    # 10
    print(solution.longestOnes(nums, k))

    nums = [0, 0, 1, 1, 1, 0, 0]
    k = 0
    # 3
    print(solution.longestOnes(nums, k))

    nums = [0, 0, 0, 1]
    k = 4
    # 4
    print(solution.longestOnes(nums, k))
