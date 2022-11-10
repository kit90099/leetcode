class Solution:
    def simplifyPath(self, path: str) -> str:
        parts = path.split("/")
        stk = []
        for p in parts:
            if p == "..":
                if stk:
                    stk.pop()
            elif p and p != ".":
                stk.append(p)
        return "/" + "/".join(stk)


s = Solution()
print(s.simplifyPath("/../"))
