import math
from typing import List, Set, Dict


class Solution:
    def numMatchingSubseq(self, s: str, words: List[str]) -> int:
        m: Dict[str, List] = dict()

        for i, c in enumerate(s):
            m.setdefault(c, [])
            m[c].append(i)

        res = 0
        for w in words:
            last = -1
            ok = True
            for i, c in enumerate(w):
                if not c in m:
                    ok = False
                    break

                posSet = m[c]
                l, r = 0, len(posSet)
                while l < r:
                    mid = math.floor(r + (l-r)/2)
                    if posSet[mid] > last:
                        r = mid
                    else:
                        l = mid+1
                if l >= len(posSet):
                    ok = False
                    break
                last = posSet[l]
            if ok:
                res += 1

        return res


s = Solution()
print(s.numMatchingSubseq("abcde", ["a","bb","acd","ace"]))
