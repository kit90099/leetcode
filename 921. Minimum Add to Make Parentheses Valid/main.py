class Solution:
    def minAddToMakeValid(self, s: str) -> int:
        res = 0
        left = 0

        for x in s:
            if x == '(':
                left += 1
            else:
                left -= 1
                if left == -1:
                    res+=1
                    left = 0

        return res+left

s = Solution()
print(s.minAddToMakeValid(""))