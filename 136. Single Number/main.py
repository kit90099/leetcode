from typing import List
class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        res = 0
        for n in nums:
            res = n ^ res
        return res

s =Solution()
s.singleNumber([3,2,2])