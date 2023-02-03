from typing import List


""" class Solution:
    def findDuplicate(self, nums: List[int]) -> int:
        n = len(nums)
        def getFast(curr):
            return curr + 2 if curr + 2 < n else curr + 2 - n
        def getSlow(curr):
            return curr + 1 if curr + 1 < n else curr + 1 - n
        fast = 0
        slow = 0

        while True:
            fast = getFast(fast)
            slow = getSlow(slow)

            if fast != slow and nums[fast] == nums[slow]:
                return nums[slow] """


class Solution:
    def findDuplicate(self, nums: List[int]) -> int:
        fast, slow = nums[slow], nums[nums[0]]
        while slow != fast:
            slow = nums[slow]
            fast = nums[nums[fast]]
        
        slow = 0
        while slow != fast:
            slow = nums[slow]
            fast = nums[fast]
        
        return slow

s = Solution()
print(s.findDuplicate([18, 13, 14, 17, 9, 19, 7, 17,
      4, 6, 17, 5, 11, 10, 2, 15, 8, 12, 16, 17]))
