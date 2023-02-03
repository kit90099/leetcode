class Solution(object):
    def calPoints(self, operations):
        """
        :type operations: List[str]
        :rtype: int
        """
        stk = []
        
        for o in operations:
            if o == "+":
                stk.append(stk[-1] + stk[-2])
            elif o == "D":
                stk.append(stk[-1] * 2)
            elif o == "C":
                stk.pop()
            else:
                stk.append(int(o))
            
        
        res = 0
        for i in stk:
            res += i
        
        return res

s = Solution()
s.calPoints(["5","-2","4","C","D","9","+","+"])