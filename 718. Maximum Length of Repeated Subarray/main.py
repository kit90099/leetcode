import math
from typing import List


class Solution:
    def findLength(self, nums1: List[int], nums2: List[int]) -> int:
        def findMaxLength(a, b, length):
            res = k = 0
            for i in range(length):
                if nums1[a+i] == nums2[b+i]:
                    k += 1
                    res = max(res, k)
                else:
                    k = 0
            return res
        m, n = len(nums1), len(nums2)
        res = 0
        for i in range(m):
            length = min(n, m-i)
            res = max(res, findMaxLength(i, 0, length))
        for i in range(n):
            length = min(m, n-i)
            res = max(res, findMaxLength(0, i, length))
        return res


s = Solution()
s.findLength([1, 2, 3, 2, 1], [3, 2, 1, 4, 7, 8])
