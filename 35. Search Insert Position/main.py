import math
from typing import List


class Solution:
    def searchInsert(self, nums: List[int], target: int) -> int:
        l, r = 0, len(nums) - 1
        while l <= r:
            mid = math.floor(l + (r - l) / 2)
            if nums[mid] == target:
                return mid
            elif nums[mid] < target:
                l = mid+1
            else:
                r = mid - 1
        return l

s= Solution()
s.searchInsert([1,3,5,6], 2)