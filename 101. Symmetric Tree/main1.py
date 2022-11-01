# Definition for a binary tree node.
from collections import deque
from typing import Optional


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def isSymmetric(self, root: Optional[TreeNode]) -> bool:
        q = deque()
        q.append(root)
        q.append(root)
        while len(q) > 0:
            a = q.popleft()
            b = q.popleft()

            if not a and not b:
                continue
            if not a or not b:
                return False
            if a.val != b.val:
                return False

            q.append(a.left)
            q.append(b.right)
            q.append(a.right)
            q.append(b.left)

        return True