import math


class Solution:
    def mySqrt(self, x: int) -> int:
        if x < 2:
            return x
        left, right = 1, x
        while left < right:
            mid = math.floor((left + right)/2)
            square = mid * mid
            if square == x:
                return mid
            elif square > x:
                right = mid
            else:
                left = mid+1
        return left - 1

s = Solution()
s.mySqrt(8)