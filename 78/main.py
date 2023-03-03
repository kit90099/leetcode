from typing import List


class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        n = len(nums)
        c = (1 << n) - 1
        ans = []

        while c != 0:
            t = c
            while i in range(n):
                if (t >> i) % 2 == 1:


s = Solution()
s.subsets([1,2,3])