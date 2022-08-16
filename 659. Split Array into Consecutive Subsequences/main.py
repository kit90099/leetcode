from typing import List


class Solution:
    def isPossible(self, nums: List[int]) -> bool:
        res = []
        used: List[bool] = [False * len(nums)]
        count = 0
        for index, item in enumerate(nums):
            if len(res) == 0 or res[-1] < item:
                res.append(item)
                used[index] = True
                if len(res) == 3:
                    count+=1
                    res = []

