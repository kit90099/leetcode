# The isBadVersion API is already defined for you.
# def isBadVersion(version: int) -> bool:

import math


class Solution:
    def firstBadVersion(self, n: int) -> int:
        l, r = 1, n
        while l <= r:
            mid = math.floor((l + (r - l) / 2))
            if isBadVersion(mid):
                r = mid - 1
            else:
                l = mid + 1