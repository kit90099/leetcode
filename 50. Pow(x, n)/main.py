""" class Solution:
    def myPow(self, x: float, n: int) -> float:
        if n == 0:
            return 1
        if n == 1:
            return x
        if n == -1:
            return 1/x
        if n % 2 == 0:
            half = self.myPow(x, n/2)
            return half * half
        else:
            half = self.myPow(x, (n - 1)/2)
            return half * half * x """


""" class Solution:
    def myPow(self, x: float, n: int) -> float:
        def quickMul(n):
            ans = 1
            con = x
            while n > 0:
                if n % 2 == 1:
                    ans *= con
                    
                con *= con
                n //= 2
            return ans
        return quickMul(n) if n >= 0 else 1/quickMul(-n)

s = Solution()
s.myPow(2,2) """


class Solution:
    def myPow(self, x: float, n: int) -> float:
        def quickMul(a, b):
            ans = 1
            while b > 0:
                if b & 1:
                    ans *= a
                a *= a
                b >>= 1
        return quickMul(x, n) if n >=0 else 1/quickMul(x, -n)