from typing import List


class Solution:
    def findPeakElement(self, nums: List[int]) -> int:
        n = len(nums)

        def get(i):
            if i == -1 or i == n:
                return float('-inf')
            return nums[i]

        l, r = 0, n-1
        while l <= r:
            mid = l + (r-l)//2
            if get(mid-1) < get(mid) > get(mid+1):
                return mid
            if get(mid) < get(mid+1):
                left = mid+1
            else:
                right = mid-1

s= Solution()
s.findPeakElement([1,2,3,1])