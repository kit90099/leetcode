from typing import List


class Solution:
    def sumSubarrayMins(self, arr: List[int]) -> int:
        if not arr:
            return 0
        n = len(arr)
        left = [-1] * n
        right = [-1] * n
        stk = []

        for i, e in enumerate(arr):
            while stk and stk[-1] > e:
                stk.pop()
            if not stk:
                left[i] = -1
            else:
                left[i] = stk[-1]
            stk.append(e)
        stk = []
        for i in range(n - 1, -1, -1):
            e = arr[i]
            while stk and stk[-1] > e:
                stk.pop()
            if not stk:
                right[i] = -1
            else:
                right[i] = stk[-1]
            stk.append(e)

        ans = 0
        for i, e in enumerate(arr):
            ans = (ans + (i - left[i]) * (right[i] - i) * arr[i]) % 1000000007
        return ans

s = Solution()
s.sumSubarrayMins([3,1,2,4])