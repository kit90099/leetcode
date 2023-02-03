class Solution:
    def divide(self, dividend: int, divisor: int) -> int:
        if divisor == -1 and dividend == -2147483648:
            return 2147483647

        isPositive = (dividend ^ divisor) >= 0
        dividend = abs(dividend)
        divisor = abs(divisor)

        ans = 0
        for i in range(31, -1, -1):
            if (dividend >> i) >= divisor: ## i is times that divisor multiply with
                dividend -= divisor << i ## if dividend is able to minus divisor << i then ans += 1 << i
                ans += 1 << i
        
        return ans if isPositive else -ans

s = Solution()
s.divide(10, 3)