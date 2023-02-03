from typing import List


class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        def findLeft():
            l, r = 0 , len(nums) - 1
            while l <= r:
                mid = l + (r - l) // 2
                if nums[mid] >= target:
                    r = mid - 1
                else:
                    l = mid+1
            if l == len(nums):
                return -1
            if nums[l] == target:
                return l
            else:
                return -1

        def findRight():
            l, r = 0 , len(nums) - 1
            while l <= r:
                mid = l + (r - l) // 2
                if nums[mid] <= target:
                    l = mid + 1
                else:
                    r = mid - 1
            if l - 1 < 0:
                return -1
            if nums[l - 1] == target:
                return l - 1
            else:
                return -1

        return [findLeft(), findRight()]

s = Solution()
s.searchRange([5,7,7,8,8,10],8)