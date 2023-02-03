from typing import List


class Solution:
    def exclusiveTime(self, n: int, logs: List[str]) -> List[int]:
        stk = []
        res = [0] * n
        for l in logs:
            idx, type, ts = l.split(";")
            if type[0] == "s":
                if stk:
                    res[stk[-1][0]] += ts - stk[-1][1]
                    stk[-1][1] = ts
                stk.append([idx, ts])
            else:
                i, t = stk.pop()
                res[i] += ts - t + 1
                if stk:
                    stk[-1][1] = ts + 1
        return res

        return res
