import collections
import math
from typing import List


class Solution:
    def calculate(self, s: str) -> int:
        list = collections.deque(s)
        stk = []
        sign = '+'
        num = 0
        while len(list) > 0:
            c = list.popleft()

            if c.isdigit():
                num = num*10 + int(c)

            if (not c.isdigit() and not c == ' ') or len(list) == 0:
                if sign == '+':
                    stk.append(num)
                elif sign == '-':
                    stk.append(-num)
                elif sign == '*':
                    stk[-1] = stk[-1]*num
                elif sign == '/':
                    stk[-1] = math.trunc(stk[-1]/num)

                sign = c
                num = 0

        return sum(stk)

s = Solution()
s.calculate("3+2*2")