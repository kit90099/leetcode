from typing import List


class Solution:
    def maxArea(self, height: List[int]) -> int:
        left, right = 0, len(height)-1
        res = 0

        while left < right:
            if height[left] < height[right]:
                res = max(height[left] * (right-left), res)
                left += 1
            else:
                res = max(height[right] * (right-left),res)
                right -= 1

        return res