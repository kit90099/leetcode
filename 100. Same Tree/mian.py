# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
from collections import deque
from typing import Optional


class Solution:
    def isSameTree(self, p: Optional[TreeNode], q: Optional[TreeNode]) -> bool:
        if not p and not q:
            return True
        if not p or not q:
            return False

        q1, q2 = deque(), deque()
        q1.append(p)
        q1.append(q)
        while q1 and q2:
            n1, n2 = q1.popleft(), q2.popleft()

            l1, l2 = n1.left, n2.left
            r1, r2 = n1.right, n2.right

            if n1.val != n2.val:
                return False

            if l1 ^ l2:
                return False
            if r1 ^ r2:
                return False

            if l1:
                q1.append(l1)
            if l2:
                q2.append(l2)
            if r1:
                q1.append(r1)
            if r2:
                q2.append(r2)
        
        return len(q1) == len(q2)