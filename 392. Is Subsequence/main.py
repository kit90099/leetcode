from typing import Dict


class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        i, j = 0, 0

        if len(s) == 0:
            return True

        while j < len(t):
            if s[i] == t[j]:
                i += 1
                if i == len(s):
                    return True
            j += 1

        return False
