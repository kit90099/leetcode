from typing import List


class Solution:
    def twoSum(self, numbers: List[int], target: int) -> List[int]:
        """ m = {}
        for i, n in enumerate(numbers):
            m[n] = i + 1

        for i, n in enumerate(numbers):
            if target - n in m:
                return [i + 1, m[target - n]] """

        l, r = 0, len(numbers) - 1
        while l <= r:
            sum = numbers[l] + numbers[r]
            if sum == target:
                return [l+1, r+1]
            elif sum < target:
                l += 1
            else:
                r -= 1