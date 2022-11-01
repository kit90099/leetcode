import math
from typing import List


class Solution:
    def reversePairs(self, nums: List[int]) -> int:
        res = 0

        def sort(nums: List[int], start: int, end: int):
            if start >= end:
                return

            mid = math.floor((start + end)/2)
            sort(nums, start, mid)
            sort(nums, mid+1, end)
            merge(nums, start, end, mid)

        def merge(nums: List[int], start: int, end: int, mid: int):
            temp = nums[:]

            j = mid + 1
            nonlocal res
            for i in range(start, mid+1):
                while j <= end and nums[i] > nums[j] * 2:
                    j += 1
                res += j-(mid+1)

            l, r = start, mid+1
            for p in range(start, end + 1):
                if r == end + 1:
                    nums[p] = temp[l]
                    l += 1
                elif l == mid+1:
                    nums[p] = temp[r]
                    r += 1
                elif temp[l] <= temp[r]:
                    nums[p] = temp[l]
                    l += 1
                else:
                    nums[p] = temp[r]
                    r += 1

        sort(nums, 0, len(nums)-1)
        return res

s=Solution()
s.reversePairs([1,3,2,3,1])