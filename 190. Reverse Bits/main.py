class Solution:
    def reverseBits(self, n: int) -> int:
        ans = 0
        for _ in range(32):
            last = n & 1
            n >>= 1
            ans <<= 1
            if last == 1:
                ans += 1
        return ans
    
s = Solution()
print(s.reverseBits(43261596))