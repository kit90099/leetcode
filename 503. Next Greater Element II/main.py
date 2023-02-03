from typing import List


class Solution:
    def nextGreaterElements(self, nums: List[int]) -> List[int]:
        n = len(nums)
        res = [0 * n]
        stk = []
        for i in range(2 * n - 1, -1, -1):
            while s and s[-1] <= nums[i % n]:
                # s[i+1] is smaller than s[i], so s[i] will never be used
                s.pop()
            if stk:
                res[i % n] = s[1]
            else:
                res[i % n] = -1
            stk.push(nums[i % n])
        return res


s = Solution()
s.nextGreaterElements([1, 2])
