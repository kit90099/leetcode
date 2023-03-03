class Solution:
    def toHex(self, num: int) -> str:
        if num == 0:
            return "0"
        hexC = "0123456789abcdef"
        ans = ""
        while num != 0 and len(ans) < 8:
            ans = hexC[num%16]+ans
            num = num >> 4
        return ans

s = Solution()
s.toHex(-1)