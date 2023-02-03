from typing import List


class Solution:
    def findMin(self, nums: List[int]) -> int:
        n = len(nums)
        if nums[0] < nums[-1]:
            return nums[0]
        if n == 1:
            return nums[0]

        def bs(l, r):
            mid = l + (r - l) // 2

            if mid == n - 1:
                return nums[-1]
            if nums[mid - 1] > nums[mid] and nums[mid] < nums[mid + 1]:
                return nums[mid]

            if nums[0] <= nums[mid]:
                ## find right
                return bs(mid + 1, r)
            else:
                return bs(l, mid - 1)

        return bs(0, n - 1)

s = Solution()
print(s.findMin([3,4,1,2,3]))