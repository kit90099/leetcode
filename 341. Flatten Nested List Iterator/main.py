# """
# This is the interface that allows for creating nested lists.
# You should not implement it, or speculate about its implementation
# """
#class NestedInteger:
#    def isInteger(self) -> bool:
#        """
#        @return True if this NestedInteger holds a single integer, rather than a nested list.
#        """
#
#    def getInteger(self) -> int:
#        """
#        @return the single integer that this NestedInteger holds, if it holds a single integer
#        Return None if this NestedInteger holds a nested list
#        """
#
#    def getList(self) -> [NestedInteger]:
#        """
#        @return the nested list that this NestedInteger holds, if it holds a nested list
#        Return None if this NestedInteger holds a single integer
#        """

from collections import deque


class NestedIterator:
    def __init__(self, nestedList: [NestedInteger]):
        self.stk = deque()
        self.stk.append(nestedList)
    
    def next(self) -> int:
        q = self.stk[-1]
        val = q[0].getInteger()
        self.stk[-1] = q[1:]
        return val
    
    def hasNext(self) -> bool:
        while self.stk:
            q = self.stk[-1]
            if not q:
                self.stk.pop()
                continue
            nest = q[0]
            if nest.isInteger():
                return True
            self.stk[-1] = q[1:]
            self.stk.append(nest.getList())
        return False
         

# Your NestedIterator object will be instantiated and called as such:
# i, v = NestedIterator(nestedList), []
# while i.hasNext(): v.append(i.next())