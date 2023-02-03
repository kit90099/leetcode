from collections import deque
from typing import List


class Solution:
    def maxSlidingWindow(self, nums: List[int], k: int) -> List[int]:
        q = deque()
        n = len(nums)
        res = []
        pos = 0

        def push(item):
            while len(q) > 0 and q[-1][0] <= item[0]:
                q.pop()
            q.append(item)
            if q[0][1] < pos - k:
                q.popleft()

        for i in range(0, k):
            pos += 1
            push([nums[i], i])
        res.append(q[0])

        for i in range(k, n):
            pos += 1
            push([nums[i], i])
            res.append(q[0])

        return res

s = Solution()
print(s.maxSlidingWindow([1,3,-1,-3,5,3,6,7], 3))
