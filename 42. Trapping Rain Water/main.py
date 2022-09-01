from typing import List


class Solution:
    def trap(self, height: List[int]) -> int:
        left, right = 0, len(height)-1
        lmax,rmax = 0,0
        res = 0

        while left < right:
            lmax = max(lmax, height[left])
            rmax = max(rmax, height[right])

            if (lmax < rmax):
                res += (lmax-height[left])
                left+=1
            else:
                res += (rmax-height[right])
                right-=1

        return res

s = Solution()
s.trap([0,1,0,2,1,0,1,3,2,1,2,1])