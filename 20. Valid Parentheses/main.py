import collections


class Solution:
    def isValid(self, s: str) -> bool:
        l = collections.deque()

        for x in s:
            if x == "(" or x == "[" or x == "{":
                l.append(x)
            else:
                if len(l) == 0:
                    return False
                if x == ")":
                    if not l.pop() == "(":
                        return False
                elif x == "]":
                    if not l.pop() == "[":
                        return False
                        return False
                elif x == "}":
                    if not l.pop() == "{":
                        return False
                
        return len(l) == 0

s = Solution()
print(s.isValid("([)"))