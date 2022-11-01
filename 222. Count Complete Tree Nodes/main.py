# Definition for a binary tree node.
class TreeNode:
    def __init__(self, val=0, left=None, right=None):
        self.val = val
        self.left = left
        self.right = right
from collections import deque
from typing import Optional


class Solution:
    def countNodes(self, root: Optional[TreeNode]) -> int:
        if not root:
            return 0
        
        q = deque()
        q.append(root)
        levels = 0
        while len(q) > 0:
            childs = 0
            size = len(q)
            levels += 1
            for i in range(0, size):
                node = q.popleft()
                if not node.right:
                    res = pow(2, levels) -1 + childs
                    if node.left:
                        res += 1
                    return res
                childs += 2
                q.append(node.left)
                q.append(node.right)
            
s = Solution()
s.countNodes(TreeNode(1, TreeNode(2,TreeNode(4), TreeNode(5)),TreeNode(3, TreeNode(6))))