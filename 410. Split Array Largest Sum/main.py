from typing import List


class Solution:
    def splitArray(self, nums: List[int], k: int) -> int:
        def split(max):
            count = 1
            sum = 0
            for n in nums:
                if sum + n > max:
                    count += 1
                    sum = n
                else:
                    sum += n
            return count

        def getMax():
            m = -1
            for n in nums:
                m = max(m, n)
            return m

        def getSum():
            m = 0
            for n in nums:
                m += n
            return m

        lo, hi = getMax(), getSum()
        while lo < hi:
            mid = lo + (hi - lo) // 2
            n = split(mid)

            if n == k:
                hi = mid
            elif n <= k:
                hi = mid
            else:
                lo = mid + 1
        return lo