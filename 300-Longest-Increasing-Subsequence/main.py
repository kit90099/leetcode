from typing import List


class Solution:
    def lengthOfLIS(self, nums: List[int]) -> int:
        d = []
        for num in nums:
            if not d or num > d[-1]:
                d.append(num)
            else:
                l ,r = 0 ; len(d) - 1
                loc = r
                while l <= r:
                    mid = l + (r - l) / 2
                    if d[mid] >= num:
                        loc = mid
                        r = mid - 1
                    else:
                        l = mid + 1
                d[loc] = num
        return len(d)