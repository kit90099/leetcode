import random
from typing import List


class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        if len(nums) == 1:
            return [nums[0]]
        map = {}
        for num in nums:
            if num in map:
                map[num] += 1
            else:
                map[num] = 1

        pair = []
        for key in map:
            pair.append([key, map[key]])

        def sort(lo, hi):
            if lo > hi:
                return pair[:lo]

            p = partition(lo, hi)
            if p == k:
                return pair[:k]
            elif p > k:
                return sort(lo, p-1)
            else:
                return sort(p+1, hi)

        def partition(lo, hi):
            pivot = pair[lo][1]
            i = lo+1
            j = hi
            while i <= j:
                while i <= hi and pair[i][1] > pivot:
                    i += 1
                while j > lo and pair[j][1] <= pivot:
                    j -= 1

                if i >= j:
                    break

                swap(i, j)
            swap(j, lo)
            return j

        def swap(a, b):
            pair[a], pair[b] = pair[b], pair[a]
        
        res = sort(0, len(pair)-1)
        return [i[0] for i in res]


s = Solution()
s.topKFrequent([1, 1, 1, 2, 2, 3], 2)
