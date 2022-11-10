class Solution:
    def reverseParentheses(self, s: str) -> str:
        stk = []
        for c in s:
            if c == ")":
                sub = []
                while not stk[-1] == "(":
                    sub.append(stk.pop())
                stk.pop()
                for e in sub:
                    stk.append(e)
            else:
                stk.append(c)

        return "".join(stk)


s = Solution()
s.reverseParentheses("(abcd)")