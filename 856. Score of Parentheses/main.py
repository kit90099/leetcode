class Solution:
    def scoreOfParentheses(self, s: str) -> int:
        stk = [0]
        for c in s:
            if c == "(":
                stk.append(0)
            else:
                v = stk.pop()
                stk[-1] += max(2 * v, 1)
        return stk[-1]


s = Solution()
print(s.scoreOfParentheses("(()())"))
