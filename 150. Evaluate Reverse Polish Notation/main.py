import math
from typing import List


class Solution:
    def evalRPN(self, tokens: List[str]) -> int:
        stk = []
        for t in tokens:
            if t.lstrip("-").isdigit():
                stk.append(int(t))
            elif t == "+":
                stk.append(stk.pop() + stk.pop())
            elif t == "-":
                b = stk.pop()
                a = stk.pop()
                stk.append(a - b)
            elif t == "*":
                stk.append(stk.pop() * stk.pop())
            elif t == "/":
                b = stk.pop()
                a = stk.pop()
                stk.append(int(a / b))

        return sum(stk)

s = Solution()
print(s.evalRPN(["10","6","9","3","+","-11","*","/","*","17","+","5","+"]))