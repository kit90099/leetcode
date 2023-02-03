""" class Solution:
    def decodeString(self, s: str) -> str:
        stk = []
        for c in s:
            if c == "]":
                sub = []
                while not stk[-1] == "[":
                    sub = [stk.pop()] + sub
                stk.pop()
                times = ""
                while stk and stk[-1].isdigit():
                    times = stk.pop() + times
                for _ in range(0, int(times)):
                    stk += sub
            else:
                stk.append(c)

        return "".join(stk) """

from collections import deque


class Solution:
    def decodeString(self, s: str) -> str:
        return helper(deque(s))


def helper(arr: deque) -> str:
    time = 0
    res = ""
    while arr:
        c = arr.popleft()
        if c.isdigit():
            time = time * 10 + int(c)
        elif c == "[":
            sub = helper(arr)
            for _ in range(0, time):
                res += sub
            time = 0
        elif c == "]":
            break
        else:
            res += c
    return res


s = Solution()
print(s.decodeString("3[a]2[bc]"))
