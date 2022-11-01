# Definition for a binary tree node.
from collections import deque
from typing import Optional, List


class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
class Solution:
    def rightSideView(self, root: Optional[TreeNode]) -> List[int]:
        if not root:
            return []
        q = deque()
        res = []
        q.append(root)
        while q:
            res.append(q[-1].val)
            size = len(q)
            for i in range(0, size):
                node = q.popleft()
                if node.right:
                    q.append(node.right)
                node = q.popleft()
                if node.left:
                    q.append(node.left)
        
        return res