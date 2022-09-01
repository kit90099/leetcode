class Solution:
    def minInsertions(self, s: str) -> int:
        res = 0
        need = 0

        for x in s:
            if x == '(':
                need += 2
                if need % 2 == 1:
                    res += 1
                    need -= 1
            else:
                need -= 1
                if need == -1:
                    res += 1
                    need = 1

        return res+need


s = Solution()
print(s.minInsertions("))"))
