class Solution:
    def minRemoveToMakeValid(self, s: str) -> str:
        count = 0
        stk = []
        for c in s:
            if c == "(":
                count+=1
                stk.append(c)
            elif c == ")":
                if count == 0:
                    continue
                else:
                    count-=1
                    stk.append(c)
            else:
                stk.append(c)
        
        if count > 0:
            i = 0
            while count > 0:
                if stk[i] == "(":
                    stk[i] = ""
                    count -= 1
                i += 1
        
        return "".join(stk)

s = Solution()
s.minRemoveToMakeValid("())()(((")