import collections
from typing import List


class Solution:
    def calculate(self, s: str) -> int:

        def calculator(s: collections.deque) -> int:
            stk = []
            sign = '+'
            num = 0
            while len(s) > 0:
                c = s.popleft()

                if c.isdigit():
                    num = num*10 + int(c)

                if c == '(':
                    num = calculator(s)

                if (not c.isdigit() and not c == ' ') or len(s) == 0:
                    if sign == '+':
                        stk.append(num)
                    elif sign == '-':
                        stk.append(-num)
                    elif sign == '*':
                        stk[-1] = stk[-1]*num
                    elif sign == '/':
                        stk[-1] = stk[-1]/num

                    sign = c
                    num = 0

                if c == ')':
                    break

            return sum(stk)

        return calculator(collections.deque(s))


s = Solution()
print(s.calculate("1+(2+3)"))
