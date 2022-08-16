class Solution:
    def trailingZeroes(self, n: int) -> int:
        divisor = 5
        res: int = 0
        while divisor < n:
            res += n / divisor
            divisor *= 5
        return res