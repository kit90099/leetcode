import math
from typing import List


class Solution:
    def countSmaller(self, nums: List[int]) -> List[int]:
        res = [0 for i in range(0, len(nums))]
        dict = []
        for idx, item in enumerate(nums):
            dict.append([idx, item])

        def mergeSort(nums: List[List[int]], start: int, end: int):
            if start >= end:
                return

            mid = math.floor((start + end) / 2)
            mergeSort(nums, start, mid)
            mergeSort(nums, mid + 1, end)

            lptr = start
            rptr = mid + 1
            temp = nums[:]
            for p in range(start, end + 1):
                if lptr == mid + 1:
                    nums[p] = temp[rptr]
                    rptr += 1
                elif rptr == end + 1:
                    nums[p] = temp[lptr]
                    lptr += 1
                    res[nums[p][0]] += (rptr - mid - 1)
                elif temp[lptr][1] <= temp[rptr][1]:
                    nums[p] = temp[lptr]
                    lptr += 1
                    res[nums[p][0]] += (rptr - mid - 1)
                else:
                    nums[p] = temp[rptr]
                    rptr += 1

        mergeSort(dict, 0, len(nums) - 1)

s = Solution()
s.countSmaller([26,78,27,100,33,67,90,23,66,5,38,7,35,23,52,22,83,51,98,69,81,32,78,28,94,13,2,97,3,76,99,51,9,21,84,66,65,36,100,41])
