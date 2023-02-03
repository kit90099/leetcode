from typing import List


class Solution:
    def fourSumCount(self, nums1: List[int], nums2: List[int], nums3: List[int], nums4: List[int]) -> int:
        map12 = {}
        for a in nums1:
            for b in nums2:
                map12.setdefault(a+b, 0)
                map12[a+b] += 1

        res = 0
        for c in nums3:
            for d in nums4:
                if -(c+d) in map12:
                    res += map[-(c+d)]
        return res