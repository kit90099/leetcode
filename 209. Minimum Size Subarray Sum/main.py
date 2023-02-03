import math
from typing import List


""" class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        # step 1 r move right
        # step 2 l move right
        n = len(nums)
        l, r = 0, 0
        sum = nums[0]
        res = 999999999
        while r < n:
            if sum >= target:
                ## s3
                res = min(res, r - l + 1)
                sum -= nums[l]
                l += 1
            else:
                r += 1
                sum += nums[r]
        return res """


class Solution:
    def minSubArrayLen(self, target: int, nums: List[int]) -> int:
        n = len(nums)
        l, r = 0, 0
        sum = 0
        res = math.inf
        while r < n:
            sum += nums[r]
            r += 1

            while sum >= target:
                res = min(res, r-l)
                sum -= nums[l]
                l += 1
        return (0 if res == math.inf else res)


s = Solution()
print(s.minSubArrayLen(7, [2, 3, 1, 2, 4, 3]))
