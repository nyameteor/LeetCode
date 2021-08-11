from typing import List


class Solution:
    def removeDuplicates(self, nums: List[int]) -> int:
        if len(nums) == 0:
            return 0
        i = 0
        for j in range(0, len(nums)):
            if j >= 1:
                if nums[j] == nums[j-1]:
                    pass
                elif nums[j] > nums[j-1]:
                    i = i + 1
                    nums[i] = nums[j]
            else:
                pass
        return i + 1
